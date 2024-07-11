package logic

import (
	"admin/app/admin/dao"
	"admin/app/gen/model"
	"admin/code"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"github.com/gin-gonic/gin"
	"strings"
)

type AdminPermissionLogic struct {
}

type IAdminPermissionLogic interface {
	List(ctx *gin.Context, params *admin_proto.PermissionListReq) (*admin_proto.PermissionListRespData, error)
	Add(ctx *gin.Context, params *admin_proto.PermissionAddReq) error
}

func newAdminPermissionLogic() IAdminPermissionLogic {
	return &AdminPermissionLogic{}
}

func (p *AdminPermissionLogic) List(ctx *gin.Context, params *admin_proto.PermissionListReq) (*admin_proto.PermissionListRespData, error) {
	total, _, err := dao.H.AdminPermission.FindList(ctx, params)
	if err != nil {
		return nil, err
	}
	data := &admin_proto.PermissionListRespData{}
	data.Total = total
	//data.Rows, err = p.handleListData(rows)
	return data, err
}

func (p *AdminPermissionLogic) Add(ctx *gin.Context, params *admin_proto.PermissionAddReq) error {
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
