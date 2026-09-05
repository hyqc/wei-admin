package logic

import (
	"admin/app/admin/dao"
	"admin/app/admin/gen/model"
	adminTypes "admin/app/admin/types"
	"admin/code"
	"admin/global"
	"admin/pkg/storage"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 分组（上传路径前缀）规则：仅允许字母数字、下划线、中划线、斜杠与点，且不允许出现 ..
var uploadGroupPattern = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_\-./]*$`)

type AdminUploadLogic struct {
}

type IAdminUploadLogic interface {
	List(ctx *gin.Context, params *admin_proto.ReqAdminUploadList) (*admin_proto.RespAdminUploadListData, error)
	// Upload 保存文件并写入记录；group 为上传路径前缀（如 /admin/user/）
	Upload(ctx *gin.Context, adminId int32, file *multipart.FileHeader, group string) (*admin_proto.AdminUploadItem, error)
	// Delete 删除记录，并尝试删除存储上的文件
	Delete(ctx *gin.Context, id int32) error
}

func newAdminUploadLogic() IAdminUploadLogic {
	return &AdminUploadLogic{}
}

// List 上传记录列表
func (a *AdminUploadLogic) List(ctx *gin.Context, params *admin_proto.ReqAdminUploadList) (*admin_proto.RespAdminUploadListData, error) {
	total, rows, err := dao.H.AdminUpload.List(ctx, params)
	if err != nil {
		return nil, err
	}
	data := &admin_proto.RespAdminUploadListData{
		List:  make([]*admin_proto.AdminUploadItem, 0, len(rows)),
		Total: total,
	}
	// 上传者名称：一次性批量查询，避免逐行回表
	names := make(map[int32]string, len(rows))
	if len(rows) > 0 {
		ids := make([]int32, 0, len(rows))
		for _, item := range rows {
			if item.AdminID > 0 {
				ids = append(ids, item.AdminID)
			}
		}
		if users, uerr := dao.H.AdminUser.FindByIds(ctx, ids); uerr == nil {
			for _, u := range users {
				names[u.ID] = u.Username
			}
		}
	}
	for _, item := range rows {
		data.List = append(data.List, a.toItem(item, names[item.AdminID]))
	}
	return data, nil
}

// Upload 上传文件
func (a *AdminUploadLogic) Upload(ctx *gin.Context, adminId int32, file *multipart.FileHeader, rawGroup string) (*admin_proto.AdminUploadItem, error) {
	if file == nil {
		return nil, code.NewCodeError(code_proto.ErrorCode_AdminUploadFileEmpty, nil)
	}
	cfg := &global.AppConfig.Upload

	group, err := normalizeUploadGroup(rawGroup)
	if err != nil {
		return nil, code.NewCodeError(code_proto.ErrorCode_AdminUploadGroupInvalid, err)
	}
	if cfg.MaxSize > 0 && file.Size > cfg.MaxSize<<20 {
		return nil, code.NewCodeError(code_proto.ErrorCode_AdminUploadFileTooLarge, nil)
	}
	ext := strings.ToLower(strings.TrimPrefix(path.Ext(file.Filename), "."))
	if len(cfg.AllowedExts) > 0 && !isAllowedExt(cfg.AllowedExts, ext) {
		return nil, code.NewCodeError(code_proto.ErrorCode_AdminUploadExtNotAllowed, nil)
	}

	f, err := file.Open()
	if err != nil {
		return nil, code.NewCodeError(code_proto.ErrorCode_AdminUploadFailed, err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, code.NewCodeError(code_proto.ErrorCode_AdminUploadFailed, err)
	}

	drv, err := storage.New(cfg)
	if err != nil {
		return nil, code.NewCodeError(code_proto.ErrorCode_AdminUploadDriverInvalid, err)
	}

	mime := file.Header.Get("Content-Type")
	now := time.Now()
	newName := fmt.Sprintf("%s_%s%s", now.Format("20060102_150405"), randHex(8), extWithDot(ext))
	// 目录结构：{分组}/{年}/{月}/{年月日_时分秒}_{随机}.{扩展名}
	objectKey := path.Join(group, now.Format("2006"), now.Format("01"), newName)

	if err = drv.Put(ctx, objectKey, data, mime); err != nil {
		return nil, code.NewCodeError(code_proto.ErrorCode_AdminUploadFailed, err)
	}

	sum := md5.Sum(data)
	record := &model.AdminUpload{
		AdminID:      adminId,
		Driver:       drv.Name(),
		UploadGroup:  group,
		ObjectKey:    objectKey,
		URL:          drv.URL(objectKey),
		OriginalName: file.Filename,
		NewName:      newName,
		Ext:          ext,
		Mime:         mime,
		Size:         file.Size,
		Md5:          hex.EncodeToString(sum[:]),
		UploadDate:   now,
	}
	if err = dao.H.AdminUpload.Create(ctx, record); err != nil {
		// 记录写入失败时回滚已上传的文件，避免产生无记录的孤儿文件
		if delErr := drv.Delete(ctx, objectKey); delErr != nil {
			global.LogSugar.Warnw("AdminUploadLogic.Upload: rollback object failed",
				zap.String("objectKey", objectKey), zap.Any("error", delErr))
		}
		return nil, err
	}
	return a.toItem(record, ""), nil
}

// Delete 删除上传记录，同时清理存储文件
func (a *AdminUploadLogic) Delete(ctx *gin.Context, id int32) error {
	record, err := dao.H.AdminUpload.Info(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	drv, err := storage.New(&global.AppConfig.Upload)
	if err != nil {
		return code.NewCodeError(code_proto.ErrorCode_AdminUploadDriverInvalid, err)
	}
	// 存储清理失败不阻断记录删除（否则文件已丢失时记录永远删不掉），仅记录日志
	if err = drv.Delete(ctx, record.ObjectKey); err != nil {
		global.LogSugar.Warnw("AdminUploadLogic.Delete: remove object failed",
			zap.String("objectKey", record.ObjectKey), zap.Any("error", err))
	}
	return dao.H.AdminUpload.Delete(ctx, id)
}

// toItem 转换为前端展示结构
func (a *AdminUploadLogic) toItem(item *model.AdminUpload, adminName string) *admin_proto.AdminUploadItem {
	return &admin_proto.AdminUploadItem{
		Id:           item.ID,
		AdminId:      item.AdminID,
		AdminName:    adminName,
		Driver:       item.Driver,
		DriverText:   driverText(item.Driver),
		UploadGroup:  item.UploadGroup,
		ObjectKey:    item.ObjectKey,
		Url:          item.URL,
		OriginalName: item.OriginalName,
		NewName:      item.NewName,
		Ext:          item.Ext,
		Mime:         item.Mime,
		Size:         item.Size,
		SizeText:     humanSize(item.Size),
		Md5:          item.Md5,
		UploadDate:   item.UploadDate.Format("2006-01-02"),
		CreatedAt:    item.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    item.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// normalizeUploadGroup 规范化分组：去掉首尾斜杠与多余层级，并做合法性校验
// 例：/admin/user/ → admin/user
func normalizeUploadGroup(raw string) (string, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return "", errors.New("upload group is empty")
	}
	// 先校验原始输入：path.Clean 会把 ../ 静默规整掉，直接 Clean 会让非法分组"变合法"
	if strings.Contains(trimmed, "..") || strings.Contains(trimmed, `\`) || strings.Contains(trimmed, "//") {
		return "", fmt.Errorf("illegal upload group: %s", raw)
	}
	group := strings.Trim(path.Clean("/"+trimmed), "/")
	if group == "" || group == "." {
		return "", errors.New("upload group is empty")
	}
	if !uploadGroupPattern.MatchString(group) {
		return "", fmt.Errorf("illegal upload group: %s", raw)
	}
	return group, nil
}

func isAllowedExt(allowed []string, ext string) bool {
	for _, item := range allowed {
		if strings.EqualFold(strings.TrimPrefix(item, "."), ext) {
			return true
		}
	}
	return false
}

func extWithDot(ext string) string {
	if ext == "" {
		return ""
	}
	return "." + ext
}

// randHex 生成随机十六进制串，用于避免同名文件冲突
func randHex(n int) string {
	buf := make([]byte, (n+1)/2)
	if _, err := rand.Read(buf); err != nil {
		// 随机源不可用时退化为时间戳，保证文件名仍然可用
		return fmt.Sprintf("%x", time.Now().UnixNano())[:n]
	}
	s := hex.EncodeToString(buf)
	if len(s) < n {
		return s
	}
	return s[:n]
}

// humanSize 把字节数转为可读文本
func humanSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	value := float64(size)
	units := []string{"KB", "MB", "GB", "TB"}
	for _, u := range units {
		value /= unit
		if value < unit {
			return fmt.Sprintf("%.2f %s", value, u)
		}
	}
	return fmt.Sprintf("%.2f PB", value/unit)
}

func driverText(driver string) string {
	switch driver {
	case storage.DriverLocal:
		return adminTypes.DriverTextLocal
	case storage.DriverAliyun:
		return adminTypes.DriverTextAliyun
	case storage.DriverQcloud:
		return adminTypes.DriverTextQcloud
	case storage.DriverS3:
		return adminTypes.DriverTextS3
	}
	return driver
}
