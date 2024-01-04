package model

import (
	"admin/app/gen/model"
	"admin/app/gen/query"
	"context"
)

type IAdminMenu interface {
	FindMyMenusByAdminId(ctx context.Context, adminId, menuId int) (*model.AdminUser, error) // 根据管理员名称查询详情
	FindAdministerMenus(ctx context.Context) (*model.AdminUser, error)                       // 获取超管可以访问的菜单
	FindMyMenus(ctx context.Context, adminId, menuId int) (*model.AdminUser, error)          // 获取超管可以访问的菜单
}

type AdminMenu struct {
}

func NewIAdminMenu() *AdminMenu {
	return &AdminMenu{}
}

// FindAdministerMenus 获取超管的全部菜单
func (a *AdminMenu) FindAdministerMenus(ctx context.Context) (*model.AdminUser, error) {
	adminMenu := query.AdminMenu.Table(query.AdminMenu.TableName())
}
