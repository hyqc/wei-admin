package service

import "admin/app/admin/model"

type AdminPermissionService struct {
	dao *model.AdminPermission
}

func NewAdminPermissionService() *AdminPermissionService {
	return &AdminPermissionService{
		dao: adminPermissionDao,
	}
}
