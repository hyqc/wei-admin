package dao

import (
	"admin/app/common"
	"admin/app/gen/custom/admin_custom"
	"admin/app/gen/model"
	"admin/app/gen/query"
	"admin/proto/admin_proto"
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gen/field"
	"time"
)

type IAdminRole interface {
	FindAdminUserByRoleIds(ctx context.Context, ids []int32) ([]*model.AdminUserRole, error)
	List(ctx *gin.Context, params *admin_proto.RoleListReq) (int64, []*model.AdminRole, error)
	Create(ctx *gin.Context, data *model.AdminRole) error
	Info(ctx *gin.Context, id int32) (*admin_custom.AdminRoleInfo, error)
	FindById(ctx *gin.Context, id int32) (*model.AdminRole, error)
	Update(ctx *gin.Context, info *model.AdminRole) error
	Enable(ctx *gin.Context, id int32, enabled bool) error
	Delete(ctx *gin.Context, id int32) error
}

type AdminRole struct {
}

func newAdminRole() IAdminRole {
	return &AdminRole{}
}

func (a *AdminRole) FindAdminUserByRoleIds(ctx context.Context, ids []int32) ([]*model.AdminUserRole, error) {
	return query.AdminUserRole.WithContext(ctx).Where(query.AdminUserRole.RoleID.In(ids...)).Find()
}

func (a *AdminRole) List(ctx *gin.Context, params *admin_proto.RoleListReq) (total int64, list []*model.AdminRole, err error) {
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

func (a *AdminRole) handleListReqSortField(sortField, sortType string) field.Expr {
	api := query.AdminRole
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

func (a *AdminRole) handleListReq(ctx context.Context, params *admin_proto.RoleListReq) (q query.IAdminRoleDo) {
	db := query.AdminRole
	q = db.WithContext(ctx)
	if params.Id > 0 {
		q = q.Where(db.ID.Eq(params.Id))
	}
	if params.Name != "" {
		q = q.Where(db.Name.Like("%" + params.Name + "%"))
	}

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

	return q
}

func (a *AdminRole) Create(ctx *gin.Context, data *model.AdminRole) error {
	return query.AdminRole.WithContext(ctx).Create(data)
}

func (a *AdminRole) Info(ctx *gin.Context, id int32) (*admin_custom.AdminRoleInfo, error) {
	t1 := query.AdminRole
	t2 := query.AdminUser
	b := t2.As("b")
	c := t2.As("c")
	data := &admin_custom.AdminRoleInfo{}
	err := t1.WithContext(ctx).Select().
		LeftJoin(b, b.ID.EqCol(t1.CreateAdminID)).
		LeftJoin(c, c.ID.EqCol(t1.ModifyAdminID)).Where(t1.ID.Eq(id)).Scan(data)
	return data, err
}

func (a *AdminRole) FindById(ctx *gin.Context, id int32) (*model.AdminRole, error) {
	return query.AdminRole.WithContext(ctx).Where(query.AdminRole.ID.Eq(id)).First()
}

func (a *AdminRole) Update(ctx *gin.Context, data *model.AdminRole) error {
	return query.AdminRole.WithContext(ctx).Where(query.AdminRole.ID.Eq(data.ID)).Save(data)
}

func (a *AdminRole) Enable(ctx *gin.Context, id int32, enabled bool) error {
	db := query.AdminRole
	_, err := db.WithContext(ctx).Where(db.ID.Eq(id)).UpdateColumn(db.IsEnabled, enabled)
	return err
}

func (a *AdminRole) Delete(ctx *gin.Context, id int32) error {
	db := query.AdminRole
	_, err := db.WithContext(ctx).Where(db.ID.Eq(id)).Delete()
	return err
}
