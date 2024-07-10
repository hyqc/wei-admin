package logic

import (
	"admin/app/admin/dao"
	"admin/app/gen/model"
	"admin/constant"
	"context"
)

// 获取我的权限
func getMyAdminPermissions(ctx context.Context, adminId int32) ([]*model.AdminPermission, error) {
	if constant.IsAdministrator(adminId) {
		// 超管
		return dao.H.AdminPermission.FindAdministerPermissions(ctx)
	}
	// 非超管
	return dao.H.AdminPermission.FindAdminPermissions(ctx, adminId, 0)
}

// 从权限中获取页面，权限数据
func getMenuIdsFromAdminPermissions(permissions []*model.AdminPermission) (pageIds []int32, permissionKeys map[string]string) {
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
