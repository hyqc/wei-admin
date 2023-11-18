package controller

import (
	"admin/pkg/core"
	"github.com/gin-gonic/gin"
)

type MenuController struct {
	core.Controller
}

// List 菜单列表
func (MenuController) List(ctx *gin.Context) {

}

// Tree 有效菜单树
func (MenuController) Tree(ctx *gin.Context) {

}

// Add 添加菜单
func (MenuController) Add(ctx *gin.Context) {

}

// Detail 菜单详情
func (MenuController) Detail(ctx *gin.Context) {

}

// Edit 编辑菜单
func (MenuController) Edit(ctx *gin.Context) {

}

// Enable 启用禁用菜单
func (MenuController) Enable(ctx *gin.Context) {

}

// Delete 删除菜单
func (MenuController) Delete(ctx *gin.Context) {

}

// Permissions 菜单权限列表
func (MenuController) Permissions(ctx *gin.Context) {

}

// Pages 页面菜单列表
func (MenuController) Pages(ctx *gin.Context) {

}

// Mode 页面模块权限列表
func (MenuController) Mode(ctx *gin.Context) {

}
