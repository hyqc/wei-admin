package logic

import (
	"admin/app/admin/dao"
	"admin/app/gen/model"
	"admin/code"
	"admin/pkg/utils"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"github.com/gin-gonic/gin"
	"strings"
)

type AdminAPILogic struct {
	db *dao.AdminAPI
}

func NewAdminAPILogic() *AdminAPILogic {
	return &AdminAPILogic{
		db: adminAPIDao,
	}
}

func (a *AdminAPILogic) List(ctx *gin.Context, params *admin_proto.ApiListReq) (data *admin_proto.ApiListResp, err error) {
	total, rows, err := a.db.FindList(ctx, params)
	if err != nil {
		return nil, err
	}
	data = &admin_proto.ApiListResp{}
	data.Total = total
	data.Rows, err = a.ListHandle(rows)
	return data, err
}

func (a *AdminAPILogic) ListHandle(list []*model.AdminAPI) (rows []*admin_proto.ApiListItem, err error) {
	for _, item := range list {
		data, err := a.ListItemHandle(item)
		if err != nil {
			return nil, err
		}
		rows = append(rows, data)
	}
	return rows, nil
}

func (a *AdminAPILogic) ListItemHandle(item *model.AdminAPI) (data *admin_proto.ApiListItem, err error) {
	data = &admin_proto.ApiListItem{}
	err = utils.BeanCopy(data, item)
	if err != nil {
		return nil, err
	}
	data.Enabled = item.IsEnabled
	data.CreatedAt = item.CreatedAt.Unix()
	data.UpdatedAt = item.UpdatedAt.Unix()
	return data, nil
}

func (a *AdminAPILogic) AllValid(ctx *gin.Context) (list []*admin_proto.ApiListItem, err error) {
	data, err := a.db.FindAllValid(ctx)
	if err != nil {
		return nil, err
	}
	list, err = a.ListHandle(data)
	return list, err
}

func (a *AdminAPILogic) Add(ctx *gin.Context, params *admin_proto.ApiAddReq) (err error) {
	data := &model.AdminAPI{
		Path:      params.Path,
		Key:       params.Key,
		Name:      params.Name,
		Describe:  params.Describe,
		IsEnabled: params.Enabled,
	}
	if err := a.db.Create(ctx, data); err != nil {
		msg := err.Error()
		if strings.Contains(msg, "uk_name") {
			return code.NewCodeError(code_proto.ErrorCode_AdminApiNameExist, err)
		}
		if strings.Contains(msg, "uk_path") {
			return code.NewCodeError(code_proto.ErrorCode_AdminApiPathExist, err)
		}
		if strings.Contains(msg, "uk_key") {
			return code.NewCodeError(code_proto.ErrorCode_AdminApiKeyExist, err)
		}
		return err
	}
	return nil
}
