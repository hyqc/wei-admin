package service

import (
	"admin/app/admin/model"
	"admin/app/constant"
	"admin/proto/admin_menu"
)

type AdminMenuService struct {
	dao *model.AdminMenu
}

func NewAdminMenuService() *AdminMenuService {
	return &AdminMenuService{
		dao: adminMenuDao,
	}
}

func (m *AdminMenuService) getMyMenus(adminId int) ([]*admin_menu.MenuItem, error) {
	if constant.IsAdministrator(adminId) {
		// 超管
		return nil, nil
	}
	// 非超管
	return nil, nil
}
