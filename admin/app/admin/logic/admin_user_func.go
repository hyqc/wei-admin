package logic

import (
	"admin/app/gen/model"
	"admin/config"
	"admin/pkg/core"
	"admin/pkg/utils"
	"admin/proto/admin_proto"
	"context"
	"time"
)

func getAccountInfo(ctx context.Context, data *model.AdminUser, refreshToken bool, seconds int64) (*admin_proto.AdminInfo, error) {
	data.Password = ""
	resp := &admin_proto.AdminInfo{}
	if err := utils.BeanCopy(resp, data); err != nil {
		return nil, err
	}
	if refreshToken {
		token, err := createToken(data.ID, data.Username, seconds)
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

	menusMap := make(map[string]*admin_proto.MenuItem)
	for _, item := range menus {
		menusMap[item.Key] = item
	}

	resp.Menus = menusMap
	resp.Permissions = perms

	return resp, err
}

func createToken(adminId int32, username string, seconds int64) (string, error) {
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
	permissions, err := getMyAdminPermissions(ctx, adminId)
	if err != nil {
		return nil, nil, err
	}
	pageIds, permissionKeys := getMenuIdsFromAdminPermissions(permissions)

	// 菜单
	menus, err = getMyMenusMap(ctx, pageIds)
	return menus, permissionKeys, err
}
