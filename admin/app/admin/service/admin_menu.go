package service

import (
	"admin/app/admin/dao"
	"admin/proto/admin_menu"
	"context"
)

type AdminMenuService struct {
	dao *dao.AdminMenu
}

func NewAdminMenuService() *AdminMenuService {
	return &AdminMenuService{
		dao: adminMenuDao,
	}
}

func (m *AdminMenuService) getMyMenus(ctx context.Context, menuIds []int32) ([]*admin_menu.MenuItem, error) {
	// 获取全部的菜单

	// 非超管
	return nil, nil
}
