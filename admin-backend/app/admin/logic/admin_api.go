package logic

import (
	"admin/app/admin/dao"
	"admin/app/gen/model"
	"admin/code"
	"admin/pkg/utils"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"time"
)

type AdminAPILogic struct {
}

type IAdminAPILogic interface {
	List(ctx *gin.Context, params *admin_proto.ReqApiList) (data *admin_proto.RespApiListData, err error)
	AllValid(ctx *gin.Context) (list []*admin_proto.ApiItem, err error)
	Add(ctx *gin.Context, params *admin_proto.ReqApiAdd) (err error)
	Info(ctx *gin.Context, params *admin_proto.ReqApiInfo) (*admin_proto.ApiItem, error)
	Edit(ctx *gin.Context, params *admin_proto.ReqApiEdit) error
	Enable(ctx *gin.Context, params *admin_proto.ReqApiEnable) error
	Delete(ctx *gin.Context, params *admin_proto.ReqApiDelete) error
}

func newAdminAPILogic() IAdminAPILogic {
	return &AdminAPILogic{}
}

func (a *AdminAPILogic) List(ctx *gin.Context, params *admin_proto.ReqApiList) (data *admin_proto.RespApiListData, err error) {
	total, rows, err := dao.H.AdminAPI.List(ctx, params)
	if err != nil {
		return nil, err
	}
	data = &admin_proto.RespApiListData{}
	data.Total = total
	data.List, err = a.HandleListData(rows)
	return data, err
}

func (a *AdminAPILogic) HandleListData(rows []*model.AdminAPI) (list []*admin_proto.ApiItem, err error) {
	for _, item := range rows {
		data, err := a.HandleItemData(item)
		if err != nil {
			return nil, err
		}
		list = append(list, data)
	}
	return list, nil
}

func (a *AdminAPILogic) HandleItemData(item *model.AdminAPI) (data *admin_proto.ApiItem, err error) {
	data = &admin_proto.ApiItem{}
	err = utils.BeanCopy(data, item)
	if err != nil {
		return nil, err
	}
	data.CreatedAt = item.CreatedAt.Format(time.DateTime)
	data.UpdatedAt = item.UpdatedAt.Format(time.DateTime)
	return data, nil
}

func (a *AdminAPILogic) AllValid(ctx *gin.Context) (list []*admin_proto.ApiItem, err error) {
	data, err := dao.H.AdminAPI.FindAllValid(ctx)
	if err != nil {
		return nil, err
	}
	list, err = a.HandleListData(data)
	return list, err
}

func (a *AdminAPILogic) Add(ctx *gin.Context, params *admin_proto.ReqApiAdd) (err error) {
	data := &model.AdminAPI{
		Path:      params.Path,
		Key:       params.Key,
		Name:      params.Name,
		Describe:  params.Describe,
		IsEnabled: params.Enabled,
	}
	if err := dao.H.AdminAPI.Create(ctx, data); err != nil {
		msg := err.Error()
		if strings.Contains(msg, "uk_name") {
			return code.NewCodeError(code_proto.ErrorCode_AdminApiNameExist, err)
		}
		if strings.Contains(msg, "uk_path") {
			return code.NewCodeError(code_proto.ErrorCode_AdminApiPathExist, err)
		}
		if strings.Contains(msg, "uk_key") {
			return code.NewCodeError(code_proto.ErrorCode_AdminApiKeyExist, err)
		}
		return err
	}
	return nil
}

func (a *AdminAPILogic) Info(ctx *gin.Context, params *admin_proto.ReqApiInfo) (*admin_proto.ApiItem, error) {
	data, err := dao.H.AdminAPI.Info(ctx, params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return nil, err
	}
	return a.HandleItemData(data)
}

func (a *AdminAPILogic) Edit(ctx *gin.Context, params *admin_proto.ReqApiEdit) error {
	info, err := dao.H.AdminAPI.Info(ctx, params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	info.Key = params.Key
	info.Path = params.Path
	info.Name = params.Name
	info.Describe = params.Describe
	info.IsEnabled = params.Enabled
	return dao.H.AdminAPI.Update(ctx, info)
}

func (a *AdminAPILogic) Enable(ctx *gin.Context, params *admin_proto.ReqApiEnable) error {
	info, err := dao.H.AdminAPI.Info(ctx, params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	if info.IsEnabled == params.Enabled {
		return nil
	}
	return dao.H.AdminAPI.Enable(ctx, params.Id, params.Enabled)
}

func (a *AdminAPILogic) Delete(ctx *gin.Context, params *admin_proto.ReqApiDelete) error {
	info, err := dao.H.AdminAPI.Info(ctx, params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	if info.IsEnabled {
		return code.NewCodeError(code_proto.ErrorCode_RecordNValidCanNotDeleted, nil)
	}
	return dao.H.AdminAPI.Delete(ctx, params.Id)
}
