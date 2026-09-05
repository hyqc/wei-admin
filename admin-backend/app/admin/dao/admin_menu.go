package dao

import (
	"admin/app/admin/gen/model"
	"admin/app/admin/gen/query"
	"admin/app/common"
	"admin/global"
	"admin/pkg/utils"
	"admin/proto/admin_proto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"time"
)

type IAdminMenu interface {
	Create(ctx *gin.Context, data *model.AdminMenu) error
	Update(ctx *gin.Context, data *model.AdminMenu) error
	Enable(ctx *gin.Context, id int32, enabled bool) error
	Delete(ctx *gin.Context, id int32) error
	FindMyMenus(ctx context.Context, adminId, menuId int) ([]*model.AdminMenu, error) // 根据管理员名称查询详情
	FindList(ctx *gin.Context, params *admin_proto.ReqAdminMenuList) (total int64, list []*model.AdminMenu, err error)
	FindAll(ctx *gin.Context) ([]*model.AdminMenu, error)         // 全部的菜单，包括禁用的
	FindAllValid(ctx context.Context) ([]*model.AdminMenu, error) // 获取全部有效菜单
	FindById(ctx *gin.Context, id int32) (*model.AdminMenu, error)
	CountByParentId(ctx *gin.Context, parentId int32) (int64, error) // 子菜单数量，用于判断是否为目录型菜单
	FindPages(ctx *gin.Context) ([]*model.AdminMenu, error)
	FindByIds(ctx *gin.Context, ids []int32) ([]*model.AdminMenu, error)
	Show(ctx *gin.Context, menuId int32, field string, show bool) error
}

type AdminMenu struct {
}

func newAdminMenu() IAdminMenu {
	return &AdminMenu{}
}

func (a *AdminMenu) Create(ctx *gin.Context, data *model.AdminMenu) error {
	return query.AdminMenu.WithContext(ctx).Create(data)
}

func (a *AdminMenu) Update(ctx *gin.Context, data *model.AdminMenu) error {
	return query.AdminMenu.WithContext(ctx).Where(query.AdminMenu.ID.Eq(data.ID)).Save(data)
}

func (a *AdminMenu) Enable(ctx *gin.Context, id int32, enabled bool) error {
	db := query.AdminMenu
	_, err := db.WithContext(ctx).Where(db.ID.Eq(id)).UpdateColumn(db.IsEnabled, enabled)
	return err
}
func (a *AdminMenu) Show(ctx *gin.Context, menuId int32, f string, show bool) error {
	db := query.AdminMenu
	_, err := db.WithContext(ctx).Where(db.ID.Eq(menuId)).Updates(map[string]any{
		fmt.Sprintf("is_%s", utils.CamelToSnake(f)): show,
	})
	return err
}

// Delete 删除菜单，并级联清理该菜单下的权限点及其接口绑定、角色绑定，避免产生孤儿权限点
func (a *AdminMenu) Delete(ctx *gin.Context, id int32) error {
	t := global.AppDB.DefaultMysql()
	err := t.Transaction(func(tx *gorm.DB) error {
		// 事务内必须使用绑定了 tx 的 query，否则操作不走事务
		q := query.Use(tx)
		menu := q.AdminMenu
		per := q.AdminPermission
		pa := q.AdminPermissionAPI
		rp := q.AdminRolePermission

		permissions, terr := per.WithContext(ctx).Where(per.MenuID.Eq(id)).Select(per.ID).Find()
		if terr != nil {
			return terr
		}
		if len(permissions) > 0 {
			permissionIds := make([]int32, 0, len(permissions))
			for _, item := range permissions {
				permissionIds = append(permissionIds, item.ID)
			}
			if _, terr = pa.WithContext(ctx).Where(pa.PermissionID.In(permissionIds...)).Delete(); terr != nil {
				return terr
			}
			if _, terr = rp.WithContext(ctx).Where(rp.PermissionID.In(permissionIds...)).Delete(); terr != nil {
				return terr
			}
		}
		if _, terr = per.WithContext(ctx).Where(per.MenuID.Eq(id)).Delete(); terr != nil {
			return terr
		}
		_, terr = menu.WithContext(ctx).Where(menu.ID.Eq(id)).Delete()
		return terr
	})

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

func (a *AdminMenu) FindList(ctx *gin.Context, params *admin_proto.ReqAdminMenuList) (total int64, list []*model.AdminMenu, err error) {
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

// CountByParentId 统计某菜单的直接子菜单数量（>0 表示它是目录型菜单，不是页面）
func (a *AdminMenu) CountByParentId(ctx *gin.Context, parentId int32) (int64, error) {
	return query.AdminMenu.WithContext(ctx).Where(query.AdminMenu.ParentID.Eq(parentId)).Count()
}

func (a *AdminMenu) FindPages(ctx *gin.Context) ([]*model.AdminMenu, error) {
	DB := query.AdminMenu
	perDB := query.AdminPermission
	return DB.WithContext(ctx).Distinct(DB.ID).Select(DB.Name, DB.ID, DB.Key, DB.Path, DB.ParentID).
		Join(perDB, perDB.MenuID.EqCol(DB.ID)).
		Where(DB.IsEnabled.Is(true), DB.IsHideInMenu.Is(false)).
		Find()
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
	if sortType == "" {
		sortType = common.DESC
	}
	if sortType == common.DESC {
		return res.Desc()
	}
	return res
}

func (a *AdminMenu) handleListReq(ctx context.Context, params *admin_proto.ReqAdminMenuList) (q query.IAdminMenuDo) {
	DB := query.AdminMenu
	q = DB.WithContext(ctx)

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

	if params.Path != "" {
		q = q.Where(DB.Path.Like("%" + params.Path + "%"))
	}
	if params.Name != "" {
		q = q.Where(DB.Name.Like("%" + params.Name + "%"))
	}
	if params.Key != "" {
		q = q.Where(DB.Key.Like("%" + params.Key + "%"))
	}
	if params.ParentId > 0 {
		q = q.Where(DB.ParentID.Eq(params.ParentId))
	}

	return q
}

func (a *AdminMenu) FindByIds(ctx *gin.Context, ids []int32) (list []*model.AdminMenu, err error) {
	return query.AdminMenu.WithContext(ctx).Where(query.AdminMenu.ID.In(ids...)).Find()
}
