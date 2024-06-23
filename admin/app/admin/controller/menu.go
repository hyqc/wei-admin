package controller

import (
	"admin/app/admin/logic"
	"admin/app/common"
	"admin/code"
	"admin/config"
	"admin/constant"
	"admin/pkg/core"
	"admin/pkg/validator"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MenuController struct {
	core.Controller
}

var (
	menuLogic = logic.NewAdminMenuLogic()
)

// List 菜单列表
func (MenuController) List(ctx *gin.Context) {
	msg := "MenuController.List"
	params := &admin_proto.MenuListReq{Base: common.NewListBaseReq()}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := menuLogic.List(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
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
