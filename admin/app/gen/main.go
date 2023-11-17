package main

import (
	"admin/pkg/gormgen"
)

func main() {
	tables := []gormgen.GenType{
		{
			Table: "admin_user", Type: "AdminUser",
		},
		{
			Table: "admin_user_role", Type: "AdminUserRole",
		},
		{
			Table: "admin_role", Type: "AdminRole",
		},
		{
			Table: "admin_role_permission", Type: "AdminRolePermission",
		},
		{
			Table: "admin_permission", Type: "AdminPermission",
		},
		{
			Table: "admin_permission_api", Type: "AdminPermissionAPI",
		},
		{
			Table: "admin_api", Type: "AdminAPI",
		},
		{
			Table: "admin_menu", Type: "AdminMenu",
		},
	}
	gormgen.Init("root", "root", "127.0.0.1:3306", "wei", gormgen.Utf8mb4, tables)
}
