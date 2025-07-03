package logic

import (
	"admin/app/admin/dao"
	model "admin/app/admin/gen/model"
	"admin/app/common"
	"admin/code"
	"admin/constant"
	"admin/pkg/utils"
	"admin/pkg/utils/array"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminRoleLogic struct {
}

type IAdminRoleLogic interface {
	List(ctx *gin.Context, params *admin_proto.ReqAdminRoleList) (*admin_proto.RespAdminRoleListData, error)
	Add(ctx *gin.Context, params *admin_proto.ReqAdminRoleAdd) error
	Info(ctx *gin.Context, params *admin_proto.ReqAdminRoleInfo) (*admin_proto.RespAdminRoleInfoData, error)
	Edit(ctx *gin.Context, params *admin_proto.ReqAdminRoleEdit) error
	Enable(ctx *gin.Context, params *admin_proto.ReqAdminRoleEnable) error
	Delete(ctx *gin.Context, params *admin_proto.ReqAdminRoleDelete) error
	RolePermissions(ctx *gin.Context, params *admin_proto.ReqAdminRolePermissions) ([]*admin_proto.RolePermissionItem, error)
	RoleBindPermissions(ctx *gin.Context, params *admin_proto.ReqAdminRoleBindPermissions) error
	All(ctx *gin.Context) ([]*admin_proto.RoleItem, error)
}

func newAdminRoleLogic() IAdminRoleLogic {
	return &AdminRoleLogic{}
}

func (a *AdminRoleLogic) List(ctx *gin.Context, params *admin_proto.ReqAdminRoleList) (*admin_proto.RespAdminRoleListData, error) {
	total, rows, err := dao.H.AdminRole.List(ctx, params)
	if err != nil {
		return nil, err
	}
	data := &admin_proto.RespAdminRoleListData{}
	data.Total = total
	data.List, err = a.HandleListData(rows)
	return data, err
}

func (a *AdminRoleLogic) HandleListData(rows []*model.AdminRole) (list []*admin_proto.RoleItem, err error) {
	for _, item := range rows {
		data, err := a.HandleItemData(item)
		if err != nil {
			return nil, err
		}
		list = append(list, data)
	}
	return list, nil
}

func (a *AdminRoleLogic) HandleItemData(item *model.AdminRole) (data *admin_proto.RoleItem, err error) {
	data = &admin_proto.RoleItem{}
	err = utils.BeanCopy(data, item)
	if err != nil {
		return nil, err
	}
	data.Id = item.ID
	data.CreatedAt = utils.HandleTime2String(item.CreatedAt)
	data.UpdatedAt = utils.HandleTime2String(item.UpdatedAt)
	return data, nil
}

func (a *AdminRoleLogic) Add(ctx *gin.Context, params *admin_proto.ReqAdminRoleAdd) error {
	adminId := constant.GetCustomClaims(ctx).AccountId
	data := &model.AdminRole{
		Name:          params.Name,
		Describe:      params.Describe,
		ModifyAdminID: adminId,
		CreateAdminID: adminId,
		IsEnabled:     params.Enabled,
	}
	return dao.H.AdminRole.Create(ctx, data)
}

func (a *AdminRoleLogic) Info(ctx *gin.Context, params *admin_proto.ReqAdminRoleInfo) (*admin_proto.RespAdminRoleInfoData, error) {
	adminInfo, err := dao.H.AdminRole.Info(ctx, params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return nil, err
	}
	permissionsList, err := dao.H.AdminRole.FindAdminRolePermissionByRoleId(ctx, params.Id)
	if err != nil {
		return nil, err
	}
	rest := &admin_proto.RespAdminRoleInfoData{
		Id:              adminInfo.ID,
		Name:            adminInfo.Name,
		Describe:        adminInfo.Describe,
		IsEnabled:       adminInfo.IsEnabled,
		CreateAdminId:   adminInfo.CreateAdminId,
		CreateAdminName: adminInfo.CreateAdminName,
		ModifyAdminId:   adminInfo.ModifyAdminId,
		ModifyAdminName: adminInfo.ModifyAdminName,
		CreatedAt:       utils.HandleTime2String(adminInfo.CreatedAt),
		UpdatedAt:       utils.HandleTime2String(adminInfo.UpdatedAt),
		Permissions:     make([]*admin_proto.RolePermissionItem, 0),
		PermissionIds:   make([]int32, 0),
	}
	for _, item := range permissionsList {
		enumItem, ok := common.AdminPermissionEnumMap[item.PermissionType]
		if !ok {
			return nil, fmt.Errorf("配置错误")
		}
		rest.PermissionIds = append(rest.PermissionIds, item.PermissionID)
		rest.Permissions = append(rest.Permissions, &admin_proto.RolePermissionItem{
			RoleId:             adminInfo.ID,
			PermissionId:       item.PermissionID,
			PermissionName:     item.PermissionName,
			PermissionKey:      item.PermissionKey,
			PermissionType:     item.PermissionType,
			PermissionTypeText: enumItem.Name,
		})
	}
	rest.PermissionIds = array.Deduplicate(rest.PermissionIds, true, true)
	return rest, nil
}

func (a *AdminRoleLogic) Edit(ctx *gin.Context, params *admin_proto.ReqAdminRoleEdit) error {
	adminInfo, err := dao.H.AdminRole.FindById(ctx, params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	adminInfo.Name = params.Name
	adminInfo.Describe = params.Describe
	adminInfo.IsEnabled = params.IsEnabled
	return dao.H.AdminRole.Update(ctx, adminInfo)
}

func (a *AdminRoleLogic) Enable(ctx *gin.Context, params *admin_proto.ReqAdminRoleEnable) error {
	info, err := dao.H.AdminRole.FindById(ctx, params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	if info.IsEnabled == params.Enabled {
		return nil
	}
	return dao.H.AdminRole.Enable(ctx, params.Id, params.Enabled)
}

func (a *AdminRoleLogic) Delete(ctx *gin.Context, params *admin_proto.ReqAdminRoleDelete) error {
	info, err := dao.H.AdminRole.FindById(ctx, params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	if info.IsEnabled {
		return code.NewCodeError(code_proto.ErrorCode_RecordNValidCanNotDeleted, nil)
	}
	return dao.H.AdminRole.Delete(ctx, params.Id)
}

func (a *AdminRoleLogic) RolePermissions(ctx *gin.Context, params *admin_proto.ReqAdminRolePermissions) (list []*admin_proto.RolePermissionItem, err error) {
	data, err := dao.H.AdminRole.FindAdminRolePermissionByRoleId(ctx, params.Id)
	if err != nil {
		return nil, err
	}
	for _, item := range data {
		list = append(list, &admin_proto.RolePermissionItem{
			PermissionId:       item.PermissionID,
			PermissionName:     item.PermissionName,
			PermissionKey:      item.PermissionKey,
			PermissionType:     item.PermissionType,
			PermissionTypeText: dao.AdminPermissionTypeTextMap[dao.AdminPermissionType(item.PermissionType)],
			RoleId:             item.RoleID,
		})
	}
	return list, nil
}

func (a *AdminRoleLogic) RoleBindPermissions(ctx *gin.Context, params *admin_proto.ReqAdminRoleBindPermissions) error {
	params.PermissionIds = array.Deduplicate(params.PermissionIds, true, true)
	if len(params.PermissionIds) == 0 {
		return code.NewCode(code_proto.ErrorCode_RequestParamsInvalid)
	}
	_, err := dao.H.AdminRole.FindById(ctx, params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	permissions, err := dao.H.AdminPermission.FindByIds(ctx, params.PermissionIds)
	if err != nil {
		return err
	}
	if len(permissions) == 0 {
		return code.NewCode(code_proto.ErrorCode_AdminPermissionExist)
	}
	data := make([]*model.AdminRolePermission, 0, len(params.PermissionIds))
	for _, item := range params.PermissionIds {
		data = append(data, &model.AdminRolePermission{
			RoleID:       params.Id,
			PermissionID: item,
		})
	}
	return dao.H.AdminRole.BindPermissions(ctx, params.Id, data)
}

func (a *AdminRoleLogic) All(ctx *gin.Context) (list []*admin_proto.RoleItem, err error) {
	data, err := dao.H.AdminRole.FindAllValid(ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range data {
		list = append(list, &admin_proto.RoleItem{
			Id:   item.ID,
			Name: item.Name,
		})
	}
	return list, nil
}
