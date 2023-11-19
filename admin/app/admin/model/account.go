package model

import (
	"admin/app/gen/model"
	"admin/app/gen/query"
	"context"
)

type IAdminAccount interface {
	FindAdminUserByUsernameAndPwd(ctx context.Context, username, pwd string) (*model.AdminUser, error)
}

type AdminAccount struct {
}

func NewAdminAccount() *AdminAccount {
	return &AdminAccount{}
}

func (a *AdminAccount) FindAdminUserByUsername(ctx context.Context, username string) (*model.AdminUser, error) {
	adminUser := query.AdminUser.Table(query.AdminUser.TableName())
	return adminUser.WithContext(ctx).Where(adminUser.Username.Eq(username)).First()
}
