package service

import (
	"admin/app/admin/dao"
	"admin/code"
	"admin/config"
	"admin/pkg/utils"
	"admin/proto/admin_account"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AdminUserService struct {
	dao *dao.AdminUser
}

func NewAdminUserService() *AdminUserService {
	return &AdminUserService{
		dao: dao.NewAdminUser(),
	}
}

func (a *AdminUserService) Login(ctx context.Context, params *admin_account.LoginReq, clientIp string) (*admin_account.LoginDataResp, error) {
	data, err := a.dao.FindAdminUserByUsername(ctx, params.Username)
	if err != nil {
		return nil, err
	}
	if !(utils.PasswordUtil{}).Matches(params.Password, data.Password) {
		return nil, code.NewCodeError(code.AdminAccountPasswordInvalid)
	}
	data.Password = ""

	// 生成token
	jti, err := config.AppSnoyflake.NextID()
	if err != nil {
		return nil, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": data.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7),
		"iat": time.Now().Unix(),
		"jti": jti,
	})

	resp := &admin_account.LoginDataResp{}
	if err := utils.BeanCopy(data, resp); err != nil {
		return nil, err
	}
	resp.Token = token.Raw

	// 权限
	ps := PermissionService{}
	permissions, err := ps.FindMyPermission(ctx, data.ID)
	if err != nil {
		return nil, err
	}
	pageIds, permis := ps.Permissions2MenuIds(permissions)
	// 菜单
	menus, err := AdminMenuSrv.getMyMenusMap(ctx, pageIds)
	if err != nil {
		return nil, err
	}
	resp.Menus = menus
	resp.Permissions = permis
	// 更新登录
	data.LastLoginTime = time.Now()
	data.LoginTotal += 1
	ip, err := dao.SetAdminUserLastLoginIp(clientIp, data.LastLoginIP)
	data.LastLoginIP = ip
	if err := a.dao.UpdateAdminUserLoginData(ctx, data.ID, data); err != nil {
		return nil, err
	}

	return resp, err
}
