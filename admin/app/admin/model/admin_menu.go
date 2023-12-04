package model

import (
	"admin/app/gen/model"
	"context"
)

type IAdminMenu interface {
	FindMyMenusByAdminId(ctx context.Context, adminId int) ([]*model.AdminMenu, error) // 根据管理员名称查询详情
}

type AdminMenu struct {
}

func NewAdminMenu() *AdminMenu {
	return &AdminMenu{}
}

// FindMyMenusByAdminId 获取我的可以访问的菜单列表
func (a *AdminMenu) FindMyMenusByAdminId(ctx context.Context, adminId int) ([]*model.AdminMenu, error) {
	return nil, nil
}
