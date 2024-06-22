package dao

import (
	"admin/app/common"
	"admin/app/gen/model"
	"admin/app/gen/query"
	adminproto "admin/proto"
	"context"
	"gorm.io/gen/field"
	"time"
)

type IAdminAPI interface {
	FindList(ctx context.Context, params *adminproto.ApiListReq) (total int64, list []*model.AdminAPI, err error)
	FindAllValid(ctx context.Context) (list []*model.AdminAPI, err error) // 查找所有有效的接口
}

type AdminAPI struct {
}

func NewAdminAPI() *AdminAPI {
	return &AdminAPI{}
}

func (a *AdminAPI) FindList(ctx context.Context, params *adminproto.ApiListReq) (total int64, list []*model.AdminAPI, err error) {
	offset, limit := common.ListBaseReqHandle(params)
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

func (a *AdminAPI) handleListReqSortField(sortField, sortType string) field.Expr {
	api := query.AdminAPI.Table(query.AdminAPI.TableName())
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

func (a *AdminAPI) handleListReq(ctx context.Context, params *adminproto.ApiListReq) (q query.IAdminAPIDo) {
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
