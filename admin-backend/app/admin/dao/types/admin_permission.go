package types

import "time"

type AdminPermissionMenu struct {
	ID        int32      `json:"id"`
	MenuID    int32      `json:"menu_id"`
	MenuName  string     `json:"menu_name"`
	MenuPath  string     `json:"menu_path"`
	Key       string     `json:"key"`
	Name      string     `json:"name"`
	Describe  string     `json:"describe"`
	Type      string     `json:"type"`
	Enabled   bool       `json:"enabled"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// AdminApiPermissionItem API + Permission
type AdminApiPermissionItem struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:接口ID" json:"id"`                // 接口ID
	Path         string `gorm:"column:path;not null;comment:接口路由" json:"path"`                                 // 接口路由
	Key          string `gorm:"column:key;not null;comment:接口唯一名称" json:"key"`                                 // 接口唯一名称
	Name         string `gorm:"column:name;not null;comment:接口名称" json:"name"`                                 // 接口名称
	Describe     string `gorm:"column:describe;not null;comment:接口描述" json:"describe"`                         // 接口描述
	IsEnabled    bool   `gorm:"column:is_enabled;not null;default:1;comment:接口状态：1：正常，0：禁用" json:"is_enabled"` // 接口状态：1：正常，0：禁用
	PermissionId int32  `gorm:"column:permission_id;not null;comment:权限ID" json:"permission_id"`
}
