package logic

import (
	"admin/app/admin/dao"
	"admin/app/gen/model"
	"admin/code"
	"admin/config"
	"admin/pkg/core"
	"admin/pkg/utils"
	"admin/pkg/utils/pwd"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type AdminUserLogic struct {
	*dao.AdminUser
}

func NewAdminUserLogic() *AdminUserLogic {
	return &AdminUserLogic{
		dao.NewAdminUser(),
	}
}

func (a *AdminUserLogic) Login(ctx context.Context, params *admin_proto.LoginReq, clientIp string) (*admin_proto.AdminInfo, error) {
	data, err := a.FindAdminUserByUsername(ctx, params.Username)
	if err != nil {
		return nil, err
	}
	if !pwd.Matches(params.Password, data.Password) {
		return nil, code.NewCodeError(code_proto.ErrorCode_AdminAccountPasswordInvalid, nil)
	}

	result, err := a.getMyInfo(ctx, data, true, config.AppConfig.Server.JWT.UsefulLife)
	if err != nil {
		return nil, err
	}
	// 更新登录
	data.LastLoginTime = time.Now()
	data.LoginTotal += 1
	ip, err := dao.SetAdminUserLastLoginIp(clientIp, data.LastLoginIP)
	data.LastLoginIP = ip
	if err := a.UpdateAdminUserLoginData(ctx, data.ID, data); err != nil {
		return nil, err
	}

	return result, nil
}

func (a *AdminUserLogic) Info(ctx context.Context, adminId int32, refreshToken bool, seconds int64) (*admin_proto.AdminInfo, error) {
	data, err := a.FindAdminUserByAdminId(ctx, adminId)
	if err != nil {
		return nil, err
	}

	return a.getMyInfo(ctx, data, refreshToken, seconds)
}

func (a *AdminUserLogic) Edit(ctx context.Context, adminId int32, params *admin_proto.AccountEditReq) error {
	data, err := a.FindAdminUserByAdminId(ctx, adminId)
	if err != nil {
		return err
	}
	data.Nickname = params.Nickname
	data.Avatar = params.Avatar
	data.Email = params.Email
	return a.UpdateAdminUser(ctx, data)
}

func (a *AdminUserLogic) EditPassword(ctx *gin.Context, adminId int32, params *admin_proto.AccountPasswordEditReq) error {
	data, err := a.FindAdminUserByAdminId(ctx, adminId)
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
	return a.UpdateAdminUser(ctx, data)
}

func (a *AdminUserLogic) MyMenus(ctx *gin.Context, adminId int32) (menus []*admin_proto.MenuItem, err error) {
	menus, _, err = getMyMenusAndPermissions(ctx, adminId)
	return menus, err
}

func (a *AdminUserLogic) MyPermission(ctx *gin.Context, adminId int32) (permissionKeys map[string]string, err error) {
	// 权限
	permissions, err := AdminPermissionSrv.FindMyPermission(ctx, adminId)
	if err != nil {
		return nil, err
	}
	_, permissionKeys = AdminPermissionSrv.Permissions2MenuIds(permissions)
	return permissionKeys, nil
}

func (a *AdminUserLogic) getMyInfo(ctx context.Context, data *model.AdminUser, refreshToken bool, seconds int64) (*admin_proto.AdminInfo, error) {
	data.Password = ""
	resp := &admin_proto.AdminInfo{}
	if err := utils.BeanCopy(resp, data); err != nil {
		return nil, err
	}
	if refreshToken {
		token, err := a.createToken(data.ID, data.Username, seconds)
		if err != nil {
			return nil, err
		}
		resp.Token = token
	}

	// 菜单
	menus, perms, err := getMyMenusAndPermissions(ctx, data.ID)
	if err != nil {
		return nil, err
	}
	resp.Menus = menus
	resp.Permissions = perms

	return resp, err
}

func (a *AdminUserLogic) createToken(adminId int32, username string, seconds int64) (string, error) {
	// 生成token
	jti, err := core.Sonyflake.NextID()
	if err != nil {
		return "", err
	}
	token, err := core.JWTCreate(core.CustomClaimsOption{
		AccountId:     adminId,
		AccountName:   username,
		ExpireSeconds: time.Duration(seconds),
		UUID:          jti,
		Secret:        config.AppConfig.Server.JWT.Secret,
	})
	return token, err
}

func getMyMenusAndPermissions(ctx context.Context, adminId int32) (menus []*admin_proto.MenuItem, permissionKeys map[string]string, err error) {
	// 权限
	permissions, err := AdminPermissionSrv.FindMyPermission(ctx, adminId)
	if err != nil {
		return nil, nil, err
	}
	pageIds, permissionKeys := AdminPermissionSrv.Permissions2MenuIds(permissions)

	// 菜单
	menus, err = AdminMenuSrv.getMyMenusMap(ctx, pageIds)
	return menus, permissionKeys, err
}
