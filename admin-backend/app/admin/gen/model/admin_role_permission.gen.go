// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAdminRolePermission = "admin_role_permission"

// AdminRolePermission mapped from table <admin_role_permission>
type AdminRolePermission struct {
	RoleID       int32 `gorm:"column:role_id;primaryKey;comment:角色ID" json:"role_id"`             // 角色ID
	PermissionID int32 `gorm:"column:permission_id;primaryKey;comment:权限ID" json:"permission_id"` // 权限ID
}

// TableName AdminRolePermission's table name
func (*AdminRolePermission) TableName() string {
	return TableNameAdminRolePermission
}
