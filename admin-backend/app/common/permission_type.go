package common

import "admin/proto/admin_proto"

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
	AdminPermissionEnumView = &AdminPermissionEnum{
		Type: "view",
		Key:  "View",
		Name: "查看",
	}
	AdminPermissionEnumEdit = &AdminPermissionEnum{
		Type: "edit",
		Key:  "AccountEdit",
		Name: "编辑",
	}
	AdminPermissionEnumDelete = &AdminPermissionEnum{
		Type: "delete",
		Key:  "Delete",
		Name: "编辑",
	}
)

func AdminPermissionEnumList(menuId int32, key string) (list []*admin_proto.PermissionApiItem) {
	for _, item := range AdminPermissionEnumItems {
		list = append(list, &admin_proto.PermissionApiItem{
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
