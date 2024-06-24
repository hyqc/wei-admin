package common

import "admin/proto/admin_proto"

func MenuTree(menusMap map[int32][]*admin_proto.MenuTreeItem, menusIds map[int32]bool, parentId, level, maxDeep int32) (data []*admin_proto.MenuTreeItem) {
	children, ok := menusMap[parentId]
	if !ok {
		return data
	}

	var result []*admin_proto.MenuTreeItem
	for _, menuItem := range children {
		menuItem.Level = level
		menuItem.Children = make([]*admin_proto.MenuTreeItem, 0, 10)

		if level+1 <= maxDeep {
			menuItem.Children = MenuTree(menusMap, menusIds, menuItem.Id, level+1, maxDeep)
			if len(menuItem.Children) == 0 {
				menuItem.Children = nil
			}
		}

		// 过滤条件调整为Go的实现方式
		if menusIds != nil && menuItem.Children == nil && len(menuItem.Children) > 0 {
			if !menusIds[parentId] && !menusIds[menuItem.Id] {
				continue
			}
		}

		result = append(result, menuItem)
	}

	return result
}

func GetMenuTreeWithTop(children []*admin_proto.MenuTreeItem) []*admin_proto.MenuTreeItem {
	menus := make([]*admin_proto.MenuTreeItem, 0, 1)
	topMenu := &admin_proto.MenuTreeItem{
		Level:              0,
		Id:                 0,
		ParentId:           0,
		Key:                "",
		Name:               "顶级菜单",
		Describe:           "顶级菜单",
		Path:               "/",
		Redirect:           "",
		Component:          "",
		Sort:               0,
		Icon:               "",
		HideChildrenInMenu: false,
		HideInMenu:         false,
		Enabled:            true,
		Children:           children,
	}
	return append(menus, topMenu)
}
