package model

import (
	"admin/app/gen/model"
	"admin/app/gen/query"
	"context"
)

type IAdminPermission interface {
	GetAdministerPermissions(ctx context.Context) ([]*model.AdminPermission, error) // 根据管理员名称查询详情
}

type AdminPermission struct {
}

func NewAdminPermission() *AdminPermission {
	return &AdminPermission{}
}

func (p AdminPermission) GetAdministerPermissions(ctx context.Context) ([]*model.AdminPermission, error) {
	permission := query.AdminPermission.Table(query.AdminPermission.TableName())
	menu := query.AdminMenu.Table(query.AdminMenu.TableName())
	return permission.WithContext(ctx).
		Join(query.AdminMenu, menu.ID.EqCol(permission.MenuID), menu.IsEnabled.Is(true)).
		Where(permission.IsEnabled.Is(true)).Find()
}
