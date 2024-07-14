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
	"github.com/gin-gonic/gin"
	"time"
)

type AdminUserLogic struct {
}

type IAdminUserLogic interface {
	Login(ctx context.Context, params *admin_proto.LoginReq, clientIp string) (*admin_proto.AdminInfo, error)
	Info(ctx context.Context, adminId int32, refreshToken bool, seconds int64) (*admin_proto.AdminInfo, error)
	Edit(ctx context.Context, adminId int32, params *admin_proto.AccountEditReq) error
	EditPassword(ctx *gin.Context, adminId int32, params *admin_proto.AccountPasswordEditReq) error
	MyMenus(ctx *gin.Context, adminId int32) (menus []*admin_proto.MenuItem, err error)
	MyPermission(ctx *gin.Context, adminId int32) (permissionKeys map[string]string, err error)
	Add(ctx *gin.Context, params *admin_proto.AdminUserAddReq) error
	List(ctx *gin.Context, params *admin_proto.AdminUserListReq) (*admin_proto.AdminUserListResp, error)
}

func newAdminUserLogic() IAdminUserLogic {
	return &AdminUserLogic{}
}

func (a *AdminUserLogic) Login(ctx context.Context, params *admin_proto.LoginReq, clientIp string) (*admin_proto.AdminInfo, error) {
	data, err := dao.H.AdminUser.FindAdminUserByUsername(ctx, params.Username)
	if err != nil {
		return nil, err
	}
	if !pwd.Matches(params.Password, data.Password) {
		return nil, code.NewCodeError(code_proto.ErrorCode_AdminAccountPasswordInvalid, nil)
	}

	result, err := getAccountInfo(ctx, data, true, config.AppConfig.Server.JWT.UsefulLife)
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

	return result, nil
}

func (a *AdminUserLogic) Info(ctx context.Context, adminId int32, refreshToken bool, seconds int64) (*admin_proto.AdminInfo, error) {
	data, err := dao.H.AdminUser.FindAdminUserByAdminId(ctx, adminId)
	if err != nil {
		return nil, err
	}

	return getAccountInfo(ctx, data, refreshToken, seconds)
}

func (a *AdminUserLogic) Edit(ctx context.Context, adminId int32, params *admin_proto.AccountEditReq) error {
	data, err := dao.H.AdminUser.FindAdminUserByAdminId(ctx, adminId)
	if err != nil {
		return err
	}
	data.Nickname = params.Nickname
	data.Avatar = params.Avatar
	data.Email = params.Email
	return dao.H.AdminUser.UpdateAdminUser(ctx, data)
}

func (a *AdminUserLogic) EditPassword(ctx *gin.Context, adminId int32, params *admin_proto.AccountPasswordEditReq) error {
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

func (a *AdminUserLogic) List(ctx *gin.Context, params *admin_proto.AdminUserListReq) (*admin_proto.AdminUserListResp, error) {
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
	data := &admin_proto.AdminUserListResp{}
	data.Total = total
	data.Rows, err = a.HandleListData(rows)
	return data, err
}

func (a *AdminUserLogic) HandleListData(list []*model.AdminUser) (rows []*admin_proto.AdminUserListItem, err error) {
	for _, item := range list {
		data, err := a.HandleItemData(item)
		if err != nil {
			return nil, err
		}
		rows = append(rows, data)
	}
	return rows, nil
}

func (a *AdminUserLogic) HandleItemData(item *model.AdminUser) (data *admin_proto.AdminUserListItem, err error) {
	data = &admin_proto.AdminUserListItem{}
	err = utils.BeanCopy(data, item)
	if err != nil {
		return nil, err
	}
	data.Enabled = item.IsEnabled
	data.CreatedAt = item.CreatedAt.Format(time.DateTime)
	data.UpdatedAt = item.UpdatedAt.Format(time.DateTime)
	return data, nil
}

func (a *AdminUserLogic) Add(ctx *gin.Context, params *admin_proto.AdminUserAddReq) error {
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
