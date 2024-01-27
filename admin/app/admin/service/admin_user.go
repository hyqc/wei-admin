package service

import (
	"admin/app/admin/dao"
	"admin/app/gen/model"
	"admin/code"
	"admin/pkg/core"
	"admin/pkg/utils"
	"admin/pkg/utils/pwd"
	"admin/pkg/uuid"
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
	if !pwd.Matches(params.Password, data.Password) {
		return nil, code.NewCodeError(code.AdminAccountPasswordInvalid)
	}

	result, err := a.getMyInfo(ctx, data, true)
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

func (a *AdminUserService) Detail(ctx context.Context, adminId int32, refreshToken bool) (*admin_account.LoginDataResp, error) {
	data, err := a.dao.FindAdminUserByAdminId(ctx, adminId)
	if err != nil {
		return nil, err
	}

	return a.getMyInfo(ctx, data, refreshToken)
}

func (a *AdminUserService) getMyInfo(ctx context.Context, data *model.AdminUser, refreshToken bool) (*admin_account.LoginDataResp, error) {
	data.Password = ""
	resp := &admin_account.LoginDataResp{}
	if err := utils.BeanCopy(data, resp); err != nil {
		return nil, err
	}
	if refreshToken {
		token, err := a.createToken(data.ID, data.Username, 7)
		if err != nil {
			return nil, err
		}
		resp.Token = token.Raw
	}

	// 权限
	ps := PermissionService{}
	permissions, err := ps.FindMyPermission(ctx, data.ID)
	if err != nil {
		return nil, err
	}
	pageIds, perms := ps.Permissions2MenuIds(permissions)

	// 菜单
	menus, err := AdminMenuSrv.getMyMenusMap(ctx, pageIds)
	if err != nil {
		return nil, err
	}
	resp.Menus = menus
	resp.Permissions = perms

	return resp, err
}

func (a *AdminUserService) createToken(adminId int32, username string, days int64) (*jwt.Token, error) {
	// 生成token
	jti, err := uuid.Sonyflake.NextID()
	if err != nil {
		return nil, err
	}

	token := core.TokenCreate(adminId, username, jti, days)
	return token, err
}
