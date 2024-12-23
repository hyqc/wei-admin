package dao

import (
	"admin/app/admin/dao/types"
	"admin/app/admin/gen/model"
	"admin/app/admin/gen/query"
	"admin/app/common"
	"admin/proto/admin_proto"
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gen/field"
	"time"
)

type IAdminAPI interface {
	Create(ctx *gin.Context, data *model.AdminAPI) error
	Info(ctx *gin.Context, id int32) (*model.AdminAPI, error)
	Update(ctx *gin.Context, data *model.AdminAPI) error
	Enable(ctx *gin.Context, id int32, enabled bool) error
	Delete(ctx *gin.Context, id int32) error
	List(ctx context.Context, params *admin_proto.ReqAdminApiList) (total int64, list []*model.AdminAPI, err error)
	FindAllValid(ctx context.Context) (list []*model.AdminAPI, err error) // 查找所有有效的接口
	FindByPermissionIds(ctx *gin.Context, ids []int32) (data []*types.AdminApiPermissionItem, err error)
	FindByIds(ctx *gin.Context, ids []int32, enabled bool) ([]*model.AdminAPI, error)
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
	db := query.AdminAPI
	_, err := db.WithContext(ctx).Where(db.ID.Eq(id)).UpdateColumn(db.IsEnabled, enabled)
	return err
}

func (a *AdminAPI) Delete(ctx *gin.Context, id int32) error {
	db := query.AdminAPI
	_, err := db.WithContext(ctx).Where(db.ID.Eq(id)).Delete()
	return err
}

func (a *AdminAPI) List(ctx context.Context, params *admin_proto.ReqAdminApiList) (total int64, list []*model.AdminAPI, err error) {
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
	db := query.AdminAPI
	return db.WithContext(ctx).Where(db.IsEnabled.Is(true)).Find()
}
func (a *AdminAPI) FindByIds(ctx *gin.Context, ids []int32, enabled bool) ([]*model.AdminAPI, error) {
	db := query.AdminAPI
	return db.WithContext(ctx).Where(db.IsEnabled.Is(enabled), db.ID.In(ids...)).Find()
}

func (a *AdminAPI) Info(ctx *gin.Context, id int32) (*model.AdminAPI, error) {
	return query.AdminAPI.WithContext(ctx).Where(query.AdminAPI.ID.Eq(id)).First()
}

func (a *AdminAPI) FindByPermissionIds(ctx *gin.Context, ids []int32) (data []*types.AdminApiPermissionItem, err error) {
	api := query.AdminAPI
	per := query.AdminPermissionAPI
	data = make([]*types.AdminApiPermissionItem, 0)
	err = api.WithContext(ctx).
		Select(api.ID, api.Key, api.Name, api.IsEnabled, api.Path, api.Describe, per.PermissionID).
		Join(per, per.APIID.EqCol(api.ID)).
		Where(per.PermissionID.In(ids...)).Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
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
	if sortType == "" {
		sortType = common.DESC
	}
	if sortType == common.DESC {
		return res.Desc()
	}
	return res
}

func (a *AdminAPI) handleListReq(ctx context.Context, params *admin_proto.ReqAdminApiList) (q query.IAdminAPIDo) {
	db := query.AdminAPI
	q = db.WithContext(ctx)

	switch params.Base.Enabled {
	case common.EnabledValidQueryValue:
		q = q.Where(db.IsEnabled.Is(true))
	case common.EnabledInvalidQueryValue:
		q = q.Where(db.IsEnabled.Is(false))
	}

	if params.Base.CreateStartTime > 0 {
		q = q.Where(db.CreatedAt.Gte(time.Unix(params.Base.CreateStartTime, 0)))
	}

	if params.Base.CreateEndTime > 0 {
		q = q.Where(db.CreatedAt.Lte(time.Unix(params.Base.CreateEndTime, 0)))
	}

	if params.Path != "" {
		q = q.Where(db.Path.Like("%" + params.Path + "%"))
	}
	if params.Name != "" {
		q = q.Where(db.Name.Like("%" + params.Name + "%"))
	}
	if params.Key != "" {
		q = q.Where(db.Key.Like("%" + params.Key + "%"))
	}

	return q
}
