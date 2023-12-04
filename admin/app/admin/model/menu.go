package model

import (
	"admin/app/gen/model"
	"context"
)

type IAdminMenu interface {
	FindMyMenusByAdminId(ctx context.Context, username string) (*model.AdminUser, error) // 根据管理员名称查询详情
}

type AdminMenu struct {
}

func NewIAdminMenu() *AdminMenu {
	return &AdminMenu{}
}

// FindMyMenusByAdminId 获取我的可以访问的菜单列表
func (a *AdminMenu) FindMyMenusByAdminId(adminId int) {

}
