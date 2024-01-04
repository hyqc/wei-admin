package service

import (
	"admin/app/admin/dao"
	"admin/app/constant"
	"admin/app/gen/model"
	"context"
)

type AdminPermissionService struct {
	dao *dao.AdminPermission
}

func NewAdminPermissionService() *AdminPermissionService {
	return &AdminPermissionService{
		dao: adminPermissionDao,
	}
}

// GetMyPermission 获取我的权限列表
func (p *AdminPermissionService) GetMyPermission(ctx context.Context, adminId int32) ([]*model.AdminPermission, error) {
	if constant.IsAdministrator(adminId) {
		// 超管
		return adminPermissionDao.GetAdministerPermissions(ctx)
	}
	// 非超管
	return adminPermissionDao.GetMyPermissions(ctx, adminId, 0)
}
