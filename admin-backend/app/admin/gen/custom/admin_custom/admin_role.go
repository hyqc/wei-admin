package admin_custom

import "time"

type AdminRoleInfo struct {
	ID              int32      `json:"id"`
	Name            string     `json:"name"`
	Describe        string     `json:"describe"`
	IsEnabled       bool       `json:"is_enabled"`
	CreateAdminId   int32      `json:"create_admin_id"`
	CreateAdminName string     `json:"create_admin_name"`
	ModifyAdminId   int32      `json:"modify_admin_id"`
	ModifyAdminName string     `json:"modify_admin_name"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

type AdminRolePermissionItem struct {
	RoleID         int32  `json:"role_id"`
	PermissionID   int32  `json:"permission_id"`
	PermissionKey  string `json:"permission_key"`
	PermissionName string `json:"permission_name"`
	PermissionType string `json:"permission_type"`
}
