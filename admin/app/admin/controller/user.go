package controller

import (
	"admin/pkg/core"
	"github.com/gin-gonic/gin"
)

type AdminUserController struct {
	core.Controller
}

// List 管理员列表
func (AdminUserController) List(ctx *gin.Context) {

}

// All 全部有效管理员
func (AdminUserController) All(ctx *gin.Context) {

}

// Add 添加管理员
func (AdminUserController) Add(ctx *gin.Context) {

}

// Detail 详情
func (AdminUserController) Detail(ctx *gin.Context) {

}

// Edit 编辑管理员
func (AdminUserController) Edit(ctx *gin.Context) {

}

// Enable 启用禁用
func (AdminUserController) Enable(ctx *gin.Context) {

}

// Delete 删除管理员
func (AdminUserController) Delete(ctx *gin.Context) {

}

// BindRoles 管理员绑定角色
func (AdminUserController) BindRoles(ctx *gin.Context) {

}
