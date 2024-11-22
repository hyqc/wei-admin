package main

import (
	"admin/pkg/gormgen"
)

func main() {
	//isEnabledFieldType := []gen.ModelOpt{gen.FieldType("is_enabled", "bool")}

	tables := []gormgen.GenType{
		{
			Table: "admin_user",
			Type:  "AdminUser",
			//Fields: isEnabledFieldType,
		},
		{
			Table: "admin_user_role",
			Type:  "AdminUserRole",
		},
		{
			Table: "admin_role",
			Type:  "AdminRole",
			//Fields: isEnabledFieldType,
		},
		{
			Table: "admin_role_permission",
			Type:  "AdminRolePermission",
		},
		{
			Table: "admin_permission",
			Type:  "AdminPermission",
			//Fields: isEnabledFieldType,
		},
		{
			Table: "admin_permission_api",
			Type:  "AdminPermissionAPI",
		},
		{
			Table: "admin_api",
			Type:  "AdminAPI",
			//Fields: isEnabledFieldType,
		},
		{
			Table: "admin_menu",
			Type:  "AdminMenu",
			//Fields: append(isEnabledFieldType,
			//	gen.FieldType("is_hide_children_in_menu", "bool"),
			//	gen.FieldType("is_hide_in_menu", "bool"),
			//),
		},
	}
	gormgen.Init("root", "123456", "127.0.0.1:3306", "wei", gormgen.Utf8mb4, tables, "./app/gen/query")
}
