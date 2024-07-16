package admin_custom

import "time"

type AdminRoleInfo struct {
	ID              int32      `json:"id"`
	Name            string     `json:"name"`
	Describe        string     `json:"describe"`
	Enabled         bool       `json:"enabled"`
	CreateAdminId   int32      `json:"create_admin_id"`
	CreateAdminName string     `json:"create_admin_name"`
	ModifyAdminId   int32      `json:"modify_admin_id"`
	ModifyAdminName string     `json:"modify_admin_name"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}
