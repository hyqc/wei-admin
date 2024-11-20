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
	List(ctx *gin.Context, params *admin_proto.ReqPermissionList) (*admin_proto.RespPermissionListData, error)
	Add(ctx *gin.Context, params *admin_proto.ReqPermissionAdd) error
	Info(ctx *gin.Context, params *admin_proto.ReqPermissionInfo) (*admin_proto.PermissionInfo, error)
	Edit(ctx *gin.Context, params *admin_proto.ReqPermissionEdit) error
	Enable(ctx *gin.Context, params *admin_proto.ReqPermissionEnable) error
	Delete(ctx *gin.Context, params *admin_proto.ReqPermissionDelete) error
	BindAPI(ctx *gin.Context, params *admin_proto.ReqPermissionBindApis) error
	AddMenuPermissions(ctx *gin.Context, params *admin_proto.ReqPermissionBindMenu) error
}

func newAdminPermissionLogic() IAdminPermissionLogic {
	return &AdminPermissionLogic{}
}

func (a *AdminPermissionLogic) List(ctx *gin.Context, params *admin_proto.ReqPermissionList) (*admin_proto.RespPermissionListData, error) {
	total, list, err := dao.H.AdminPermission.List(ctx, params)
	if err != nil {
		return nil, err
	}

	data := &admin_proto.RespPermissionListData{
		List:  make([]*admin_proto.PermissionListItem, 0),
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
		data.List = a.handleListData(list, apisMap, menusMap)
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

func (a *AdminPermissionLogic) Add(ctx *gin.Context, params *admin_proto.ReqPermissionAdd) error {
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

func (a *AdminPermissionLogic) Edit(ctx *gin.Context, params *admin_proto.ReqPermissionEdit) error {
	data := &model.AdminPermission{
		ID:        params.Id,
		MenuID:    params.MenuId,
		Key:       params.Key,
		Name:      params.Name,
		Type:      params.Type,
		Describe:  params.Describe,
		IsEnabled: params.Enabled,
	}
	if err := dao.H.AdminPermission.Update(ctx, data); err != nil {
		if strings.Contains(err.Error(), "uk_key") {
			return code.NewCodeError(code_proto.ErrorCode_AdminPermissionKeyExist, err)
		}
		return err
	}
	return nil
}

func (a *AdminPermissionLogic) Info(ctx *gin.Context, params *admin_proto.ReqPermissionInfo) (*admin_proto.PermissionInfo, error) {
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

func (a *AdminPermissionLogic) Enable(ctx *gin.Context, params *admin_proto.ReqPermissionEnable) error {
	info, err := dao.H.AdminPermission.Info(ctx, params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	if info.IsEnabled == params.Enabled {
		return nil
	}
	return dao.H.AdminPermission.Enable(ctx, params.Id, params.Enabled)
}

func (a *AdminPermissionLogic) Delete(ctx *gin.Context, params *admin_proto.ReqPermissionDelete) error {
	info, err := dao.H.AdminPermission.Info(ctx, params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	if info.IsEnabled {
		return code.NewCodeError(code_proto.ErrorCode_RecordNValidCanNotDeleted, nil)
	}
	return dao.H.AdminPermission.Delete(ctx, params.Id)
}

func (a *AdminPermissionLogic) BindAPI(ctx *gin.Context, params *admin_proto.ReqPermissionBindApis) error {
	_, err := dao.H.AdminPermission.Info(ctx, params.PermissionId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	apisList, err := dao.H.AdminAPI.FindByIds(ctx, params.ApiIds, true)
	if err != nil {
		return err
	}
	if len(apisList) == 0 {
		return code.NewCodeError(code_proto.ErrorCode_AdminApiNotExist, err)
	}
	permissionAps := make([]*model.AdminPermissionAPI, 0, len(apisList))
	for _, item := range apisList {
		permissionAps = append(permissionAps, &model.AdminPermissionAPI{
			PermissionID: params.PermissionId,
			APIID:        item.ID,
		})
	}
	return dao.H.AdminPermission.BindApis(ctx, params.PermissionId, permissionAps)
}

func (a *AdminPermissionLogic) AddMenuPermissions(ctx *gin.Context, params *admin_proto.ReqPermissionBindMenu) error {
	_, err := dao.H.AdminMenu.FindById(ctx, params.MenuId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_AdminMenuNotExist, err)
		}
		return err
	}
	data := make([]*model.AdminPermission, 0, len(params.Permissions))
	for _, item := range params.Permissions {
		if item.Key == "" {
			return code.NewCodeError(code_proto.ErrorCode_AdminPermissionKeyNeed, nil)
		}
		if item.Name == "" {
			return code.NewCodeError(code_proto.ErrorCode_AdminPermissionNameNeed, nil)
		}
		if dao.GetAdminPermissionTypeText(dao.AdminPermissionType(item.Type)) == "" {
			return code.NewCodeError(code_proto.ErrorCode_AdminPermissionTypeInvalid, nil)
		}
		data = append(data, &model.AdminPermission{
			MenuID:    params.MenuId,
			Key:       item.Key,
			Name:      item.Name,
			Type:      item.Type,
			Describe:  item.Describe,
			IsEnabled: item.Enabled,
		})
	}
	return dao.H.AdminPermission.BatchAddPermissions(ctx, data)
}
