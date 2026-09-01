package logic

import (
	"admin/app/admin/dao"
	model2 "admin/app/admin/gen/model"
	"admin/code"
	"admin/constant"
	"admin/global"
	"admin/pkg/utils"
	"admin/pkg/utils/array"
	"admin/pkg/utils/pwd"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"context"
	"errors"
	"strings"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type AdminUserLogic struct {
}

type IAdminUserLogic interface {
	AccountLogin(ctx context.Context, params *admin_proto.ReqLogin, clientIp string) (*admin_proto.RespLoginData, error)
	AccountInfo(ctx context.Context, adminId int32, refreshToken bool, seconds int64) (*admin_proto.RespAccountInfoData, error)
	AccountEdit(ctx context.Context, adminId int32, params *admin_proto.ReqAccountEdit) error
	AccountEditPassword(ctx *gin.Context, adminId int32, params *admin_proto.ReqAccountPasswordEdit) error
	MyMenus(ctx *gin.Context, adminId int32) (menus []*admin_proto.MenuItem, err error)
	MyPermission(ctx *gin.Context, adminId int32) (permissionKeys map[string]string, err error)
	List(ctx *gin.Context, params *admin_proto.ReqAdminUserList) (*admin_proto.RespAdminUserListData, error)
	Add(ctx *gin.Context, params *admin_proto.ReqAdminUserAdd) error
	Edit(ctx *gin.Context, params *admin_proto.ReqAdminUserEdit) error
	EditPassword(ctx *gin.Context, params *admin_proto.ReqAdminUserEditPassword) error
	Enable(ctx *gin.Context, params *admin_proto.ReqAdminUserEnabled) error
	Delete(ctx *gin.Context, params *admin_proto.ReqAdminUserDelete) error
	BindRoles(ctx *gin.Context, params *admin_proto.ReqAdminUserBindRoles) error
}

func newAdminUserLogic() IAdminUserLogic {
	return &AdminUserLogic{}
}

func (a *AdminUserLogic) AccountLogin(ctx context.Context, params *admin_proto.ReqLogin, clientIp string) (*admin_proto.RespLoginData, error) {
	data, err := dao.H.AdminUser.FindAdminUserByUsername(ctx, params.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, code.NewCodeError(code_proto.ErrorCode_AdminAccountNotExist, nil)
	}
	if !pwd.Matches(params.Password, data.Password) {
		return nil, code.NewCodeError(code_proto.ErrorCode_AdminAccountPasswordInvalid, nil)
	}

	info, err := getAccountInfo(ctx, data, true, global.AppConfig.JWT.Expire)
	if err != nil {
		return nil, err
	}
	// 更新登录
	now := time.Now()
	data.LastLoginTime = &now
	data.LoginTotal += 1
	ip, err := dao.SetAdminUserLastLoginIp(clientIp, data.LastLoginIP)
	data.LastLoginIP = ip
	if err := dao.H.AdminUser.UpdateAdminUserLoginData(ctx, data.ID, data); err != nil {
		return nil, err
	}

	return &admin_proto.RespLoginData{
		Data: info,
	}, nil
}

func (a *AdminUserLogic) AccountInfo(ctx context.Context, adminId int32, refreshToken bool, seconds int64) (*admin_proto.RespAccountInfoData, error) {
	data, err := dao.H.AdminUser.FindAdminUserByAdminId(ctx, adminId)
	if err != nil {
		return nil, err
	}

	info, err := getAccountInfo(ctx, data, refreshToken, time.Duration(seconds))
	if err != nil {
		return nil, err
	}
	return &admin_proto.RespAccountInfoData{
		Data: info,
	}, nil
}

func (a *AdminUserLogic) AccountEdit(ctx context.Context, adminId int32, params *admin_proto.ReqAccountEdit) error {
	data, err := dao.H.AdminUser.FindAdminUserByAdminId(ctx, adminId)
	if err != nil {
		return err
	}
	data.Nickname = params.Nickname
	data.Avatar = params.Avatar
	data.Email = params.Email
	return dao.H.AdminUser.UpdateAdminUser(ctx, data)
}

func (a *AdminUserLogic) AccountEditPassword(ctx *gin.Context, adminId int32, params *admin_proto.ReqAccountPasswordEdit) error {
	data, err := dao.H.AdminUser.FindAdminUserByAdminId(ctx, adminId)
	if err != nil {
		return err
	}
	if !pwd.Matches(params.OldPassword, data.Password) {
		return code.NewCodeError(code_proto.ErrorCode_AdminAccountPasswordInvalid, nil)
	}
	data.Password, err = pwd.Encode(params.NewPassword)
	if err != nil {
		return err
	}
	return dao.H.AdminUser.UpdateAdminUser(ctx, data)
}

func (a *AdminUserLogic) MyMenus(ctx *gin.Context, adminId int32) (menus []*admin_proto.MenuItem, err error) {
	menus, _, err = getMyMenusAndPermissions(ctx, adminId)
	return menus, err
}

func (a *AdminUserLogic) MyPermission(ctx *gin.Context, adminId int32) (permissionKeys map[string]string, err error) {
	// 权限
	permissions, err := getMyAdminPermissions(ctx, adminId)
	if err != nil {
		return nil, err
	}
	_, permissionKeys = getMenuIdsFromAdminPermissions(permissions)
	return permissionKeys, nil
}

func (a *AdminUserLogic) List(ctx *gin.Context, params *admin_proto.ReqAdminUserList) (*admin_proto.RespAdminUserListData, error) {
	adminIds := make([]int32, 0)
	if params.RoleIds != nil && len(params.RoleIds) > 0 {
		adminRoles, err := dao.H.AdminRole.FindAdminUserByRoleIds(ctx, params.RoleIds)
		if err != nil {
			return nil, err
		}
		for _, item := range adminRoles {
			adminIds = append(adminIds, item.AdminID)
		}
	}
	total, rows, err := dao.H.AdminUser.List(ctx, params, adminIds)
	if err != nil {
		return nil, err
	}
	data := &admin_proto.RespAdminUserListData{}
	if total == 0 {
		return nil, nil
	}
	for _, item := range rows {
		adminIds = append(adminIds, item.ID)
	}
	adminIds = array.Deduplicate(adminIds, true, true)
	roles, err := dao.H.AdminUser.FindAdminUserRolesByAdminId(ctx, adminIds)
	if err != nil {
		return nil, err
	}
	rolesMap := make(map[int32][]*admin_proto.AdminUserRoleItem)
	for _, item := range roles {
		if _, ok := rolesMap[item.AdminId]; !ok {
			rolesMap[item.AdminId] = make([]*admin_proto.AdminUserRoleItem, 0)
		}
		rolesMap[item.AdminId] = append(rolesMap[item.AdminId], &admin_proto.AdminUserRoleItem{
			RoleId:   item.RoleId,
			RoleName: item.RoleName,
		})
	}
	data.Total = total
	data.PageSize = params.Base.PageSize
	data.PageNum = params.Base.PageNum
	data.List, err = a.HandleListData(rows, rolesMap)
	return data, err
}

func (a *AdminUserLogic) HandleListData(rows []*model2.AdminUser, rolesMap map[int32][]*admin_proto.AdminUserRoleItem) (list []*admin_proto.AdminUserListItem, err error) {
	for _, item := range rows {
		data, err := a.HandleItemData(item)
		if err != nil {
			return nil, err
		}
		if roles, ok := rolesMap[item.ID]; ok {
			data.Roles = roles
		}
		list = append(list, data)
	}
	return list, nil
}

func (a *AdminUserLogic) HandleItemData(item *model2.AdminUser) (data *admin_proto.AdminUserListItem, err error) {
	data = &admin_proto.AdminUserListItem{}
	err = utils.BeanCopy(data, item)
	if err != nil {
		return nil, err
	}
	data.AdminId = item.ID
	data.CreatedAt = utils.HandleTime2String(item.CreatedAt)
	data.UpdatedAt = utils.HandleTime2String(item.UpdatedAt)
	// 登录IP：数据库存 JSON 数组（最多保留最近两次登录 IP），第一个为上次登录IP，最后一个为本次登录IP
	lastIp, currentIp, err := getLoginIps(item.LastLoginIP)
	if err != nil {
		return nil, err
	}
	data.LastLoginIp = lastIp
	data.CurrentLoginIp = currentIp
	// 登录时间：数据库仅保存最近一次登录时间，即本次登录时间
	data.LastLoginTime = utils.HandleTimePointer2String(item.LastLoginTime)
	data.CurrentLoginTime = data.LastLoginTime
	// 超管标识：ID 为 1 的账号自动拥有全部权限
	data.IsSuperAdmin = constant.IsAdministrator(item.ID)
	return data, nil
}

func (a *AdminUserLogic) Add(ctx *gin.Context, params *admin_proto.ReqAdminUserAdd) error {
	password, err := pwd.Encode(params.Password)
	if err != nil {
		return err
	}
	data := &model2.AdminUser{
		Username:    params.Username,
		Password:    password,
		Nickname:    params.Nickname,
		Email:       params.Email,
		Avatar:      params.Avatar,
		IsEnabled:   true, // 前端不传启用状态，新增账号默认启用
		LastLoginIP: "[]",
	}
	if err = dao.H.AdminUser.Create(ctx, data); err != nil {
		// 账号唯一键冲突，返回“账号已存在”
		if strings.Contains(err.Error(), "uk_username") {
			return code.NewCodeError(code_proto.ErrorCode_AdminAccountNameExist, err)
		}
		return err
	}
	// 前端在创建时即绑定角色
	if len(params.RoleIds) > 0 {
		adminUserRoles := make([]*model2.AdminUserRole, 0, len(params.RoleIds))
		for _, roleId := range params.RoleIds {
			adminUserRoles = append(adminUserRoles, &model2.AdminUserRole{
				AdminID: data.ID,
				RoleID:  roleId,
			})
		}
		if err = dao.H.AdminUser.AddRoles(ctx, adminUserRoles); err != nil {
			return err
		}
	}
	return nil
}

func (a *AdminUserLogic) Edit(ctx *gin.Context, params *admin_proto.ReqAdminUserEdit) error {
	data, err := dao.H.AdminUser.FindAdminUserByAdminId(ctx, params.AdminId)
	if err != nil {
		return err
	}
	// 前端仅提交昵称与邮箱，账号/启用状态/头像保持不变
	data.Nickname = params.Nickname
	data.Email = params.Email
	return dao.H.AdminUser.UpdateAdminUser(ctx, data)
}

func (a *AdminUserLogic) EditPassword(ctx *gin.Context, params *admin_proto.ReqAdminUserEditPassword) error {
	data, err := dao.H.AdminUser.FindAdminUserByAdminId(ctx, params.AdminId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == gorm.ErrRecordNotFound {
		return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
	}
	// 前端不传确认密码，仅在确认密码有值时才比对
	if len(params.ConfirmPassword) != 0 && params.Password != params.ConfirmPassword {
		return code.NewCode(code_proto.ErrorCode_RequestParamsInvalid)
	}
	data.Password, err = pwd.Encode(params.Password)
	if err != nil {
		return err
	}
	return dao.H.AdminUser.UpdateAdminUser(ctx, data)
}

func (a *AdminUserLogic) Enable(ctx *gin.Context, params *admin_proto.ReqAdminUserEnabled) error {
	info, err := dao.H.AdminUser.FindAdminUserByAdminId(ctx, params.AdminId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	if info.IsEnabled == params.Enabled {
		return nil
	}
	return dao.H.AdminUser.Enable(ctx, params.AdminId, params.Enabled)
}

func (a *AdminUserLogic) Delete(ctx *gin.Context, params *admin_proto.ReqAdminUserDelete) error {
	info, err := dao.H.AdminUser.FindAdminUserByAdminId(ctx, params.AdminId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	if info.IsEnabled {
		return code.NewCodeError(code_proto.ErrorCode_RecordNValidCanNotDeleted, nil)
	}
	return dao.H.AdminUser.Delete(ctx, params.AdminId)
}

func (a *AdminUserLogic) BindRoles(ctx *gin.Context, params *admin_proto.ReqAdminUserBindRoles) error {
	// 超管账号自动拥有全部权限，不允许绑定角色
	if constant.IsAdministrator(params.AdminId) {
		return code.NewCodeError(code_proto.ErrorCode_AdminSuperAccountNotAllow, nil)
	}
	_, err := dao.H.AdminUser.FindAdminUserByAdminId(ctx, params.AdminId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_AdminAccountNotExist, err)
		}
		return err
	}
	//if !info.IsEnabled {
	//	return code.NewCodeError(code_proto.ErrorCode_AdminAccountInvalid, err)
	//}
	adminUserRoles := make([]*model2.AdminUserRole, 0, len(params.RoleIds))
	for _, roleId := range params.RoleIds {
		adminUserRoles = append(adminUserRoles, &model2.AdminUserRole{
			AdminID: params.AdminId,
			RoleID:  roleId,
		})
	}
	return dao.H.AdminUser.AddRoles(ctx, adminUserRoles)
}
