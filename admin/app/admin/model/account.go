package model

import (
	"admin/app/gen/model"
	"admin/app/gen/query"
	"context"
	"strings"
)

type IAdminAccount interface {
	FindAdminUserByUsername(ctx context.Context, username string) (*model.AdminUser, error)   // 根据管理员名称查询详情
	UpdateAdminUserLoginData(ctx context.Context, adminId int32, data *model.AdminUser) error // 更新管理员的登录信息
}

type AdminAccount struct {
}

var (
	adminUser = query.AdminUser.Table(query.AdminUser.TableName())
)

func NewAdminAccount() *AdminAccount {
	return &AdminAccount{}
}

func (a *AdminAccount) FindAdminUserByUsername(ctx context.Context, username string) (*model.AdminUser, error) {
	return adminUser.WithContext(ctx).Where(adminUser.Username.Eq(username)).First()
}

func (a *AdminAccount) UpdateAdminUserLoginData(ctx context.Context, adminId int32, data *model.AdminUser) error {
	_, err := adminUser.WithContext(ctx).Where(adminUser.ID.Eq(adminId)).
		Select(adminUser.LoginTotal, adminUser.LastLoginIP, adminUser.LastLoginTime).
		Updates(data)
	return err
}

// HandleAdminUserLastLoginIp 处理AdminUser.LastLoginIp，最多保存上一次登录的IP和本次登录的IP
func HandleAdminUserLastLoginIp(clientIp, ips string) string {
	const sep = ","
	clientIp = strings.Trim(clientIp, " ")
	if ips == "" {
		return clientIp
	}
	arr := strings.Split(ips, sep)
	tmp := make([]string, 0, len(arr))
	for _, ip := range arr {
		if ip != "" {
			tmp = append(tmp, ip)
		}
	}
	if len(tmp) == 0 {
		return clientIp
	}
	if len(tmp) == 1 {
		tmp = append(tmp, clientIp)
		return strings.Join(tmp, sep)
	}
	tmp[0], tmp[1] = tmp[1], clientIp
	return strings.Join(tmp, sep)
}
