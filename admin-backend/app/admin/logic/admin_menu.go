package logic

import (
	"admin/app/admin/dao"
	"admin/app/admin/gen/model"
	"admin/app/common"
	"admin/code"
	"admin/pkg/utils"
	"admin/pkg/utils/array"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sort"
)

type AdminMenuLogic struct {
}

type IAdminMenuLogic interface {
	List(ctx *gin.Context, params *admin_proto.ReqMenuList) (data *admin_proto.RespMenuListData, err error)
	Tree(ctx *gin.Context) ([]*admin_proto.MenuTreeItem, error)
	Add(ctx *gin.Context, params *admin_proto.ReqMenuAdd) error
	Info(ctx *gin.Context, params *admin_proto.ReqMenuInfo) (*admin_proto.MenuItem, error)
	Edit(ctx *gin.Context, params *admin_proto.ReqMenuEdit) error
	Enable(ctx *gin.Context, params *admin_proto.ReqMenuEnable) error
	Delete(ctx *gin.Context, params *admin_proto.ReqMenuDelete) error
	Permissions(ctx *gin.Context, params *admin_proto.ReqMenuPermissions) (*admin_proto.MenuPermissions, error)
	Pages(ctx *gin.Context, params *admin_proto.ReqMenuPages) (list []*admin_proto.MenuTreeItem, err error)
	AllMode(ctx *gin.Context) (*admin_proto.RespMenuModeData, error)
}

func newAdminMenuLogic() IAdminMenuLogic {
	return &AdminMenuLogic{}
}

func (a *AdminMenuLogic) List(ctx *gin.Context, params *admin_proto.ReqMenuList) (data *admin_proto.RespMenuListData, err error) {
	total, rows, err := dao.H.AdminMenu.FindList(ctx, params)
	if err != nil {
		return nil, err
	}
	data = &admin_proto.RespMenuListData{}
	data.Total = total
	data.List, err = a.handleListData(rows)
	return data, err
}

func (a *AdminMenuLogic) Tree(ctx *gin.Context) ([]*admin_proto.MenuTreeItem, error) {
	allMenus, err := dao.H.AdminMenu.FindAll(ctx)
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

func (a *AdminMenuLogic) Add(ctx *gin.Context, params *admin_proto.ReqMenuAdd) error {
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
	return dao.H.AdminMenu.Create(ctx, data)
}

func (a *AdminMenuLogic) Info(ctx *gin.Context, params *admin_proto.ReqMenuInfo) (*admin_proto.MenuItem, error) {
	data, err := dao.H.AdminMenu.FindById(ctx, params.MenuId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return nil, err
	}
	return a.handleItemData(data)
}

func (a *AdminMenuLogic) Edit(ctx *gin.Context, params *admin_proto.ReqMenuEdit) error {
	info, err := dao.H.AdminMenu.FindById(ctx, params.MenuId)
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
	return dao.H.AdminMenu.Update(ctx, info)
}

func (a *AdminMenuLogic) Enable(ctx *gin.Context, params *admin_proto.ReqMenuEnable) error {
	info, err := dao.H.AdminMenu.FindById(ctx, params.MenuId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	if info.IsEnabled == params.Enabled {
		return nil
	}
	return dao.H.AdminMenu.Enable(ctx, params.MenuId, params.Enabled)
}

func (a *AdminMenuLogic) Delete(ctx *gin.Context, params *admin_proto.ReqMenuDelete) error {
	info, err := dao.H.AdminMenu.FindById(ctx, params.MenuId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	if info.IsEnabled {
		return code.NewCodeError(code_proto.ErrorCode_RecordNValidCanNotDeleted, nil)
	}
	return dao.H.AdminMenu.Delete(ctx, params.MenuId)
}

func (a *AdminMenuLogic) Permissions(ctx *gin.Context, params *admin_proto.ReqMenuPermissions) (*admin_proto.MenuPermissions, error) {
	info, err := dao.H.AdminMenu.FindById(ctx, params.MenuId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return nil, err
	}
	data := &admin_proto.MenuPermissions{
		Menu:        &admin_proto.MenuTreeItem{},
		Permissions: make([]*admin_proto.PermissionApiItem, 0),
	}
	_ = utils.BeanCopy(data.Menu, info)
	permissions, err := dao.H.AdminPermission.FindPermissionsByMenuId(ctx, info.ID)
	if err != nil {
		return nil, err
	}
	if len(permissions) == 0 {
		data.Permissions = common.AdminPermissionEnumList(info.ID, info.Key)
		return data, nil
	}
	// TODO
	permissionIdsMap := make(map[int32]struct{})
	permissionIds := make([]int32, 0, len(permissions))
	for _, item := range permissions {
		if _, ok := permissionIdsMap[item.ID]; !ok {
			permissionIdsMap[item.ID] = struct{}{}
			permissionIds = append(permissionIds, item.ID)
		}

	}
	apiList, err := dao.H.AdminAPI.FindByPermissionIds(ctx, permissionIds)
	if err != nil {
		return nil, err
	}
	apiMap := make(map[int32][]*admin_proto.ApiItem)
	for _, item := range apiList {
		if _, ok := apiMap[item.PermissionId]; !ok {
			apiMap[item.PermissionId] = make([]*admin_proto.ApiItem, 0, len(apiList))
		}
		apiMap[item.PermissionId] = append(apiMap[item.PermissionId], item)
	}
	for _, item := range data.Permissions {
		if val, ok := apiMap[item.Id]; ok {
			item.Apis = val
		}
	}
	return data, nil
}

func (a *AdminMenuLogic) Pages(ctx *gin.Context, params *admin_proto.ReqMenuPages) (list []*admin_proto.MenuTreeItem, err error) {
	//data, err := dao.H.AdminMenu.FindPages(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//if params.All {
	//	list = append(list, &admin_proto.MenuTreeItem{
	//		Level: 0,
	//		Id:    0,
	//		Key:   "All",
	//		Name:  "全部",
	//	})
	//}
	return
}

func (a *AdminMenuLogic) handleListData(rows []*model.AdminMenu) (list []*admin_proto.MenuItem, err error) {
	for _, item := range rows {
		data, err := a.handleItemData(item)
		if err != nil {
			return nil, err
		}
		list = append(list, data)
	}
	return list, nil
}

func (a *AdminMenuLogic) handleItemData(item *model.AdminMenu) (data *admin_proto.MenuItem, err error) {
	data = &admin_proto.MenuItem{}
	err = utils.BeanCopy(data, item)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *AdminMenuLogic) AllMode(ctx *gin.Context) (*admin_proto.RespMenuModeData, error) {
	//获取全部权限
	permissionList, err := dao.H.AdminPermission.FindAdministerPermissions(ctx)
	if err != nil {
		return nil, err
	}
	//获取权限对应的页面
	pagesMap := make(map[int32]*admin_proto.MenuPageItem)
	pageIds := make([]int32, 0, len(permissionList))
	for _, item := range permissionList {
		pageIds = append(pageIds, item.MenuID)
		if _, ok := pagesMap[item.MenuID]; !ok {
			pagesMap[item.MenuID] = &admin_proto.MenuPageItem{
				PageId:      item.MenuID,
				PageName:    "",
				Permissions: make([]*admin_proto.MenuPagePermissions, 0),
			}
		}
		typeName, err := common.GetPermissionTypeName(item.Type)
		if err != nil {
			return nil, err
		}
		pagesMap[item.MenuID].Permissions = append(pagesMap[item.MenuID].Permissions, &admin_proto.MenuPagePermissions{
			PermissionId:       item.ID,
			PermissionName:     item.Name,
			PermissionType:     item.Type,
			PermissionTypeName: typeName,
		})
	}

	pageIds = array.Deduplicate(pageIds, true, true)

	//获取所有页面信息
	menuList, err := dao.H.AdminMenu.FindAllValid(ctx)
	if err != nil {
		return nil, err
	}
	menuMap := make(map[int32]*model.AdminMenu)
	for _, item := range menuList {
		menuMap[item.ID] = item
	}
	for _, page := range pagesMap {
		sort.Slice(page.Permissions, func(i, j int) bool {
			return page.Permissions[i].PermissionType > page.Permissions[j].PermissionType
		})
		if menu, ok := menuMap[page.PageId]; ok {
			page.PageName = menu.Name
		}
	}
	result := &admin_proto.RespMenuModeData{
		Modes: make([]*admin_proto.MenuModeItem, 0),
	}
	//获取所有顶级页面
	menuModeList := getAllTopMenusByPageIds(menuList, pageIds, pagesMap)
	if len(menuModeList) == 0 {
		return result, nil
	}
	// 获取页面的权限
	menuModeMap := make(map[int32]*admin_proto.MenuModeItem)
	// 合并模块
	for _, item := range menuModeList {
		if _, ok := menuModeMap[item.ModelId]; !ok {
			menuModeMap[item.ModelId] = item
		} else {
			menuModeMap[item.ModelId].Pages = append(menuModeMap[item.ModelId].Pages, item.Pages...)
		}
	}
	for _, item := range menuModeMap {
		sort.Slice(item.Pages, func(i, j int) bool {
			return item.Pages[i].PageId < item.Pages[j].PageId
		})
		result.Modes = append(result.Modes, item)
	}
	return result, nil
}

func getAllTopMenusByPageIds(menus []*model.AdminMenu, pageIds []int32, pagesMap map[int32]*admin_proto.MenuPageItem) (list []*admin_proto.MenuModeItem) {
	for _, pageId := range pageIds {
		item := getTopMenuByPageId(menus, pageId, pagesMap[pageId])
		if item != nil {
			list = append(list, item)
		}
	}
	return list
}

// 根据子菜单ID获取全部的祖先菜单
// menus 所有的菜单
// pageId 页面菜单ID
func getTopMenuByPageId(menus []*model.AdminMenu, pageId int32, page *admin_proto.MenuPageItem) *admin_proto.MenuModeItem {
	for _, item := range menus {
		if item.ID == pageId {
			if item.ParentID == 0 {
				rest := &admin_proto.MenuModeItem{
					ModelId:   item.ID,
					ModelName: item.Name,
					Pages:     make([]*admin_proto.MenuPageItem, 0, 1),
				}
				if page != nil {
					rest.Pages = append(rest.Pages, page)
				}
				return rest
			}
			return getTopMenuByPageId(menus, item.ParentID, page)
		}
	}
	return nil
}
