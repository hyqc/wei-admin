package dao

import (
	"admin/app/gen/model"
	"admin/app/gen/query"
	"context"
)

type IAdminUser interface {
	FindAdminUserByUsername(ctx context.Context, username string) (*model.AdminUser, error)   // 根据管理员名称查询详情
	UpdateAdminUserLoginData(ctx context.Context, adminId int32, data *model.AdminUser) error // 更新管理员的登录信息
	FindAdminUserByAdminId(ctx context.Context, id int32) (*model.AdminUser, error)           // 根据管理员ID查询详情
	UpdateAdminUser(ctx context.Context, data *model.AdminUser) error
}

type AdminUser struct {
}

func newAdminUser() *AdminUser {
	return &AdminUser{}
}

func (a *AdminUser) FindAdminUserByUsername(ctx context.Context, username string) (*model.AdminUser, error) {
	db := query.AdminUser.Table(query.AdminUser.TableName())
	return db.WithContext(ctx).Where(db.Username.Eq(username)).First()
}

func (a *AdminUser) UpdateAdminUserLoginData(ctx context.Context, adminId int32, data *model.AdminUser) error {
	db := query.AdminUser.Table(query.AdminUser.TableName())
	_, err := db.WithContext(ctx).Where(db.ID.Eq(adminId)).
		Select(db.LoginTotal, db.LastLoginIP, db.LastLoginTime).
		Updates(data)
	return err
}

func (a *AdminUser) FindAdminUserByAdminId(ctx context.Context, id int32) (*model.AdminUser, error) {
	db := query.AdminUser.Table(query.AdminUser.TableName())
	return db.WithContext(ctx).Where(db.ID.Eq(id)).First()
}

func (a *AdminUser) UpdateAdminUser(ctx context.Context, data *model.AdminUser) error {
	db := query.AdminUser.Table(query.AdminUser.TableName())
	_, err := db.WithContext(ctx).Where(db.ID.Eq(data.ID)).Updates(data)
	return err
}
