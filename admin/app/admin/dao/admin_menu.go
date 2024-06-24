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

type IAdminMenu interface {
	Create(ctx *gin.Context, data *model.AdminMenu) error
	Update(ctx *gin.Context, data *model.AdminMenu) error
	Enable(ctx *gin.Context, id int32, enabled bool) error
	Delete(ctx *gin.Context, id int32) error
	FindMyMenus(ctx context.Context, adminId, menuId int) ([]*model.AdminMenu, error) // 根据管理员名称查询详情
	FindList(ctx *gin.Context, params *admin_proto.MenuListReq) (total int64, list []*model.AdminMenu, err error)
	FindAll(ctx *gin.Context) ([]*model.AdminMenu, error)         // 全部的菜单，包括禁用的
	FindAllValid(ctx context.Context) ([]*model.AdminMenu, error) // 获取全部有效菜单
	FindById(ctx *gin.Context, id int32) (*model.AdminMenu, error)
}

type AdminMenu struct {
}

func NewAdminMenu() *AdminMenu {
	return &AdminMenu{}
}

func (a *AdminMenu) Create(ctx *gin.Context, data *model.AdminMenu) error {
	return query.AdminMenu.WithContext(ctx).Create(data)
}

func (a *AdminMenu) Update(ctx *gin.Context, data *model.AdminMenu) error {
	return query.AdminMenu.WithContext(ctx).Where(query.AdminMenu.ID.Eq(data.ID)).Save(data)
}

func (a *AdminMenu) Enable(ctx *gin.Context, id int32, enabled bool) error {
	apiDB := query.AdminMenu
	_, err := apiDB.WithContext(ctx).Where(apiDB.ID.Eq(id)).UpdateColumn(apiDB.IsEnabled, enabled)
	return err
}

func (a *AdminMenu) Delete(ctx *gin.Context, id int32) error {
	apiDB := query.AdminMenu
	_, err := apiDB.WithContext(ctx).Where(apiDB.ID.Eq(id)).Delete()
	return err
}

// FindMyMenus 获取我的可以访问的菜单列表
func (a *AdminMenu) FindMyMenus(ctx context.Context, adminId, menuId int) ([]*model.AdminMenu, error) {
	return nil, nil
}

// FindAllValid 获取全部有效菜单
func (a *AdminMenu) FindAllValid(ctx context.Context) ([]*model.AdminMenu, error) {
	menu := query.AdminMenu
	return menu.WithContext(ctx).Where(menu.IsEnabled.Is(true)).Order(menu.Sort, menu.ParentID, menu.ID).Find()
}

func (a *AdminMenu) FindAll(ctx *gin.Context) ([]*model.AdminMenu, error) {
	menu := query.AdminMenu
	return menu.WithContext(ctx).Order(menu.Sort, menu.ParentID, menu.ID).Find()
}

func (a *AdminMenu) FindList(ctx *gin.Context, params *admin_proto.MenuListReq) (total int64, list []*model.AdminMenu, err error) {
	DB := query.AdminMenu
	offset, limit, base := common.HandleListBaseReq(params.Base)
	params.Base = base
	q := a.handleListReq(ctx, params)
	total, err = q.Count()
	if err != nil {
		return total, list, err
	}
	list, err = q.Order(DB.ParentID, DB.Sort.Desc(), DB.ID).Limit(limit).Offset(offset).Find()
	return total, list, nil
}

func (a *AdminMenu) FindById(ctx *gin.Context, id int32) (*model.AdminMenu, error) {
	return query.AdminMenu.WithContext(ctx).Where(query.AdminMenu.ID.Eq(id)).First()
}

func (a *AdminMenu) handleListReqSortField(sortField, sortType string) field.Expr {
	DB := query.AdminMenu
	var res field.OrderExpr
	switch sortField {
	case DB.CreatedAt.ColumnName().String():
		res = DB.CreatedAt
	case DB.UpdatedAt.ColumnName().String():
		res = DB.UpdatedAt
	case DB.ID.ColumnName().String():
		fallthrough
	case "":
		fallthrough
	default:
		res = DB.ID
	}
	if sortType == common.DESC {
		return res.Desc()
	}
	return res
}

func (a *AdminMenu) handleListReq(ctx context.Context, params *admin_proto.MenuListReq) (q query.IAdminMenuDo) {
	DB := query.AdminMenu
	q = DB.WithContext(ctx)
	if params.Path != "" {
		q = q.Where(DB.Path.Like(params.Key))
	}
	if params.Name != "" {
		q = q.Where(DB.Name.Like(params.Key))
	}
	if params.Key != "" {
		q = q.Where(DB.Key.Like(params.Key))
	}
	if params.ParentId > 0 {
		q = q.Where(DB.ParentID.Eq(params.ParentId))
	}

	switch params.Base.Enabled {
	case common.EnabledValidQueryValue:
		q = q.Where(DB.IsEnabled.Is(true))
	case common.EnabledInvalidQueryValue:
		q = q.Where(DB.IsEnabled.Is(false))
	}

	if params.Base.CreateStartTime > 0 {
		q = q.Where(DB.CreatedAt.Gte(time.Unix(params.Base.CreateStartTime, 0)))
	}

	if params.Base.CreateEndTime > 0 {
		q = q.Where(DB.CreatedAt.Lte(time.Unix(params.Base.CreateEndTime, 0)))
	}

	return q
}
