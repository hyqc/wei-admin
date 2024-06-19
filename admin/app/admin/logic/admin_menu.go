package logic

import (
	"admin/app/admin/dao"
	"admin/app/gen/model"
	adminproto "admin/proto"
	"context"
)

type AdminMenuLogic struct {
	*dao.AdminMenu
}

func NewAdminMenuLogic() *AdminMenuLogic {
	return &AdminMenuLogic{
		adminMenuDao,
	}
}

func (a *AdminMenuLogic) getMyMenusMap(ctx context.Context, pageIds []int32) ([]*adminproto.MenuItem, error) {
	allMenus, err := a.FindAll(ctx)
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
	menus := make([]*adminproto.MenuItem, 0, len(arr))
	m := make(map[string]struct{}, len(arr))
	for _, item := range arr {
		if _, ok := m[item.Key]; !ok {
			m[item.Key] = struct{}{}
			menus = append(menus, item)
		}
	}
	return menus, nil
}

func getAllRelatedMenusByPageIds(menusMap map[int32]*model.AdminMenu, pageIds []int32) (data []*adminproto.MenuItem) {
	for _, pageId := range pageIds {
		arr := getAllFatherMenuByChildrenId(menusMap, pageId)
		if len(arr) > 0 {
			data = append(data, arr...)
		}
	}
	return data
}

func getAllChildrenPagesByPageIds(menusMap map[int32]*model.AdminMenu, pageIds []int32) (data []*adminproto.MenuItem) {
	for _, pageId := range pageIds {
		arr := getAllChildrenMenuByChildrenId(menusMap, pageId)
		if len(arr) > 0 {
			data = append(data, arr...)
		}
	}
	return data
}

func getAllFatherMenuByChildrenId(menusMap map[int32]*model.AdminMenu, parentId int32) (data []*adminproto.MenuItem) {
	if menu, ok := menusMap[parentId]; ok {
		data = append(data, &adminproto.MenuItem{
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

func getAllChildrenMenuByChildrenId(menusMap map[int32]*model.AdminMenu, parentId int32) (data []*adminproto.MenuItem) {
	for menuId, menu := range menusMap {
		if menu.ParentID == parentId {
			data = append(data, &adminproto.MenuItem{
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
