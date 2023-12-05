package dao

import (
	"admin/app/gen/model"
	"admin/app/gen/query"
	"context"
)

type IAdminUser interface {
	FindAdminUserByUsername(ctx context.Context, username string) (*model.AdminUser, error)   // 根据管理员名称查询详情
	UpdateAdminUserLoginData(ctx context.Context, adminId int32, data *model.AdminUser) error // 更新管理员的登录信息
}

type AdminUser struct {
}

func NewAdminUser() *AdminUser {
	return &AdminUser{}
}

func (a *AdminUser) FindAdminUserByUsername(ctx context.Context, username string) (*model.AdminUser, error) {
	adminUser := query.AdminUser.Table(query.AdminUser.TableName())
	return adminUser.WithContext(ctx).Where(adminUser.Username.Eq(username)).First()
}

func (a *AdminUser) UpdateAdminUserLoginData(ctx context.Context, adminId int32, data *model.AdminUser) error {
	adminUser := query.AdminUser.Table(query.AdminUser.TableName())
	_, err := adminUser.WithContext(ctx).Where(adminUser.ID.Eq(adminId)).
		Select(adminUser.LoginTotal, adminUser.LastLoginIP, adminUser.LastLoginTime).
		Updates(data)
	return err
}
