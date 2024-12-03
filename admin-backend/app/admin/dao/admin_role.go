package dao

import (
	"admin/app/admin/dao/types"
	model2 "admin/app/admin/gen/model"
	query2 "admin/app/admin/gen/query"
	"admin/app/common"
	"admin/config"
	"admin/proto/admin_proto"
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"time"
)

type IAdminRole interface {
	FindAdminUserByRoleIds(ctx context.Context, ids []int32) ([]*model2.AdminUserRole, error)
	List(ctx context.Context, params *admin_proto.ReqAdminRoleList) (int64, []*model2.AdminRole, error)
	Create(ctx context.Context, data *model2.AdminRole) error
	Info(ctx context.Context, id int32) (*types.AdminRoleInfo, error)
	FindById(ctx context.Context, id int32) (*model2.AdminRole, error)
	Update(ctx context.Context, info *model2.AdminRole) error
	Enable(ctx context.Context, id int32, enabled bool) error
	Delete(ctx context.Context, id int32) error
	FindAdminRolePermissionByRoleId(ctx context.Context, id int32) ([]*types.AdminRolePermissionItem, error)
	BindPermissions(ctx context.Context, id int32, data []*model2.AdminRolePermission) error
	FindRolesById(ctx context.Context, id int32) ([]*model2.AdminRole, error)
	FindAllValid(ctx *gin.Context) ([]*model2.AdminRole, error)
}

type AdminRole struct {
}

func newAdminRole() IAdminRole {
	return &AdminRole{}
}

func (a *AdminRole) FindAdminUserByRoleIds(ctx context.Context, ids []int32) ([]*model2.AdminUserRole, error) {
	return query2.AdminUserRole.WithContext(ctx).Where(query2.AdminUserRole.RoleID.In(ids...)).Find()
}

func (a *AdminRole) List(ctx context.Context, params *admin_proto.ReqAdminRoleList) (total int64, list []*model2.AdminRole, err error) {
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
	api := query2.AdminRole
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

func (a *AdminRole) handleListReq(ctx context.Context, params *admin_proto.ReqAdminRoleList) (q query2.IAdminRoleDo) {
	db := query2.AdminRole
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

func (a *AdminRole) Create(ctx context.Context, data *model2.AdminRole) error {
	return query2.AdminRole.WithContext(ctx).Create(data)
}

func (a *AdminRole) Info(ctx context.Context, id int32) (*types.AdminRoleInfo, error) {
	t1 := query2.AdminRole
	t2 := query2.AdminUser
	b := t2.As("b")
	c := t2.As("c")
	data := &types.AdminRoleInfo{}
	err := t1.WithContext(ctx).Select(t1.ID, t1.Name, t1.Describe, t1.IsEnabled, t1.CreateAdminID, b.Username.As("create_admin_name"), t1.ModifyAdminID, c.Username.As("modify_admin_name"), t1.CreatedAt, t1.UpdatedAt).
		LeftJoin(b, b.ID.EqCol(t1.CreateAdminID)).
		LeftJoin(c, c.ID.EqCol(t1.ModifyAdminID)).Where(t1.ID.Eq(id)).Scan(data)
	return data, err
}

func (a *AdminRole) FindById(ctx context.Context, id int32) (*model2.AdminRole, error) {
	return query2.AdminRole.WithContext(ctx).Where(query2.AdminRole.ID.Eq(id)).First()
}

func (a *AdminRole) Update(ctx context.Context, data *model2.AdminRole) error {
	return query2.AdminRole.WithContext(ctx).Where(query2.AdminRole.ID.Eq(data.ID)).Save(data)
}

func (a *AdminRole) Enable(ctx context.Context, id int32, enabled bool) error {
	db := query2.AdminRole
	_, err := db.WithContext(ctx).Where(db.ID.Eq(id)).UpdateColumn(db.IsEnabled, enabled)
	return err
}

func (a *AdminRole) Delete(ctx context.Context, id int32) error {
	db := query2.AdminRole
	_, err := db.WithContext(ctx).Where(db.ID.Eq(id)).Delete()
	return err
}

func (a *AdminRole) FindAdminRolePermissionByRoleId(ctx context.Context, roleId int32) (list []*types.AdminRolePermissionItem, err error) {
	ta := query2.AdminRolePermission.As("a")
	tb := query2.AdminPermission.As("b")
	err = ta.WithContext(ctx).
		Select(ta.RoleID, ta.PermissionID, tb.Key.As("permission_key"), tb.Name.As("permission_name"), tb.Type.As("permission_type")).
		Join(tb, tb.ID.EqCol(ta.PermissionID)).
		Where(tb.IsEnabled.Is(true), ta.RoleID.Eq(roleId)).
		Scan(&list)
	return list, err
}

func (a *AdminRole) BindPermissions(ctx context.Context, roleId int32, data []*model2.AdminRolePermission) error {
	pa := query2.AdminRolePermission
	err := config.AppConfig.DBClient.Wei.Transaction(func(tx *gorm.DB) error {
		if _, err := pa.WithContext(ctx).Where(pa.PermissionID.Eq(roleId)).Delete(); err != nil {
			return err
		}
		return pa.WithContext(ctx).Create(data...)
	})
	return err
}

func (a *AdminRole) FindRolesById(ctx context.Context, id int32) ([]*model2.AdminRole, error) {
	return query2.AdminRole.WithContext(ctx).Where(query2.AdminRole.ID.Eq(id)).Find()
}

func (a *AdminRole) FindAllValid(ctx *gin.Context) ([]*model2.AdminRole, error) {
	return query2.AdminRole.WithContext(ctx).Where(query2.AdminRole.IsEnabled.Is(true)).Find()
}
