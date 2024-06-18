package logic

import (
	"admin/app/admin/dao"
	"admin/app/gen/model"
	"admin/constant"
	"context"
)

type PermissionLogic struct {
	dao *dao.AdminPermission
}

func NewAdminPermissionLogic() *PermissionLogic {
	return &PermissionLogic{
		dao: adminPermissionDao,
	}
}

func (p *PermissionLogic) FindMyPermission(ctx context.Context, adminId int32) ([]*model.AdminPermission, error) {
	if constant.IsAdministrator(adminId) {
		// 超管
		return adminPermissionDao.FindAdministerPermissions(ctx)
	}
	// 非超管
	return adminPermissionDao.FindAdminPermissions(ctx, adminId, 0)
}

func (p *PermissionLogic) Permissions2MenuIds(permissions []*model.AdminPermission) (pageIds []int32, permissionKeys map[string]string) {
	// 管理员可以访问的菜单
	menuIdsM := make(map[int32]struct{})
	permissionKeys = make(map[string]string)
	for _, item := range permissions {
		if _, ok := menuIdsM[item.MenuID]; !ok {
			menuIdsM[item.MenuID] = struct{}{}
			pageIds = append(pageIds, item.MenuID)
		}
		permissionKeys[item.Key] = item.Name
	}
	return
}
