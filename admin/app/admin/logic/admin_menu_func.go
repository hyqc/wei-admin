package logic

import (
	"admin/app/admin/dao"
	"admin/app/gen/model"
	"admin/proto/admin_proto"
	"context"
)

func getMyMenusMap(ctx context.Context, pageIds []int32) ([]*admin_proto.MenuItem, error) {
	allMenus, err := dao.H.AdminMenu.FindAllValid(ctx)
	if err != nil {
		return nil, err
	}
	allMenusMap := make(map[int32]*model.AdminMenu)
	for _, item := range allMenus {
		allMenusMap[item.ID] = item
	}
	arr1 := getAllRelatedMenusByPageIds(allMenusMap, pageIds)
	arr2 := getAllChildrenPagesByPageIds(allMenusMap, pageIds)
	arr := append(arr1, arr2...)
	menus := make([]*admin_proto.MenuItem, 0, len(arr))
	m := make(map[string]struct{}, len(arr))
	for _, item := range arr {
		if _, ok := m[item.Key]; !ok {
			m[item.Key] = struct{}{}
			menus = append(menus, item)
		}
	}
	return menus, nil
}

func getAllRelatedMenusByPageIds(menusMap map[int32]*model.AdminMenu, pageIds []int32) (data []*admin_proto.MenuItem) {
	for _, pageId := range pageIds {
		arr := getAllFatherMenuByChildrenId(menusMap, pageId)
		if len(arr) > 0 {
			data = append(data, arr...)
		}
	}
	return data
}

func getAllChildrenPagesByPageIds(menusMap map[int32]*model.AdminMenu, pageIds []int32) (data []*admin_proto.MenuItem) {
	for _, pageId := range pageIds {
		arr := getAllChildrenMenuByChildrenId(menusMap, pageId)
		if len(arr) > 0 {
			data = append(data, arr...)
		}
	}
	return data
}

func getAllFatherMenuByChildrenId(menusMap map[int32]*model.AdminMenu, parentId int32) (data []*admin_proto.MenuItem) {
	if menu, ok := menusMap[parentId]; ok {
		data = append(data, &admin_proto.MenuItem{
			Key:       menu.Key,
			Path:      menu.Path,
			Name:      menu.Name,
			Icon:      menu.Icon,
			Component: menu.Component,
		})
		arr := getAllFatherMenuByChildrenId(menusMap, menu.ParentID)
		if len(arr) > 0 {
			data = append(data, arr...)
		}
	}
	return data
}

func getAllChildrenMenuByChildrenId(menusMap map[int32]*model.AdminMenu, parentId int32) (data []*admin_proto.MenuItem) {
	for menuId, menu := range menusMap {
		if menu.ParentID == parentId {
			data = append(data, &admin_proto.MenuItem{
				Key:       menu.Key,
				Path:      menu.Path,
				Name:      menu.Name,
				Icon:      menu.Icon,
				Component: menu.Component,
			})

			tmpIds := make(map[int32]*model.AdminMenu)
			for _, item := range menusMap {
				if item.ParentID != parentId {
					tmpIds[item.ID] = item
				}
			}
			arr := getAllChildrenMenuByChildrenId(menusMap, menuId)
			if len(arr) > 0 {
				data = append(data, arr...)
			}
		}
	}
	return data
}
