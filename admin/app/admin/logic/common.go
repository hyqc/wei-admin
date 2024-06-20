package logic

import "admin/app/admin/dao"

var (
	adminUserDao       = dao.NewAdminUser()
	adminMenuDao       = dao.NewAdminMenu()
	adminPermissionDao = dao.NewAdminPermission()
	adminAPIDao        = dao.NewAdminAPI()

	AdminUserSrv       = NewAdminUserLogic()
	AdminMenuSrv       = NewAdminMenuLogic()
	AdminPermissionSrv = NewAdminPermissionLogic()
)
