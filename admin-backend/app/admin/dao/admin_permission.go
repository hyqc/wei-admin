package dao

import (
	"admin/app/admin/gen/custom/admin_custom"
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

type IAdminPermission interface {
	FindAdministerPermissions(ctx context.Context) ([]*model2.AdminPermission, error) // 根据管理员名称查询详情
	FindAdminPermissions(ctx context.Context, adminId, menuId int32) ([]*model2.AdminPermission, error)
	FindPermissionsByMenuId(ctx context.Context, menuId int32) ([]*model2.AdminPermission, error)
	List(ctx *gin.Context, params *admin_proto.ReqPermissionList) (total int64, list []*model2.AdminPermission, err error)
	Create(ctx *gin.Context, data *model2.AdminPermission) error
	Info(ctx *gin.Context, id int32) (*model2.AdminPermission, error)
	Update(ctx *gin.Context, data *model2.AdminPermission) error
	Enable(ctx *gin.Context, id int32, enabled bool) error
	FindPermissionMenuInfoById(ctx *gin.Context, permissionId int32) (*admin_custom.AdminPermissionMenu, error)
	Delete(ctx *gin.Context, id int32) error
	BindApis(ctx *gin.Context, permissionId int32, permissionApes []*model2.AdminPermissionAPI) error
	BatchAddPermissions(ctx *gin.Context, data []*model2.AdminPermission) error
	FindByIds(ctx *gin.Context, ids []int32) ([]*model2.AdminPermission, error)
}

type AdminPermission struct {
}

func newAdminPermission() *AdminPermission {
	return &AdminPermission{}
}

type AdminPermissionType string

const (
	AdminPermissionTypeAll    AdminPermissionType = "all"
	AdminPermissionTypeView   AdminPermissionType = "view"
	AdminPermissionTypeEdit   AdminPermissionType = "edit"
	AdminPermissionTypeDelete AdminPermissionType = "delete"
)

var (
	AdminPermissionTypeTextMap = map[AdminPermissionType]string{
		AdminPermissionTypeView:   "查看",
		AdminPermissionTypeEdit:   "编辑",
		AdminPermissionTypeDelete: "删除",
	}
)

func GetAdminPermissionTypeText(t AdminPermissionType) string {
	if _, ok := AdminPermissionTypeTextMap[t]; ok {
		return AdminPermissionTypeTextMap[t]
	}
	return ""
}

// FindAdministerPermissions 获取超管对应的权限
func (a *AdminPermission) FindAdministerPermissions(ctx context.Context) ([]*model2.AdminPermission, error) {
	permission := query2.AdminPermission
	menu := query2.AdminMenu
	return permission.WithContext(ctx).
		Join(menu, menu.ID.EqCol(permission.MenuID), menu.IsEnabled.Is(true)).
		Where(permission.IsEnabled.Is(true)).Order(permission.MenuID).Find()
}

// FindAdminPermissions 获取非超管对于的权限
func (a *AdminPermission) FindAdminPermissions(ctx context.Context, adminId, menuId int32) ([]*model2.AdminPermission, error) {
	permission := query2.AdminPermission
	menu := query2.AdminMenu
	role := query2.AdminRole
	rolePermission := query2.AdminRolePermission
	userRole := query2.AdminUserRole
	user := query2.AdminUser

	db := permission.WithContext(ctx).
		Join(rolePermission, rolePermission.PermissionID.EqCol(permission.ID)).
		Join(role, role.ID.EqCol(rolePermission.RoleID), role.IsEnabled.Is(true)).
		Join(userRole, userRole.RoleID.EqCol(role.ID), userRole.AdminID.Eq(adminId)).
		Join(user, user.ID.EqCol(userRole.AdminID)).
		Join(menu, menu.ID.EqCol(permission.MenuID), menu.IsEnabled.Is(true)).
		Where(user.ID.Eq(adminId))
	if menuId > 0 {
		db = db.Where(permission.MenuID.Eq(menuId))
	}
	return db.Order(permission.MenuID).Find()
}

func (a *AdminPermission) FindPermissionsByMenuId(ctx context.Context, menuId int32) ([]*model2.AdminPermission, error) {
	permission := query2.AdminPermission
	return permission.WithContext(ctx).Where(permission.MenuID.Eq(menuId)).Find()
}

func (a *AdminPermission) List(ctx *gin.Context, params *admin_proto.ReqPermissionList) (total int64, list []*model2.AdminPermission, err error) {
	DB := query2.AdminPermission
	offset, limit, base := common.HandleListBaseReq(params.Base)
	params.Base = base
	q := a.handleListReq(ctx, params)
	total, err = q.Count()
	if err != nil {
		return total, list, err
	}
	list, err = q.Order(DB.MenuID, DB.ID).Limit(limit).Offset(offset).Find()
	return total, list, err
}

func (a *AdminPermission) handleListReq(ctx context.Context, params *admin_proto.ReqPermissionList) (q query2.IAdminPermissionDo) {
	DB := query2.AdminPermission
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

	if params.MenuId > 0 {
		q = q.Where(DB.MenuID.Eq(params.MenuId))
	}
	if params.Name != "" {
		q = q.Where(DB.Name.Like("%" + params.Name + "%"))
	}
	if params.Key != "" {
		q = q.Where(DB.Key.Like("%" + params.Key + "%"))
	}
	if params.Type != "" && params.Type != string(AdminPermissionTypeAll) {
		q = q.Where(DB.Type.Eq(params.Type))
	}

	return q
}

func (a *AdminPermission) Create(ctx *gin.Context, data *model2.AdminPermission) error {
	return query2.AdminPermission.WithContext(ctx).Create(data)
}

func (a *AdminPermission) Info(ctx *gin.Context, id int32) (*model2.AdminPermission, error) {
	return query2.AdminPermission.WithContext(ctx).Where(query2.AdminPermission.ID.Eq(id)).First()
}
func (a *AdminPermission) Update(ctx *gin.Context, data *model2.AdminPermission) error {
	_, err := query2.AdminPermission.WithContext(ctx).Where(query2.AdminPermission.ID.Eq(data.ID)).Updates(data)
	return err
}

func (a *AdminPermission) Enable(ctx *gin.Context, id int32, enabled bool) error {
	db := query2.AdminPermission
	_, err := db.WithContext(ctx).Where(db.ID.Eq(id)).UpdateColumn(db.IsEnabled, enabled)
	return err
}

func (a *AdminPermission) Delete(ctx *gin.Context, id int32) error {
	db := query2.AdminPermission
	_, err := db.WithContext(ctx).Where(db.ID.Eq(id)).Delete()
	return err
}

func (a *AdminPermission) BatchAddPermissions(ctx *gin.Context, data []*model2.AdminPermission) error {
	return query2.AdminPermission.WithContext(ctx).Create(data...)
}

func (a *AdminPermission) BindApis(ctx *gin.Context, permissionId int32, permissionApes []*model2.AdminPermissionAPI) error {
	pa := query2.AdminPermissionAPI
	err := config.AppConfig.DBClient.Wei.Transaction(func(tx *gorm.DB) error {
		if _, err := pa.WithContext(ctx).Where(pa.PermissionID.Eq(permissionId)).Delete(); err != nil {
			return err
		}
		return pa.WithContext(ctx).Create(permissionApes...)
	})
	return err
}

func (a *AdminPermission) FindPermissionMenuInfoById(ctx *gin.Context, permissionId int32) (*admin_custom.AdminPermissionMenu, error) {
	p := query2.AdminPermission
	m := query2.AdminMenu
	data := &admin_custom.AdminPermissionMenu{}
	fields := []field.Expr{
		p.ID,
		p.MenuID,
		p.Name,
		p.Key,
		p.Type,
		p.Describe,
		p.CreatedAt,
		p.UpdatedAt,
		m.ID.As("menu_id"),
		p.IsEnabled.As("enabled"),
		m.Name.As("menu_name"),
		m.Path.As("menu_path"),
	}
	err := p.WithContext(ctx).Select(fields...).LeftJoin(m, m.ID.EqCol(p.MenuID)).Where(p.ID.Eq(permissionId)).Scan(data)
	return data, err
}

func (a *AdminPermission) FindByIds(ctx *gin.Context, ids []int32) ([]*model2.AdminPermission, error) {
	return query2.AdminPermission.WithContext(ctx).Where(query2.AdminPermission.ID.In(ids...)).Find()
}
