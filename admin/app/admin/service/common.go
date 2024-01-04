package service

import "admin/app/admin/dao"

var (
	adminUserDao       = dao.NewAdminUser()
	adminMenuDao       = dao.NewAdminMenu()
	adminPermissionDao = dao.NewAdminPermission()

	AdminUserSrv       = NewAdminUserService()
	AdminMenuSrv       = NewAdminMenuService()
	AdminPermissionSrv = NewAdminPermissionService()
)
