package logic

import (
	"admin/app/admin/dao"
	"admin/app/admin/gen/model"
	"admin/global"
	"admin/pkg/utils"
	"admin/pkg/utils/jwt"
	"admin/proto/admin_proto"
	"context"
	"encoding/json"
	"time"
)

func getAccountInfo(ctx context.Context, data *model.AdminUser, refreshToken bool, seconds time.Duration) (*admin_proto.AdminInfo, error) {
	data.Password = ""
	resp := &admin_proto.AdminInfo{
		AdminId:       data.ID,
		Username:      data.Username,
		Nickname:      data.Nickname,
		Avatar:        data.Avatar,
		Email:         data.Email,
		CreatedAt:     utils.HandleTime2String(data.CreatedAt),
		UpdatedAt:     utils.HandleTime2String(data.UpdatedAt),
		LastLoginTime: utils.HandleTimePointer2String(data.LastLoginTime),
		LoginTotal:    data.LoginTotal,
		Enabled:       data.IsEnabled,
		Menus:         make(map[string]*admin_proto.MenuItem),
		Permissions:   make(map[string]string),
	}
	// 登录IP：last_login_ip 为 JSON 数组，最多保留最近两次登录IP（第一个为上次，最后一个为本次）
	lastLoginIp, currentLoginIp, err := getLoginIps(data.LastLoginIP)
	if err != nil {
		return nil, err
	}
	resp.LastLoginIp = lastLoginIp
	resp.CurrentLoginIp = currentLoginIp
	// 数据库只保存最近一次登录时间，即本次登录时间
	resp.CurrentLoginTime = resp.LastLoginTime
	if refreshToken {
		token, err := createToken(data.ID, data.Username, int64(seconds))
		if err != nil {
			return nil, err
		}
		resp.Token = token.Token
		// 前端以 expire(秒) * 1000 作为过期时间戳，这里返回过期时间的秒级时间戳
		resp.Expire = token.ExpiryAt.Unix()
		resp.ExpireDataTime = utils.HandleTime2String(token.ExpiryAt)
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

	// 角色
	roles, err := dao.H.AdminUser.FindAdminUserRolesByAdminId(ctx, []int32{data.ID})
	if err != nil {
		return nil, err
	}
	if len(roles) != 0 {
		for _, item := range roles {
			resp.Roles = append(resp.Roles, &admin_proto.AdminUserRoleItem{
				RoleId:   item.RoleId,
				RoleName: item.RoleName,
			})
		}
	}

	resp.Menus = menusMap
	resp.Permissions = perms

	return resp, err
}

func createToken(adminId int32, username string, seconds int64) (*jwt.Token, error) {
	// 生成token
	data := jwt.ClaimsData{
		AccountId:     adminId,
		ExpireSeconds: seconds,
	}
	return global.AppAuth.Generate(data)
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

// getLoginIps 解析 last_login_ip（JSON 数组，最多保留最近两次登录IP），返回上次登录IP与本次登录IP
func getLoginIps(ips string) (lastIp, currentIp string, err error) {
	res := make([]string, 0)
	if err = json.Unmarshal([]byte(ips), &res); err != nil {
		return "", "", err
	}
	if len(res) == 0 {
		return "", "", nil
	}
	return res[0], res[len(res)-1], nil
}
