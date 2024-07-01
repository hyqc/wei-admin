package dao

import (
	"admin/app/gen/model"
	"admin/app/gen/query"
	"context"
)

type IAdminPermission interface {
	FindAdministerPermissions(ctx context.Context) ([]*model.AdminPermission, error) // 根据管理员名称查询详情
	FindAdminPermissions(ctx context.Context, adminId, menuId int32) ([]*model.AdminPermission, error)
	FindPermissionsByMenuId(ctx context.Context, menuId int32) ([]*model.AdminPermission, error)
}

type AdminPermission struct {
}

func NewAdminPermission() *AdminPermission {
	return &AdminPermission{}
}

// FindAdministerPermissions 获取超管对应的权限
func (a *AdminPermission) FindAdministerPermissions(ctx context.Context) ([]*model.AdminPermission, error) {
	permission := query.AdminPermission
	menu := query.AdminMenu
	return permission.WithContext(ctx).
		Join(menu, menu.ID.EqCol(permission.MenuID), menu.IsEnabled.Is(true)).
		Where(permission.IsEnabled.Is(true)).Order(permission.MenuID).Find()
}

// FindAdminPermissions 获取非超管对于的权限
func (a *AdminPermission) FindAdminPermissions(ctx context.Context, adminId, menuId int32) ([]*model.AdminPermission, error) {
	permission := query.AdminPermission
	menu := query.AdminMenu
	role := query.AdminRole
	rolePermission := query.AdminRolePermission
	userRole := query.AdminUserRole

	db := permission.WithContext(ctx).
		Join(menu, menu.ID.EqCol(permission.MenuID), menu.IsEnabled.Is(true)).
		Join(rolePermission, rolePermission.RoleID.EqCol(userRole.RoleID)).
		Join(role, role.ID.EqCol(userRole.RoleID), role.IsEnabled.Is(true)).
		Join(userRole, userRole.RoleID.EqCol(role.ID), userRole.AdminID.Eq(adminId)).
		Where(userRole.AdminID.Eq(adminId))
	if menuId > 0 {
		db.Where(permission.MenuID.Eq(menuId))
	}
	return db.Order(permission.MenuID).Find()
}

func (a *AdminPermission) FindPermissionsByMenuId(ctx context.Context, menuId int32) ([]*model.AdminPermission, error) {
	permission := query.AdminPermission
	return permission.WithContext(ctx).Where(permission.MenuID.Eq(menuId)).Find()
}
