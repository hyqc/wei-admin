package logic

import (
	"admin/app/admin/dao"
	model2 "admin/app/admin/gen/model"
	"admin/code"
	"admin/pkg/utils"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

type AdminPermissionLogic struct {
}

type IAdminPermissionLogic interface {
	List(ctx *gin.Context, params *admin_proto.ReqAdminPermissionList) (*admin_proto.RespAdminPermissionListData, error)
	Add(ctx *gin.Context, params *admin_proto.ReqAdminPermissionAdd) error
	Info(ctx *gin.Context, params *admin_proto.ReqAdminPermissionInfo) (*admin_proto.AdminPermissionInfo, error)
	Edit(ctx *gin.Context, params *admin_proto.ReqAdminPermissionEdit) error
	Enable(ctx *gin.Context, params *admin_proto.ReqAdminPermissionEnable) error
	Delete(ctx *gin.Context, params *admin_proto.ReqAdminPermissionDelete) error
	BindAPI(ctx *gin.Context, params *admin_proto.ReqAdminPermissionBindApis) error
	UnBindAPI(ctx *gin.Context, params *admin_proto.ReqAdminPermissionUnBindApi) error
	AddMenuPermissions(ctx *gin.Context, params *admin_proto.ReqAdminPermissionBindMenu) error
}

func newAdminPermissionLogic() IAdminPermissionLogic {
	return &AdminPermissionLogic{}
}

func (a *AdminPermissionLogic) List(ctx *gin.Context, params *admin_proto.ReqAdminPermissionList) (*admin_proto.RespAdminPermissionListData, error) {
	total, list, err := dao.H.AdminPermission.List(ctx, params)
	if err != nil {
		return nil, err
	}

	data := &admin_proto.RespAdminPermissionListData{
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
		apisMap := make(map[int32][]*admin_proto.AdminApiItem)
		if len(permissionIds) > 0 {
			// 根据权限id获取api列表
			apiList, err := dao.H.AdminAPI.FindByPermissionIds(ctx, permissionIds)
			if err != nil {
				return nil, err
			}

			for _, item := range apiList {
				if _, ok := apisMap[item.PermissionId]; !ok {
					apisMap[item.PermissionId] = make([]*admin_proto.AdminApiItem, 0)
				}
				apisMap[item.PermissionId] = append(apisMap[item.PermissionId], &admin_proto.AdminApiItem{
					Id:           item.ID,
					Path:         item.Path,
					Key:          item.Key,
					Name:         item.Name,
					IsEnabled:    item.IsEnabled,
					PermissionId: item.PermissionId,
					Describe:     item.Describe,
				})
			}
		}

		menusMap := make(map[int32]*model2.AdminMenu)
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

func (a *AdminPermissionLogic) handleListData(list []*model2.AdminPermission, apisMap map[int32][]*admin_proto.AdminApiItem, menusMap map[int32]*model2.AdminMenu) (data []*admin_proto.PermissionListItem) {
	for _, item := range list {
		data = append(data, a.handleListItemData(item, apisMap, menusMap))
	}
	return data
}

func (a *AdminPermissionLogic) handleListItemData(item *model2.AdminPermission, apisMap map[int32][]*admin_proto.AdminApiItem, menusMap map[int32]*model2.AdminMenu) *admin_proto.PermissionListItem {
	tmp := &admin_proto.PermissionListItem{
		Id:        item.ID,
		MenuId:    item.MenuID,
		Key:       item.Key,
		Name:      item.Name,
		Describe:  item.Describe,
		Type:      item.Type,
		TypeText:  dao.GetAdminPermissionTypeText(dao.AdminPermissionType(item.Type)),
		IsEnabled: item.IsEnabled,
		CreatedAt: utils.HandleTime2String(item.CreatedAt),
		UpdatedAt: utils.HandleTime2String(item.UpdatedAt),
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

func (a *AdminPermissionLogic) Add(ctx *gin.Context, params *admin_proto.ReqAdminPermissionAdd) error {
	data := &model2.AdminPermission{
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

func (a *AdminPermissionLogic) Edit(ctx *gin.Context, params *admin_proto.ReqAdminPermissionEdit) error {
	data := &model2.AdminPermission{
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

func (a *AdminPermissionLogic) Info(ctx *gin.Context, params *admin_proto.ReqAdminPermissionInfo) (*admin_proto.AdminPermissionInfo, error) {
	data, err := dao.H.AdminPermission.FindPermissionMenuInfoById(ctx, params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
	}
	apiListTmp, err := dao.H.AdminAPI.FindByPermissionIds(ctx, []int32{data.ID})
	if err != nil {
		return nil, err
	}
	apiList := make([]*admin_proto.AdminApiItem, 0, len(apiListTmp))
	for _, item := range apiListTmp {
		apiList = append(apiList, &admin_proto.AdminApiItem{
			Id:        item.ID,
			Path:      item.Path,
			Key:       item.Key,
			Name:      item.Name,
			IsEnabled: item.IsEnabled,
			Describe:  item.Describe,
		})
	}
	return &admin_proto.AdminPermissionInfo{
		Id:        data.ID,
		MenuId:    data.MenuID,
		MenuName:  data.MenuName,
		MenuPath:  data.MenuPath,
		Key:       data.Key,
		Name:      data.Name,
		Describe:  data.Describe,
		Type:      data.Type,
		IsEnabled: data.Enabled,
		Apis:      apiList,
		CreatedAt: utils.HandleTimePointer2String(data.CreatedAt),
		UpdatedAt: utils.HandleTimePointer2String(data.UpdatedAt),
	}, nil
}

func (a *AdminPermissionLogic) Enable(ctx *gin.Context, params *admin_proto.ReqAdminPermissionEnable) error {
	info, err := dao.H.AdminPermission.Info(ctx, params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	if info.IsEnabled == params.IsEnabled {
		return nil
	}
	return dao.H.AdminPermission.Enable(ctx, params.Id, params.IsEnabled)
}

func (a *AdminPermissionLogic) Delete(ctx *gin.Context, params *admin_proto.ReqAdminPermissionDelete) error {
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

func (a *AdminPermissionLogic) BindAPI(ctx *gin.Context, params *admin_proto.ReqAdminPermissionBindApis) error {
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
	permissionAps := make([]*model2.AdminPermissionAPI, 0, len(apisList))
	for _, item := range apisList {
		permissionAps = append(permissionAps, &model2.AdminPermissionAPI{
			PermissionID: params.PermissionId,
			APIID:        item.ID,
		})
	}
	return dao.H.AdminPermission.BindApis(ctx, params.PermissionId, permissionAps)
}
func (a *AdminPermissionLogic) UnBindAPI(ctx *gin.Context, params *admin_proto.ReqAdminPermissionUnBindApi) error {
	_, err := dao.H.AdminPermission.Info(ctx, params.PermissionId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	apisList, err := dao.H.AdminAPI.FindByIds(ctx, []int32{params.ApiId}, true)
	if err != nil {
		return err
	}
	if len(apisList) == 0 {
		return code.NewCodeError(code_proto.ErrorCode_AdminApiNotExist, err)
	}
	permissionAps := make([]*model2.AdminPermissionAPI, 0, len(apisList))
	for _, item := range apisList {
		permissionAps = append(permissionAps, &model2.AdminPermissionAPI{
			PermissionID: params.PermissionId,
			APIID:        item.ID,
		})
	}
	return dao.H.AdminPermission.UnBindApi(ctx, params.PermissionId, params.ApiId)
}

func (a *AdminPermissionLogic) AddMenuPermissions(ctx *gin.Context, params *admin_proto.ReqAdminPermissionBindMenu) error {
	_, err := dao.H.AdminMenu.FindById(ctx, params.MenuId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_AdminMenuNotExist, err)
		}
		return err
	}
	saveData := make([]*model2.AdminPermission, 0, len(params.Permissions))
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
		saveData = append(saveData, &model2.AdminPermission{
			ID:        item.Id,
			MenuID:    params.MenuId,
			Key:       item.Key,
			Name:      item.Name,
			Type:      item.Type,
			Describe:  item.Describe,
			IsEnabled: item.Enabled,
		})
	}
	return dao.H.AdminPermission.BatchAddPermissions(ctx, saveData)
}
