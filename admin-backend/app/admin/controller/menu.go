package controller

import (
	"admin/app/admin/logic"
	"admin/app/admin/validate"
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

// List 菜单列表
func (MenuController) List(ctx *gin.Context) {
	msg := "MenuController.List"
	params := &admin_proto.ReqMenuList{Base: common.NewListBaseReq()}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminMenuReq.ListReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminMenu.List(ctx, params)
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
	msg := "MenuController.Tree"
	result := code.NewCode(code_proto.ErrorCode_Success)
	data, err := logic.H.AdminMenu.Tree(ctx)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Add 添加菜单
func (MenuController) Add(ctx *gin.Context) {
	msg := "MenuController.Add"
	params := &admin_proto.ReqMenuAdd{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminMenuReq.AddReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminMenu.Add(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Info 菜单详情
func (MenuController) Info(ctx *gin.Context) {
	msg := "MenuController.Info"
	params := &admin_proto.ReqMenuInfo{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminMenuReq.InfoReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	info, err := logic.H.AdminMenu.Info(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(info)
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Edit 编辑菜单
func (MenuController) Edit(ctx *gin.Context) {
	msg := "MenuController.Edit"
	params := &admin_proto.ReqMenuEdit{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminMenuReq.EditReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminMenu.Edit(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Enable 启用禁用菜单
func (MenuController) Enable(ctx *gin.Context) {
	msg := "MenuController.Enable"
	params := &admin_proto.ReqMenuEnable{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminMenuReq.EnableReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminMenu.Enable(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Delete 删除菜单
func (MenuController) Delete(ctx *gin.Context) {
	msg := "MenuController.Delete"
	params := &admin_proto.ReqMenuDelete{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminMenuReq.DeleteReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminMenu.Delete(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Permissions 菜单权限列表
func (MenuController) Permissions(ctx *gin.Context) {
	msg := "MenuController.Permissions"
	params := &admin_proto.ReqMenuPermissions{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminMenuReq.PermissionsReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminMenu.Permissions(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Pages 页面菜单列表
func (MenuController) Pages(ctx *gin.Context) {
	msg := "MenuController.Permissions"
	params := &admin_proto.ReqMenuPages{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminMenuReq.PagesReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminMenu.Pages(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Mode 页面模块权限列表
func (MenuController) Mode(ctx *gin.Context) {

}
