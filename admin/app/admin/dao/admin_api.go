package dao

import (
	"admin/app/common"
	"admin/app/gen/model"
	"admin/app/gen/query"
	"admin/proto/admin_proto"
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gen/field"
	"time"
)

type IAdminAPI interface {
	Create(ctx *gin.Context, data *model.AdminAPI) error
	Update(ctx *gin.Context, data *model.AdminAPI) error
	Enable(ctx *gin.Context, id int32, enabled bool) error
	Delete(ctx *gin.Context, id int32) error
	FindList(ctx context.Context, params *admin_proto.ApiListReq) (total int64, list []*model.AdminAPI, err error)
	FindAllValid(ctx context.Context) (list []*model.AdminAPI, err error) // 查找所有有效的接口
	FindById(ctx *gin.Context, id int32) (*model.AdminAPI, error)
	FindByPermissionIds(ctx *gin.Context, ids []int32) (data []*admin_proto.ApiItem, err error)
}

type AdminAPI struct {
}

func newAdminAPI() *AdminAPI {
	return &AdminAPI{}
}

func (a *AdminAPI) Create(ctx *gin.Context, data *model.AdminAPI) error {
	return query.AdminAPI.WithContext(ctx).Create(data)
}

func (a *AdminAPI) Update(ctx *gin.Context, data *model.AdminAPI) error {
	return query.AdminAPI.WithContext(ctx).Where(query.AdminAPI.ID.Eq(data.ID)).Save(data)
}

func (a *AdminAPI) Enable(ctx *gin.Context, id int32, enabled bool) error {
	apiDB := query.AdminAPI
	_, err := apiDB.WithContext(ctx).Where(apiDB.ID.Eq(id)).UpdateColumn(apiDB.IsEnabled, enabled)
	return err
}

func (a *AdminAPI) Delete(ctx *gin.Context, id int32) error {
	apiDB := query.AdminAPI
	_, err := apiDB.WithContext(ctx).Where(apiDB.ID.Eq(id)).Delete()
	return err
}

func (a *AdminAPI) FindList(ctx context.Context, params *admin_proto.ApiListReq) (total int64, list []*model.AdminAPI, err error) {
	offset, limit, base := common.HandleListBaseReq(params.Base)
	params.Base = base
	q := a.handleListReq(ctx, params)
	total, err = q.Count()
	if err != nil {
		return total, list, err
	}
	list, err = q.Order(a.handleListReqSortField(params.Base.SortField, params.Base.SortType)).Limit(limit).Offset(offset).Find()
	return total, list, nil
}

func (a *AdminAPI) FindAllValid(ctx context.Context) (list []*model.AdminAPI, err error) {
	apiDB := query.AdminAPI
	return apiDB.WithContext(ctx).Where(apiDB.IsEnabled.Is(true)).Find()
}

func (a *AdminAPI) FindById(ctx *gin.Context, id int32) (*model.AdminAPI, error) {
	return query.AdminAPI.WithContext(ctx).Where(query.AdminAPI.ID.Eq(id)).First()
}

func (a *AdminAPI) FindByPermissionIds(ctx *gin.Context, ids []int32) (data []*admin_proto.ApiItem, err error) {
	api := query.AdminAPI
	per := query.AdminPermissionAPI
	per.WithContext(ctx).Join(api, api.ID.EqCol(per.APIID)).Where(per.PermissionID.In(ids...)).Scan(&data).Error()
	return data, err
}

func (a *AdminAPI) handleListReqSortField(sortField, sortType string) field.Expr {
	api := query.AdminAPI
	var res field.OrderExpr
	switch sortField {
	case api.CreatedAt.ColumnName().String():
		res = api.CreatedAt
	case api.UpdatedAt.ColumnName().String():
		res = api.UpdatedAt
	case api.ID.ColumnName().String():
		fallthrough
	case "":
		fallthrough
	default:
		res = api.ID
	}
	if sortType == common.DESC {
		return res.Desc()
	}
	return res
}

func (a *AdminAPI) handleListReq(ctx context.Context, params *admin_proto.ApiListReq) (q query.IAdminAPIDo) {
	apiDB := query.AdminAPI
	q = apiDB.WithContext(ctx)
	if params.Path != "" {
		q = q.Where(apiDB.Path.Like(params.Key))
	}
	if params.Name != "" {
		q = q.Where(apiDB.Name.Like(params.Key))
	}
	if params.Key != "" {
		q = q.Where(apiDB.Key.Like(params.Key))
	}

	switch params.Base.Enabled {
	case common.EnabledValidQueryValue:
		q = q.Where(apiDB.IsEnabled.Is(true))
	case common.EnabledInvalidQueryValue:
		q = q.Where(apiDB.IsEnabled.Is(false))
	}

	if params.Base.CreateStartTime > 0 {
		q = q.Where(apiDB.CreatedAt.Gte(time.Unix(params.Base.CreateStartTime, 0)))
	}

	if params.Base.CreateEndTime > 0 {
		q = q.Where(apiDB.CreatedAt.Lte(time.Unix(params.Base.CreateEndTime, 0)))
	}

	return q
}
