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

type IAdminPermission interface {
	FindAdministerPermissions(ctx context.Context) ([]*model.AdminPermission, error) // 根据管理员名称查询详情
	FindAdminPermissions(ctx context.Context, adminId, menuId int32) ([]*model.AdminPermission, error)
	FindPermissionsByMenuId(ctx context.Context, menuId int32) ([]*model.AdminPermission, error)
	List(ctx *gin.Context, params *admin_proto.PermissionListReq) (total int64, list []*model.AdminPermission, err error)
	Create(ctx *gin.Context, data *model.AdminPermission) error
	FindById(ctx *gin.Context, id int32) (*model.AdminPermission, error)
	FindPermissionMenuInfoById(ctx *gin.Context, permissionId int32) (*admin_custom.AdminPermissionMenu, error)
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
func (a *AdminPermission) FindAdministerPermissions(ctx context.Context) ([]*model.AdminPermission, error) {
	permission := query.AdminPermission
	menu := query.AdminMenu
	return permission.WithContext(ctx).
		Join(menu, menu.ID.EqCol(permission.MenuID), menu.IsEnabled.Is(true)).
		Where(permission.IsEnabled.Is(true)).Order(permission.MenuID).Find()
}

// FindAdminPermissions 获取非超管对于的权限
func (a *AdminPermission) FindAdminPermissions(ctx context.Context, adminId, menuId int32) ([]*model.AdminPermission, error) {
	permission := query.AdminPermission
	menu := query.AdminMenu
	role := query.AdminRole
	rolePermission := query.AdminRolePermission
	userRole := query.AdminUserRole

	db := permission.WithContext(ctx).
		Join(menu, menu.ID.EqCol(permission.MenuID), menu.IsEnabled.Is(true)).
		Join(rolePermission, rolePermission.RoleID.EqCol(userRole.RoleID)).
		Join(role, role.ID.EqCol(userRole.RoleID), role.IsEnabled.Is(true)).
		Join(userRole, userRole.RoleID.EqCol(role.ID), userRole.AdminID.Eq(adminId)).
		Where(userRole.AdminID.Eq(adminId))
	if menuId > 0 {
		db.Where(permission.MenuID.Eq(menuId))
	}
	return db.Order(permission.MenuID).Find()
}

func (a *AdminPermission) FindPermissionsByMenuId(ctx context.Context, menuId int32) ([]*model.AdminPermission, error) {
	permission := query.AdminPermission
	return permission.WithContext(ctx).Where(permission.MenuID.Eq(menuId)).Find()
}

func (a *AdminPermission) List(ctx *gin.Context, params *admin_proto.PermissionListReq) (total int64, list []*model.AdminPermission, err error) {
	DB := query.AdminPermission
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

func (a *AdminPermission) handleListReq(ctx context.Context, params *admin_proto.PermissionListReq) (q query.IAdminPermissionDo) {
	DB := query.AdminPermission
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

func (a *AdminPermission) Create(ctx *gin.Context, data *model.AdminPermission) error {
	return query.AdminPermission.WithContext(ctx).Create(data)
}

func (a *AdminPermission) FindById(ctx *gin.Context, id int32) (*model.AdminPermission, error) {
	return query.AdminPermission.WithContext(ctx).Where(query.AdminPermission.ID.Eq(id)).First()
}
func (a *AdminPermission) FindPermissionMenuInfoById(ctx *gin.Context, permissionId int32) (*admin_custom.AdminPermissionMenu, error) {
	p := query.AdminPermission
	m := query.AdminMenu
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
