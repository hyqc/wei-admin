package dao

import (
	"admin/app/gen/model"
	"admin/app/gen/query"
	"context"
)

type IAdminRole interface {
	FindAdminUserByRoleIds(ctx context.Context, ids []int32) ([]*model.AdminUserRole, error)
}

type AdminRole struct {
}

func newAdminRole() IAdminRole {
	return &AdminRole{}
}

func (a *AdminRole) FindAdminUserByRoleIds(ctx context.Context, ids []int32) ([]*model.AdminUserRole, error) {
	return query.AdminUserRole.WithContext(ctx).Where(query.AdminUserRole.RoleID.In(ids...)).Find()
}
