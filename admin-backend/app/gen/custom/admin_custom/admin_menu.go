package admin_custom

import "time"

type AdminMenuItem struct {
	Id                 int32      `json:"id"`                    // 菜单ID
	Key                string     `json:"key"`                   // 菜单唯一键
	Name               string     `json:"name"`                  // 菜单名称
	ParentId           int32      `json:"parent_id"`             // 菜单父级ID
	Describe           string     `json:"describe"`              // 菜单描述
	Path               string     `json:"path"`                  // 菜单路径
	Redirect           string     `json:"redirect"`              // 菜单重定向
	Component          string     `json:"component"`             // 菜单组件路径
	Sort               int32      `json:"sort"`                  // 菜单排序
	Icon               string     `json:"icon"`                  // 菜单图标
	HideChildrenInMenu bool       `json:"hide_children_in_menu"` // 是否隐藏子菜单
	HideInMenu         bool       `json:"hide_in_menu"`          // 是否隐藏菜单
	Enabled            bool       `json:"enabled"`               // 是否启用
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
}

type AdminMenuTreeItem struct {
	Level    int32                `json:"level"`    // 菜单层级
	Children []*AdminMenuTreeItem `json:"children"` // 层级菜单
}
