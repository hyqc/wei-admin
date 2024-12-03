package types

type AdminUserRole struct {
	AdminId  int32  `gorm:"column:admin_id" json:"admin_id"`
	RoleId   int32  `gorm:"column:role_id" json:"role_id"`
	RoleName string `gorm:"column:role_name" json:"role_name"`
}
