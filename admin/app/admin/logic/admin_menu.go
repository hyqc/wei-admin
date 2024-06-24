package logic

import (
	"admin/app/admin/dao"
	"admin/app/common"
	"admin/app/gen/model"
	"admin/code"
	"admin/pkg/utils"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminMenuLogic struct {
	db *dao.AdminMenu
}

func NewAdminMenuLogic() *AdminMenuLogic {
	return &AdminMenuLogic{
		db: adminMenuDao,
	}
}

func (a *AdminMenuLogic) List(ctx *gin.Context, params *admin_proto.MenuListReq) (data *admin_proto.MenuListRespData, err error) {
	total, rows, err := a.db.FindList(ctx, params)
	if err != nil {
		return nil, err
	}
	data = &admin_proto.MenuListRespData{}
	data.Total = total
	data.Rows, err = a.HandleListData(rows)
	return data, err
}

func (a *AdminMenuLogic) Tree(ctx *gin.Context) ([]*admin_proto.MenuTreeItem, error) {
	allMenus, err := a.db.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	menusMap := make(map[int32][]*admin_proto.MenuTreeItem)
	for _, item := range allMenus {
		if _, ok := menusMap[item.ParentID]; !ok {
			menusMap[item.ParentID] = make([]*admin_proto.MenuTreeItem, 0)
		}
		tmp := &admin_proto.MenuTreeItem{}
		_ = utils.BeanCopy(tmp, item)
		tmp.Id = item.ID
		tmp.ParentId = item.ParentID
		tmp.Enabled = item.IsEnabled
		tmp.HideInMenu = item.IsHideInMenu
		tmp.HideChildrenInMenu = item.IsHideChildrenInMenu
		tmp.CreateTime = item.CreatedAt.Unix()
		tmp.ModifyTime = item.UpdatedAt.Unix()
		tmp.Children = make([]*admin_proto.MenuTreeItem, 0)
		menusMap[item.ParentID] = append(menusMap[item.ParentID], tmp)
	}
	data := common.MenuTree(menusMap, nil, 0, 1, 4)
	return common.GetMenuTreeWithTop(data), nil
}

func (a *AdminMenuLogic) Add(ctx *gin.Context, params *admin_proto.MenuAddReq) error {
	data := &model.AdminMenu{
		ParentID:             params.ParentId,
		Path:                 params.Path,
		Name:                 params.Name,
		Key:                  params.Key,
		Describe:             params.Describe,
		Icon:                 params.Icon,
		Sort:                 params.Sort,
		Redirect:             params.Redirect,
		Component:            params.Component,
		IsHideInMenu:         params.HideInMenu,
		IsHideChildrenInMenu: params.HideChildrenInMenu,
		IsEnabled:            params.Enabled,
	}
	return a.db.Create(ctx, data)
}

func (a *AdminMenuLogic) Info(ctx *gin.Context, params *admin_proto.MenuInfoReq) (*admin_proto.MenuItem, error) {
	data, err := a.db.FindById(ctx, params.MenuId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return nil, err
	}
	return a.HandleItemData(data)
}

func (a *AdminMenuLogic) Edit(ctx *gin.Context, params *admin_proto.MenuEditReq) error {
	info, err := a.db.FindById(ctx, params.MenuId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	info.ParentID = params.ParentId
	info.Path = params.Path
	info.Name = params.Name
	info.Key = params.Key
	info.Describe = params.Describe
	info.Icon = params.Icon
	info.Sort = params.Sort
	info.Redirect = params.Redirect
	info.IsHideInMenu = params.HideInMenu
	info.IsHideChildrenInMenu = params.HideChildrenInMenu
	info.IsEnabled = params.Enabled
	return a.db.Update(ctx, info)
}

func (a *AdminMenuLogic) Enable(ctx *gin.Context, params *admin_proto.MenuEnableReq) error {
	info, err := a.db.FindById(ctx, params.MenuId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	if info.IsEnabled == params.Enabled {
		return nil
	}
	return a.db.Enable(ctx, params.MenuId, params.Enabled)
}

func (a *AdminMenuLogic) Delete(ctx *gin.Context, params *admin_proto.MenuDeleteReq) error {
	info, err := a.db.FindById(ctx, params.MenuId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	if info.IsEnabled {
		return code.NewCodeError(code_proto.ErrorCode_RecordNValidCanNotDeleted, nil)
	}
	return a.db.Delete(ctx, params.MenuId)
}

func (a *AdminMenuLogic) HandleListData(list []*model.AdminMenu) (rows []*admin_proto.MenuItem, err error) {
	for _, item := range list {
		data, err := a.HandleItemData(item)
		if err != nil {
			return nil, err
		}
		rows = append(rows, data)
	}
	return rows, nil
}

func (a *AdminMenuLogic) HandleItemData(item *model.AdminMenu) (data *admin_proto.MenuItem, err error) {
	data = &admin_proto.MenuItem{}
	err = utils.BeanCopy(data, item)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *AdminMenuLogic) getMyMenusMap(ctx context.Context, pageIds []int32) ([]*admin_proto.MenuItem, error) {
	allMenus, err := a.db.FindAllValid(ctx)
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
