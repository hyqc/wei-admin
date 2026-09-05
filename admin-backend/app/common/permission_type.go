package common

import (
	"admin/proto/admin_proto"
)

// AdminPermissionEnum 管理后台权限枚举
type AdminPermissionEnum struct {
	Type string `json:"type"`
	Key  string `json:"key"`
	Name string `json:"name"`
}

// 权限动作类型固定为三类，不允许新增：
//   - view：查看（读操作）
//   - edit：编辑（写操作，新增、重置、绑定等一切写操作统一归入此类）
//   - delete：删除
var (
	AdminPermissionEnumView = &AdminPermissionEnum{
		Type: "view",
		Key:  "View",
		Name: "查看",
	}
	AdminPermissionEnumEdit = &AdminPermissionEnum{
		Type: "edit",
		Key:  "Edit",
		Name: "编辑",
	}
	AdminPermissionEnumDelete = &AdminPermissionEnum{
		Type: "delete",
		Key:  "Delete",
		Name: "删除",
	}

	// AdminPermissionEnumItems 全部权限类型，供下拉选项等 UI 场景使用
	AdminPermissionEnumItems = []*AdminPermissionEnum{
		AdminPermissionEnumView,
		AdminPermissionEnumEdit,
		AdminPermissionEnumDelete,
	}

	// AdminPermissionEnumDefaultItems 菜单权限配置首次打开时的默认模板
	AdminPermissionEnumDefaultItems = []*AdminPermissionEnum{
		AdminPermissionEnumView,
		AdminPermissionEnumEdit,
		AdminPermissionEnumDelete,
	}

	AdminPermissionEnumMap = map[string]*AdminPermissionEnum{
		AdminPermissionEnumView.Type:   AdminPermissionEnumView,
		AdminPermissionEnumEdit.Type:   AdminPermissionEnumEdit,
		AdminPermissionEnumDelete.Type: AdminPermissionEnumDelete,
	}
)

func AdminPermissionEnumList(menuId int32, key string) (list []*admin_proto.MenuPermissionItem) {
	for _, item := range AdminPermissionEnumDefaultItems {
		list = append(list, &admin_proto.MenuPermissionItem{
			MenuId:   menuId,
			Type:     item.Type,
			Key:      key + item.Key,
			Name:     item.Name,
			Describe: item.Name,
			Enabled:  true,
		})
	}
	return list
}

// GetPermissionTypeName 获取权限类型展示名；未知类型直接返回原值，保证链路不断
func GetPermissionTypeName(t string) (string, error) {
	if val, ok := AdminPermissionEnumMap[t]; ok {
		return val.Name, nil
	}
	return t, nil
}
