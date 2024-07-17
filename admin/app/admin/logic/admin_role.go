package logic

import (
	"admin/app/admin/dao"
	"admin/app/gen/model"
	"admin/code"
	"admin/constant"
	"admin/pkg/utils"
	"admin/pkg/utils/array"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type AdminRoleLogic struct {
}

type IAdminRoleLogic interface {
	List(ctx *gin.Context, params *admin_proto.RoleListReq) (*admin_proto.RoleListRespData, error)
	Add(ctx *gin.Context, params *admin_proto.RoleAddReq) error
	Info(ctx *gin.Context, params *admin_proto.RoleInfoReq) (*admin_proto.RoleInfoRespData, error)
	Edit(ctx *gin.Context, params *admin_proto.RoleEditReq) error
	Enable(ctx *gin.Context, params *admin_proto.RoleEnableReq) error
	Delete(ctx *gin.Context, params *admin_proto.RoleDeleteReq) error
	RolePermissions(ctx *gin.Context, params *admin_proto.RolePermissionsReq) ([]*admin_proto.RolePermissionItem, error)
	RoleBindPermissions(ctx *gin.Context, params *admin_proto.RoleBindPermissionsReq) error
}

func newAdminRoleLogic() IAdminRoleLogic {
	return &AdminRoleLogic{}
}

func (a *AdminRoleLogic) List(ctx *gin.Context, params *admin_proto.RoleListReq) (*admin_proto.RoleListRespData, error) {
	total, rows, err := dao.H.AdminRole.List(ctx, params)
	if err != nil {
		return nil, err
	}
	data := &admin_proto.RoleListRespData{}
	data.Total = total
	data.Rows, err = a.HandleListData(rows)
	return data, err
}

func (a *AdminRoleLogic) HandleListData(list []*model.AdminRole) (rows []*admin_proto.RoleItem, err error) {
	for _, item := range list {
		data, err := a.HandleItemData(item)
		if err != nil {
			return nil, err
		}
		rows = append(rows, data)
	}
	return rows, nil
}

func (a *AdminRoleLogic) HandleItemData(item *model.AdminRole) (data *admin_proto.RoleItem, err error) {
	data = &admin_proto.RoleItem{}
	err = utils.BeanCopy(data, item)
	if err != nil {
		return nil, err
	}
	data.Enabled = item.IsEnabled
	data.CreatedAt = item.CreatedAt.Format(time.DateTime)
	data.UpdatedAt = item.UpdatedAt.Format(time.DateTime)
	return data, nil
}

func (a *AdminRoleLogic) Add(ctx *gin.Context, params *admin_proto.RoleAddReq) error {
	adminId := constant.GetCustomClaims(ctx).AdminID
	data := &model.AdminRole{
		Name:          params.Name,
		Describe:      params.Describe,
		ModifyAdminID: adminId,
		CreateAdminID: adminId,
		IsEnabled:     params.Enabled,
	}
	return dao.H.AdminRole.Create(ctx, data)
}

func (a *AdminRoleLogic) Info(ctx *gin.Context, params *admin_proto.RoleInfoReq) (*admin_proto.RoleInfoRespData, error) {
	adminInfo, err := dao.H.AdminRole.Info(ctx, params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return nil, err
	}
	return &admin_proto.RoleInfoRespData{
		Id:              adminInfo.ID,
		Name:            adminInfo.Name,
		Describe:        adminInfo.Describe,
		Enabled:         adminInfo.Enabled,
		CreateAdminId:   adminInfo.CreateAdminId,
		CreateAdminName: adminInfo.CreateAdminName,
		ModifyAdminId:   adminInfo.ModifyAdminId,
		ModifyAdminName: adminInfo.ModifyAdminName,
		CreatedAt:       adminInfo.CreatedAt.Format(time.DateTime),
		UpdatedAt:       adminInfo.UpdatedAt.Format(time.DateTime),
	}, nil
}

func (a *AdminRoleLogic) Edit(ctx *gin.Context, params *admin_proto.RoleEditReq) error {
	adminInfo, err := dao.H.AdminRole.FindById(ctx, params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.NewCodeError(code_proto.ErrorCode_RecordNotExist, err)
		}
		return err
	}
	adminInfo.Name = params.Name
	adminInfo.Describe = params.Describe
	adminInfo.IsEnabled = params.Enabled
	return dao.H.AdminRole.Update(ctx, adminInfo)
}

func (a *AdminRoleLogic) Enable(ctx *gin.Context, params *admin_proto.RoleEnableReq) error {
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

func (a *AdminRoleLogic) Delete(ctx *gin.Context, params *admin_proto.RoleDeleteReq) error {
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

func (a *AdminRoleLogic) RolePermissions(ctx *gin.Context, params *admin_proto.RolePermissionsReq) (list []*admin_proto.RolePermissionItem, err error) {
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

func (a *AdminRoleLogic) RoleBindPermissions(ctx *gin.Context, params *admin_proto.RoleBindPermissionsReq) error {
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
