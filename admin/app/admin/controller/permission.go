package controller

import (
	"admin/pkg/core"
	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	core.Controller
}

// List 权限列表
func (PermissionController) List(ctx *gin.Context) {

}

// Add 添加权限
func (PermissionController) Add(ctx *gin.Context) {

}

// Detail 详情
func (PermissionController) Detail(ctx *gin.Context) {

}

// Edit 编辑权限
func (PermissionController) Edit(ctx *gin.Context) {

}

// Enable 启用禁用
func (PermissionController) Enable(ctx *gin.Context) {

}

// Delete 删除权限
func (PermissionController) Delete(ctx *gin.Context) {

}

// BindAPI 绑定权限接口
func (PermissionController) BindAPI(ctx *gin.Context) {

}

// AddMenuPermissions 指定菜单创建查看编辑删除权限
func (PermissionController) AddMenuPermissions(ctx *gin.Context) {

}
