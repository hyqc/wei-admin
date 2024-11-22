package logic

import (
	"admin/app/admin/dao"
	"admin/app/gen/model"
	"admin/code"
	"admin/config"
	"admin/pkg/utils"
	"admin/pkg/utils/pwd"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"context"
	"errors"
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

	info, err := getAccountInfo(ctx, data, true, config.AppConfig.Server.JWT.UsefulLife)
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

	info, err := getAccountInfo(ctx, data, refreshToken, seconds)
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
	data.Password, err = pwd.Encode(params.Password)
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
	data.Total = total
	data.List, err = a.HandleListData(rows)
	return data, err
}

func (a *AdminUserLogic) HandleListData(rows []*model.AdminUser) (list []*admin_proto.AdminUserModel, err error) {
	for _, item := range rows {
		data, err := a.HandleItemData(item)
		if err != nil {
			return nil, err
		}
		list = append(list, data)
	}
	return list, nil
}

func (a *AdminUserLogic) HandleItemData(item *model.AdminUser) (data *admin_proto.AdminUserModel, err error) {
	data = &admin_proto.AdminUserModel{}
	err = utils.BeanCopy(data, item)
	if err != nil {
		return nil, err
	}
	data.CreatedAt = item.CreatedAt.Format(time.DateTime)
	data.UpdatedAt = item.UpdatedAt.Format(time.DateTime)
	return data, nil
}

func (a *AdminUserLogic) Add(ctx *gin.Context, params *admin_proto.ReqAdminUserAdd) error {
	password, err := pwd.Encode(params.Password)
	if err != nil {
		return err
	}
	data := &model.AdminUser{
		Username:  params.Username,
		Password:  password,
		Nickname:  params.Nickname,
		Email:     params.Email,
		Avatar:    params.Avatar,
		IsEnabled: params.Enabled,
	}
	return dao.H.AdminUser.Create(ctx, data)
}

func (a *AdminUserLogic) Edit(ctx *gin.Context, params *admin_proto.ReqAdminUserEdit) error {
	data, err := dao.H.AdminUser.FindAdminUserByAdminId(ctx, params.AdminId)
	if err != nil {
		return err
	}
	if params.Password != "" {
		data.Password, err = pwd.Encode(params.Password)
	}
	data.Username = params.Username
	data.Nickname = params.Nickname
	data.Email = params.Email
	data.IsEnabled = params.Enabled
	data.Avatar = params.Avatar
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
	info, err := dao.H.AdminUser.FindAdminUserByAdminId(ctx, params.AdminId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_AdminAccountNotExist, err)
		}
		return err
	}
	if !info.IsEnabled {
		return code.NewCodeError(code_proto.ErrorCode_AdminAccountInvalid, err)
	}
	adminUserRoles := make([]*model.AdminUserRole, 0, len(params.RoleIds))
	for _, roleId := range params.RoleIds {
		adminUserRoles = append(adminUserRoles, &model.AdminUserRole{
			AdminID: params.AdminId,
			RoleID:  roleId,
		})
	}
	return dao.H.AdminUser.AddRoles(ctx, adminUserRoles)
}