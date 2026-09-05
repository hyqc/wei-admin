package dao

import (
	"admin/app/admin/dao/types"
	"admin/app/admin/gen/model"
	"admin/app/admin/gen/query"
	"admin/app/common"
	"admin/code"
	"admin/constant"
	"admin/global"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"time"
)

type IAdminPermission interface {
	FindAdministerPermissions(ctx context.Context) ([]*model.AdminPermission, error) // 根据管理员名称查询详情
	FindAdminPermissions(ctx context.Context, adminId, menuId int32) ([]*model.AdminPermission, error)
	FindPermissionsByMenuId(ctx context.Context, menuId int32) ([]*model.AdminPermission, error)
	List(ctx *gin.Context, params *admin_proto.ReqAdminPermissionList) (total int64, list []*model.AdminPermission, err error)
	Create(ctx *gin.Context, data *model.AdminPermission) error
	Info(ctx *gin.Context, id int32) (*model.AdminPermission, error)
	Update(ctx *gin.Context, data *model.AdminPermission) error
	Enable(ctx *gin.Context, id int32, enabled bool) error
	FindPermissionMenuInfoById(ctx *gin.Context, permissionId int32) (*types.AdminPermissionMenu, error)
	Delete(ctx *gin.Context, id int32) error
	BindApis(ctx *gin.Context, permissionId int32, permissionApes []*model.AdminPermissionAPI) error
	UnBindApi(ctx *gin.Context, permissionId, apiId int32) error
	BatchAddPermissions(ctx *gin.Context, menuId int32, data []PermissionSyncItem) error
	FindByIds(ctx *gin.Context, ids []int32) ([]*model.AdminPermission, error)
	IsAdminCanAccessPath(ctx context.Context, adminId int32, path string) (bool, error)
}

type AdminPermission struct {
}

// PermissionSyncItem 菜单权限配置的同步项：权限点 + 该操作需要访问的接口
// ApiIds 为 nil 表示本次不改动接口绑定，非 nil（含空切片）表示全量覆盖为该集合
type PermissionSyncItem struct {
	Permission *model.AdminPermission
	ApiIds     []int32
}

func newAdminPermission() *AdminPermission {
	return &AdminPermission{}
}

type AdminPermissionType string

// 权限动作类型固定三类：view 查看（读）/ edit 编辑（写，含新增、重置、绑定等一切写操作）/ delete 删除
// AdminPermissionTypeAll 仅用于列表筛选时表示"全部类型"
const (
	AdminPermissionTypeAll    AdminPermissionType = "all"
	AdminPermissionTypeView   AdminPermissionType = "view"
	AdminPermissionTypeEdit   AdminPermissionType = "edit"
	AdminPermissionTypeDelete AdminPermissionType = "delete"
)

var (
	AdminPermissionTypeTextMap = map[AdminPermissionType]string{
		AdminPermissionTypeAll:    "全部",
		AdminPermissionTypeView:   "查看",
		AdminPermissionTypeEdit:   "编辑",
		AdminPermissionTypeDelete: "删除",
	}
)

// GetAdminPermissionTypeText 获取权限类型展示名
// 权限类型是"软枚举"：内置类型返回中文名，自定义类型（如 export/audit）直接返回原值，
// 保证列表/详情里不会出现空白类型（与 common.GetPermissionTypeName 行为保持一致）
func GetAdminPermissionTypeText(t AdminPermissionType) string {
	if val, ok := AdminPermissionTypeTextMap[t]; ok {
		return val
	}
	return string(t)
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
	user := query.AdminUser

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

func (a *AdminPermission) FindPermissionsByMenuId(ctx context.Context, menuId int32) ([]*model.AdminPermission, error) {
	permission := query.AdminPermission
	return permission.WithContext(ctx).Where(permission.MenuID.Eq(menuId)).Order(permission.ID).Find()
}

func (a *AdminPermission) List(ctx *gin.Context, params *admin_proto.ReqAdminPermissionList) (total int64, list []*model.AdminPermission, err error) {
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

func (a *AdminPermission) handleListReq(ctx context.Context, params *admin_proto.ReqAdminPermissionList) (q query.IAdminPermissionDo) {
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
	// 反查：绑定了该接口的权限点
	if params.ApiId > 0 {
		PA := query.AdminPermissionAPI
		q = q.Join(PA, PA.PermissionID.EqCol(DB.ID), PA.APIID.Eq(params.ApiId))
	}

	return q
}

func (a *AdminPermission) Create(ctx *gin.Context, data *model.AdminPermission) error {
	return query.AdminPermission.WithContext(ctx).Create(data)
}

func (a *AdminPermission) Info(ctx *gin.Context, id int32) (*model.AdminPermission, error) {
	return query.AdminPermission.WithContext(ctx).Where(query.AdminPermission.ID.Eq(id)).First()
}
func (a *AdminPermission) Update(ctx *gin.Context, data *model.AdminPermission) error {
	_, err := query.AdminPermission.WithContext(ctx).Where(query.AdminPermission.ID.Eq(data.ID)).Updates(data)
	return err
}

func (a *AdminPermission) Enable(ctx *gin.Context, id int32, enabled bool) error {
	db := query.AdminPermission
	_, err := db.WithContext(ctx).Where(db.ID.Eq(id)).UpdateColumn(db.IsEnabled, enabled)
	return err
}

func (a *AdminPermission) Delete(ctx *gin.Context, id int32) error {
	t := global.AppDB.DefaultMysql()
	err := t.Transaction(func(tx *gorm.DB) error {
		db := query.AdminPermission
		_, err := db.WithContext(ctx).Where(db.ID.Eq(id)).Delete()
		if err != nil {
			return err
		}
		pa := query.AdminPermissionAPI
		_, err = pa.WithContext(ctx).Where(pa.PermissionID.Eq(id)).Delete()
		return err
	})

	return err
}

// BatchAddPermissions 同步"菜单权限配置"提交的权限点集合：
//  1. 校验权限唯一键：本次提交内部不可重复，也不可与其它权限点冲突；
//  2. 已存在的权限点按 id 全字段更新（含 key，解决改 key 不生效的问题），
//     删除该菜单下已存在但本次未提交的权限点（级联清理接口绑定与角色绑定），
//     保证 UI 中删除"操作权限"行后对应的权限点真正被移除；
//  3. 逐项全量覆盖接口绑定：ApiIds 为 nil 表示本次不动绑定，非 nil（含空切片）表示覆盖为指定集合；
//  4. 提交空集合表示清空该菜单全部权限点。
func (a *AdminPermission) BatchAddPermissions(ctx *gin.Context, menuId int32, data []PermissionSyncItem) error {
	return global.AppDB.DefaultMysql().Transaction(func(tx *gorm.DB) error {
		// 事务内必须使用绑定了 tx 的 query，否则操作不走事务
		q := query.Use(tx)
		p := q.AdminPermission
		pa := q.AdminPermissionAPI
		rp := q.AdminRolePermission

		// 1. 唯一键校验
		submittedKeys := make(map[string]struct{}, len(data))
		for _, item := range data {
			key := item.Permission.Key
			if _, ok := submittedKeys[key]; ok {
				return code.NewCodeError(code_proto.ErrorCode_AdminPermissionKeyExist, nil)
			}
			submittedKeys[key] = struct{}{}
			conflict := p.WithContext(ctx).Where(p.Key.Eq(key))
			if item.Permission.ID > 0 {
				conflict = conflict.Where(p.ID.Neq(item.Permission.ID))
			}
			if count, err := conflict.Count(); err != nil {
				return err
			} else if count > 0 {
				return code.NewCodeError(code_proto.ErrorCode_AdminPermissionKeyExist, nil)
			}
		}

		// 2. 该菜单现有权限点
		existing, err := p.WithContext(ctx).Where(p.MenuID.Eq(menuId)).Find()
		if err != nil {
			return err
		}
		existingIds := make(map[int32]struct{}, len(existing))
		for _, e := range existing {
			existingIds[e.ID] = struct{}{}
		}

		submittedIds := make(map[int32]struct{}, len(data))
		for _, item := range data {
			perm := item.Permission
			perm.MenuID = menuId
			if perm.ID > 0 && !existingIdExists(existingIds, perm.ID) {
				// 提交了一个不属于本菜单的权限点 id，按新增处理
				perm.ID = 0
			}
			if perm.ID > 0 {
				// 已存在：显式指定列更新（结构体 Updates 会忽略零值导致无法禁用）
				if _, err := p.WithContext(ctx).Where(p.ID.Eq(perm.ID)).Updates(map[string]interface{}{
					"menu_id":    perm.MenuID,
					"key":        perm.Key,
					"name":       perm.Name,
					"type":       perm.Type,
					"describe":   perm.Describe,
					"is_enabled": perm.IsEnabled,
				}); err != nil {
					return err
				}
			} else {
				if err := p.WithContext(ctx).Create(perm); err != nil {
					return err
				}
			}
			submittedIds[perm.ID] = struct{}{}

			// 3. 接口绑定：ApiIds 为 nil 时保持原绑定不变
			if item.ApiIds != nil {
				if _, err := pa.WithContext(ctx).Where(pa.PermissionID.Eq(perm.ID)).Delete(); err != nil {
					return err
				}
				if len(item.ApiIds) > 0 {
					bindings := make([]*model.AdminPermissionAPI, 0, len(item.ApiIds))
					for _, apiId := range item.ApiIds {
						bindings = append(bindings, &model.AdminPermissionAPI{PermissionID: perm.ID, APIID: apiId})
					}
					if err := pa.WithContext(ctx).Create(bindings...); err != nil {
						return err
					}
				}
			}
		}

		// 4. 删除本次未提交的旧权限点
		toBeDeleted := make([]int32, 0)
		for id := range existingIds {
			if _, ok := submittedIds[id]; !ok {
				toBeDeleted = append(toBeDeleted, id)
			}
		}
		if len(toBeDeleted) > 0 {
			if _, err = pa.WithContext(ctx).Where(pa.PermissionID.In(toBeDeleted...)).Delete(); err != nil {
				return err
			}
			if _, err = rp.WithContext(ctx).Where(rp.PermissionID.In(toBeDeleted...)).Delete(); err != nil {
				return err
			}
			if _, err = p.WithContext(ctx).Where(p.ID.In(toBeDeleted...)).Delete(); err != nil {
				return err
			}
		}
		return nil
	})
}

func existingIdExists(m map[int32]struct{}, id int32) bool {
	_, ok := m[id]
	return ok
}

func (a *AdminPermission) BindApis(ctx *gin.Context, permissionId int32, permissionApes []*model.AdminPermissionAPI) error {
	pa := query.AdminPermissionAPI
	err := global.AppDB.DefaultMysql().Transaction(func(tx *gorm.DB) error {
		if _, err := pa.WithContext(ctx).Where(pa.PermissionID.Eq(permissionId)).Delete(); err != nil {
			return err
		}
		return pa.WithContext(ctx).Create(permissionApes...)
	})
	return err
}

func (a *AdminPermission) UnBindApi(ctx *gin.Context, permissionId, apiId int32) error {
	pa := query.AdminPermissionAPI
	_, err := pa.WithContext(ctx).Where(pa.PermissionID.Eq(permissionId), pa.APIID.Eq(apiId)).Delete()
	return err
}

func (a *AdminPermission) FindPermissionMenuInfoById(ctx *gin.Context, permissionId int32) (*types.AdminPermissionMenu, error) {
	p := query.AdminPermission
	m := query.AdminMenu
	data := &types.AdminPermissionMenu{}
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

func (a *AdminPermission) FindByIds(ctx *gin.Context, ids []int32) ([]*model.AdminPermission, error) {
	return query.AdminPermission.WithContext(ctx).Where(query.AdminPermission.ID.In(ids...)).Find()
}

func (a *AdminPermission) IsAdminCanAccessPath(ctx context.Context, adminId int32, path string) (bool, error) {
	if constant.IsAdministrator(adminId) {
		return true, nil
	}
	user := query.AdminUser
	userRole := query.AdminUserRole
	role := query.AdminRole
	rolePermission := query.AdminRolePermission
	permission := query.AdminPermission
	perApi := query.AdminPermissionAPI
	api := query.AdminAPI

	data, err := user.WithContext(ctx).
		Join(userRole, userRole.AdminID.EqCol(user.ID)).
		Join(role, role.ID.EqCol(userRole.RoleID), role.IsEnabled.Is(true)).
		Join(rolePermission, rolePermission.RoleID.EqCol(role.ID)).
		Join(permission, permission.ID.EqCol(rolePermission.PermissionID), permission.IsEnabled.Is(true)).
		Join(perApi, perApi.PermissionID.EqCol(permission.ID)).
		Join(api, api.ID.EqCol(perApi.APIID), api.IsEnabled.Is(true)).
		Where(user.ID.Eq(adminId), user.IsEnabled.Is(true), api.Path.Eq(path)).First()
	if err != nil {
		return false, err
	}
	if data == nil {
		return false, nil
	}
	if data.ID == adminId {
		return true, nil
	}
	return false, nil
}
