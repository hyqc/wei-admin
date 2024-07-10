package logic

import (
	"admin/app/admin/dao"
	"admin/proto/admin_proto"
	"github.com/gin-gonic/gin"
)

type AdminPermissionLogic struct {
	db *dao.AdminPermission
}

type IAdminPermissionLogic interface {
	List(ctx *gin.Context, params *admin_proto.PermissionListReq) (*admin_proto.PermissionListRespData, error)
}

func newAdminPermissionLogic() IAdminPermissionLogic {
	return &AdminPermissionLogic{}
}

func (p *AdminPermissionLogic) List(ctx *gin.Context, params *admin_proto.PermissionListReq) (*admin_proto.PermissionListRespData, error) {
	total, _, err := p.db.FindList(ctx, params)
	if err != nil {
		return nil, err
	}
	data := &admin_proto.PermissionListRespData{}
	data.Total = total
	//data.Rows, err = p.handleListData(rows)
	return data, err
}
