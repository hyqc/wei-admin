package logic

import (
	"admin/app/admin/dao"
	"admin/app/gen/model"
	"admin/code"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"time"
)

type AdminPermissionLogic struct {
}

type IAdminPermissionLogic interface {
	List(ctx *gin.Context, params *admin_proto.PermissionListReq) (*admin_proto.PermissionListRespData, error)
	Add(ctx *gin.Context, params *admin_proto.PermissionAddReq) error
	Info(ctx *gin.Context, params *admin_proto.PermissionInfoReq) (*admin_proto.PermissionInfo, error)
}

func newAdminPermissionLogic() IAdminPermissionLogic {
	return &AdminPermissionLogic{}
}

func (a *AdminPermissionLogic) List(ctx *gin.Context, params *admin_proto.PermissionListReq) (*admin_proto.PermissionListRespData, error) {
	total, list, err := dao.H.AdminPermission.List(ctx, params)
	if err != nil {
		return nil, err
	}

	data := &admin_proto.PermissionListRespData{
		Rows:  make([]*admin_proto.PermissionListItem, 0),
		Total: total,
	}
	if len(list) > 0 {
		permissionIds := make([]int32, 0, len(list))
		permissionIdsMap := make(map[int32]struct{})
		menuIds := make([]int32, 0, len(list))
		menuIdsMap := make(map[int32]struct{})
		for _, item := range list {
			if _, ok := permissionIdsMap[item.ID]; !ok {
				permissionIds = append(permissionIds, item.ID)
				permissionIdsMap[item.ID] = struct{}{}
			}
			if _, ok := menuIdsMap[item.MenuID]; !ok {
				menuIds = append(menuIds, item.MenuID)
				menuIdsMap[item.MenuID] = struct{}{}
			}
		}
		apisMap := make(map[int32][]*admin_proto.ApiItem)
		if len(permissionIds) > 0 {
			// 根据权限id获取api列表
			apiList, err := dao.H.AdminAPI.FindByPermissionIds(ctx, permissionIds)
			if err != nil {
				return nil, err
			}

			for _, item := range apiList {
				if _, ok := apisMap[item.PermissionId]; !ok {
					apisMap[item.PermissionId] = make([]*admin_proto.ApiItem, 0)
				}
				apisMap[item.PermissionId] = append(apisMap[item.PermissionId], item)
			}
		}

		menusMap := make(map[int32]*model.AdminMenu)
		if len(menuIds) > 0 {
			// 根据菜单id获取菜单列表
			menusList, err := dao.H.AdminMenu.FindByIds(ctx, menuIds)
			if err != nil {
				return nil, err
			}
			for _, item := range menusList {
				menusMap[item.ID] = item
			}
		}
		data.Rows = a.handleListData(list, apisMap, menusMap)
	}
	return data, err
}

func (a *AdminPermissionLogic) handleListData(list []*model.AdminPermission, apisMap map[int32][]*admin_proto.ApiItem, menusMap map[int32]*model.AdminMenu) (data []*admin_proto.PermissionListItem) {
	for _, item := range list {
		data = append(data, a.handleListItemData(item, apisMap, menusMap))
	}
	return data
}

func (a *AdminPermissionLogic) handleListItemData(item *model.AdminPermission, apisMap map[int32][]*admin_proto.ApiItem, menusMap map[int32]*model.AdminMenu) *admin_proto.PermissionListItem {
	tmp := &admin_proto.PermissionListItem{
		Id:        item.ID,
		MenuId:    item.MenuID,
		Key:       item.Key,
		Name:      item.Name,
		Describe:  item.Describe,
		Type:      item.Type,
		TypeText:  dao.GetAdminPermissionTypeText(dao.AdminPermissionType(item.Type)),
		Enabled:   item.IsEnabled,
		CreatedAt: item.CreatedAt.Format(time.DateTime),
		UpdatedAt: item.UpdatedAt.Format(time.DateTime),
	}
	if _, ok := menusMap[item.MenuID]; ok {
		tmp.MenuName = menusMap[item.MenuID].Name
		tmp.MenuPath = menusMap[item.MenuID].Path
	}
	if _, ok := apisMap[item.ID]; ok {
		tmp.Apis = apisMap[item.ID]
	}
	return tmp
}

func (a *AdminPermissionLogic) Add(ctx *gin.Context, params *admin_proto.PermissionAddReq) error {
	data := &model.AdminPermission{
		MenuID:    params.MenuId,
		Key:       params.Key,
		Name:      params.Name,
		Type:      params.Type,
		Describe:  params.Describe,
		IsEnabled: params.Enabled,
	}
	if err := dao.H.AdminPermission.Create(ctx, data); err != nil {
		msg := err.Error()
		if strings.Contains(msg, "uk_key") {
			return code.NewCodeError(code_proto.ErrorCode_AdminPermissionKeyExist, err)
		}
		if strings.Contains(msg, "uk_permission") {
			return code.NewCodeError(code_proto.ErrorCode_AdminPermissionExist, err)
		}
		return err
	}
	return nil
}

func (a *AdminPermissionLogic) Info(ctx *gin.Context, params *admin_proto.PermissionInfoReq) (*admin_proto.PermissionInfo, error) {
	data, err := dao.H.AdminPermission.FindPermissionMenuInfoById(ctx, params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
	}
	apiList, err := dao.H.AdminAPI.FindByPermissionIds(ctx, []int32{data.ID})
	if err != nil {
		return nil, err
	}
	return &admin_proto.PermissionInfo{
		Id:        data.ID,
		MenuId:    data.MenuID,
		MenuName:  data.MenuName,
		MenuPath:  data.MenuPath,
		Key:       data.Key,
		Name:      data.Name,
		Describe:  data.Describe,
		Type:      data.Type,
		Enabled:   data.Enabled,
		Apis:      apiList,
		CreatedAt: data.CreatedAt.Format(time.DateTime),
		UpdatedAt: data.UpdatedAt.Format(time.DateTime),
	}, nil
}
