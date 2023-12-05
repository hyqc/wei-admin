package dao

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

// GetAdministerPermissions 获取超管对应的权限
func (p *AdminPermission) GetAdministerPermissions(ctx context.Context) ([]*model.AdminPermission, error) {
	permission := query.AdminPermission.Table(query.AdminPermission.TableName())
	menu := query.AdminMenu.Table(query.AdminMenu.TableName())
	return permission.WithContext(ctx).
		Join(menu, menu.ID.EqCol(permission.MenuID), menu.IsEnabled.Is(true)).
		Where(permission.IsEnabled.Is(true)).Find()
}

// GetMyPermissions 获取非超管对于的权限
func (p *AdminPermission) GetMyPermissions(ctx context.Context, adminId, menuId int32) ([]*model.AdminPermission, error) {
	permission := query.AdminPermission.Table(query.AdminPermission.TableName())
	menu := query.AdminMenu.Table(query.AdminMenu.TableName())
	role := query.AdminRole.Table(query.AdminRole.TableName())
	rolePermission := query.AdminRolePermission.Table(query.AdminRolePermission.TableName())
	userRole := query.AdminUserRole.Table(query.AdminUserRole.TableName())

	db := permission.WithContext(ctx).
		Join(menu, menu.ID.EqCol(permission.MenuID), menu.IsEnabled.Is(true)).
		Join(rolePermission, rolePermission.RoleID.EqCol(userRole.RoleID)).
		Join(role, role.ID.EqCol(userRole.RoleID), role.IsEnabled.Is(true)).
		Join(userRole, userRole.RoleID.EqCol(role.ID), userRole.AdminID.Eq(adminId)).
		Where(userRole.AdminID.Eq(adminId))
	if menuId > 0 {
		db.Where(permission.MenuID.Eq(menuId))
	}
	return db.Find()
}
