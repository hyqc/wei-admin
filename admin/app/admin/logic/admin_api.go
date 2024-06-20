package logic

import (
	"admin/app/admin/dao"
	"admin/app/gen/model"
	"admin/pkg/utils"
	adminproto "admin/proto"
	"github.com/gin-gonic/gin"
)

type AdminAPILogic struct {
	*dao.AdminAPI
}

func NewAdminAPILogic() *AdminAPILogic {
	return &AdminAPILogic{
		adminAPIDao,
	}
}

func (a *AdminAPILogic) List(ctx *gin.Context, params *adminproto.ApiListReq) (data *adminproto.ApiListResp, err error) {
	total, list, err := a.FindList(ctx, params)
	if err != nil {
		return nil, err
	}
	data = &adminproto.ApiListResp{}
	data.Total = total
	data.Rows, err = a.ListHandle(list)
	return data, err
}

func (a *AdminAPILogic) ListHandle(list []*model.AdminAPI) (rows []*adminproto.ApiListItem, err error) {
	for _, item := range list {
		data, err := a.ListItemHandle(item)
		if err != nil {
			return nil, err
		}
		rows = append(rows, data)
	}
	return rows, nil
}

func (a *AdminAPILogic) ListItemHandle(item *model.AdminAPI) (data *adminproto.ApiListItem, err error) {
	data = &adminproto.ApiListItem{}
	err = utils.BeanCopy(data, item)
	if err != nil {
		return nil, err
	}
	data.Enabled = item.IsEnabled
	data.CreatedAt = item.CreatedAt.Unix()
	data.UpdatedAt = item.UpdatedAt.Unix()
	return data, nil
}
