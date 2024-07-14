package dao

import (
	"admin/app/common"
	"admin/app/gen/model"
	"admin/app/gen/query"
	"admin/proto/admin_proto"
	"context"
	"gorm.io/gen/field"
	"time"
)

type IAdminUser interface {
	FindAdminUserByUsername(ctx context.Context, username string) (*model.AdminUser, error)   // 根据管理员名称查询详情
	UpdateAdminUserLoginData(ctx context.Context, adminId int32, data *model.AdminUser) error // 更新管理员的登录信息
	FindAdminUserByAdminId(ctx context.Context, id int32) (*model.AdminUser, error)           // 根据管理员ID查询详情
	UpdateAdminUser(ctx context.Context, data *model.AdminUser) error
	Create(ctx context.Context, data *model.AdminUser) error
	List(ctx context.Context, params *admin_proto.AdminUserListReq, adminIds []int32) (total int64, list []*model.AdminUser, err error)
}

type AdminUser struct {
}

func newAdminUser() *AdminUser {
	return &AdminUser{}
}

func (a *AdminUser) FindAdminUserByUsername(ctx context.Context, username string) (*model.AdminUser, error) {
	db := query.AdminUser
	return db.WithContext(ctx).Where(db.Username.Eq(username)).First()
}

func (a *AdminUser) UpdateAdminUserLoginData(ctx context.Context, adminId int32, data *model.AdminUser) error {
	db := query.AdminUser
	_, err := db.WithContext(ctx).Where(db.ID.Eq(adminId)).
		Select(db.LoginTotal, db.LastLoginIP, db.LastLoginTime).
		Updates(data)
	return err
}

func (a *AdminUser) FindAdminUserByAdminId(ctx context.Context, id int32) (*model.AdminUser, error) {
	db := query.AdminUser
	return db.WithContext(ctx).Where(db.ID.Eq(id)).First()
}

func (a *AdminUser) UpdateAdminUser(ctx context.Context, data *model.AdminUser) error {
	db := query.AdminUser
	_, err := db.WithContext(ctx).Where(db.ID.Eq(data.ID)).Updates(data)
	return err
}

func (a *AdminUser) Create(ctx context.Context, data *model.AdminUser) error {
	return query.AdminUser.WithContext(ctx).Create(data)
}

func (a *AdminUser) List(ctx context.Context, params *admin_proto.AdminUserListReq, adminIds []int32) (total int64, list []*model.AdminUser, err error) {
	offset, limit, base := common.HandleListBaseReq(params.Base)
	params.Base = base
	q := a.handleListReq(ctx, params, adminIds)
	total, err = q.Count()
	if err != nil {
		return total, list, err
	}
	list, err = q.Order(a.handleListReqSortField(params.Base.SortField, params.Base.SortType)).Limit(limit).Offset(offset).Find()
	return total, list, nil
}

func (a *AdminUser) handleListReqSortField(sortField, sortType string) field.Expr {
	api := query.AdminUser
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

func (a *AdminUser) handleListReq(ctx context.Context, params *admin_proto.AdminUserListReq, adminIds []int32) (q query.IAdminUserDo) {
	db := query.AdminUser
	q = db.WithContext(ctx)

	if adminIds != nil && len(adminIds) > 0 {
		q = q.Where(db.ID.In(adminIds...))
	}

	if params.Username != "" {
		q = q.Where(db.Username.Like(params.Username))
	}
	if params.Email != "" {
		q = q.Where(db.Email.Like(params.Email))
	}
	if params.Nickname != "" {
		q = q.Where(db.Nickname.Like(params.Nickname))
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
