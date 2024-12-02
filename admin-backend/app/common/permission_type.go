package common

import (
	"admin/proto/admin_proto"
	"fmt"
)

// AdminPermissionEnum 管理后台权限枚举
type AdminPermissionEnum struct {
	Type string `json:"type"`
	Key  string `json:"key"`
	Name string `json:"name"`
}

var (
	AdminPermissionEnumItems = []*AdminPermissionEnum{
		AdminPermissionEnumView,
		AdminPermissionEnumEdit,
		AdminPermissionEnumDelete,
	}

	AdminPermissionEnumMap = map[string]*AdminPermissionEnum{
		AdminPermissionEnumView.Type:   AdminPermissionEnumView,
		AdminPermissionEnumEdit.Type:   AdminPermissionEnumEdit,
		AdminPermissionEnumDelete.Type: AdminPermissionEnumDelete,
	}

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
)

func AdminPermissionEnumList(menuId int32, key string) (list []*admin_proto.MenuPermissionItem) {
	for _, item := range AdminPermissionEnumItems {
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

func GetPermissionTypeName(t string) (string, error) {
	if val, ok := AdminPermissionEnumMap[t]; ok {
		return val.Name, nil
	}
	return "", fmt.Errorf("未知权限类型:%v", t)
}
