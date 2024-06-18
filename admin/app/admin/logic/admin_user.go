package logic

import (
	"admin/app/admin/dao"
	"admin/app/gen/model"
	"admin/code"
	"admin/config"
	"admin/pkg/core"
	"admin/pkg/utils"
	"admin/pkg/utils/pwd"
	admin "admin/proto"
	"context"
	"time"
)

type AdminUserLogic struct {
	dao *dao.AdminUser
}

func NewAdminUserLogic() *AdminUserLogic {
	return &AdminUserLogic{
		dao: dao.NewAdminUser(),
	}
}

func (a *AdminUserLogic) Login(ctx context.Context, params *admin.LoginReq, clientIp string) (*admin.AdminInfo, error) {
	data, err := a.dao.FindAdminUserByUsername(ctx, params.Username)
	if err != nil {
		return nil, err
	}
	if !pwd.Matches(params.Password, data.Password) {
		return nil, code.NewCodeError(code.AdminAccountPasswordInvalid)
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
	if err := a.dao.UpdateAdminUserLoginData(ctx, data.ID, data); err != nil {
		return nil, err
	}

	return result, nil
}

func (a *AdminUserLogic) Info(ctx context.Context, adminId int32, refreshToken bool, seconds int64) (*admin.AdminInfo, error) {
	data, err := a.dao.FindAdminUserByAdminId(ctx, adminId)
	if err != nil {
		return nil, err
	}

	return a.getMyInfo(ctx, data, refreshToken, seconds)
}

func (a *AdminUserLogic) Edit(ctx context.Context, params *admin.AccountEditReq) error {
	// TODO
	return nil
}

func (a *AdminUserLogic) getMyInfo(ctx context.Context, data *model.AdminUser, refreshToken bool, seconds int64) (*admin.AdminInfo, error) {
	data.Password = ""
	resp := &admin.AdminInfo{}
	if err := utils.BeanCopy(data, resp); err != nil {
		return nil, err
	}
	if refreshToken {
		token, err := a.createToken(data.ID, data.Username, seconds)
		if err != nil {
			return nil, err
		}
		resp.Token = token
	}

	// 权限
	permissions, err := AdminPermissionSrv.FindMyPermission(ctx, data.ID)
	if err != nil {
		return nil, err
	}
	pageIds, perms := AdminPermissionSrv.Permissions2MenuIds(permissions)

	// 菜单
	menus, err := AdminMenuSrv.getMyMenusMap(ctx, pageIds)
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
