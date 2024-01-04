package dao

import (
	"admin/app/gen/model"
	"admin/app/gen/query"
	"context"
)

type IAdminMenu interface {
	FindMyMenus(ctx context.Context, adminId, menuId int) ([]*model.AdminMenu, error) // 根据管理员名称查询详情
	FindAll(ctx context.Context) ([]*model.AdminMenu, error)                          // 获取全部有效菜单
}

type AdminMenu struct {
}

func NewAdminMenu() *AdminMenu {
	return &AdminMenu{}
}

// FindMyMenus 获取我的可以访问的菜单列表
func (a *AdminMenu) FindMyMenus(ctx context.Context, adminId, menuId int) ([]*model.AdminMenu, error) {
	return nil, nil
}

// FindAll 获取全部有效菜单
func (a *AdminMenu) FindAll(ctx context.Context) ([]*model.AdminMenu, error) {
	menu := query.AdminMenu.Table(query.AdminMenu.TableName())
	return menu.WithContext(ctx).Where(menu.IsEnabled.Is(true)).Order(menu.Sort, menu.ParentID, menu.ID).Find()
}
