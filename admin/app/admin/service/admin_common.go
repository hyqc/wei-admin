package service

import "admin/app/admin/model"

var (
	adminUserDao       = model.NewAdminUser()
	adminMenuDao       = model.NewAdminMenu()
	adminPermissionDao = model.NewAdminPermission()
)
