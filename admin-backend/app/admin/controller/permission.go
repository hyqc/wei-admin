package controller

import (
	"admin/app/admin/logic"
	"admin/app/common"
	"admin/code"
	"admin/constant"
	"admin/global"
	"admin/pkg/core"
	"admin/pkg/govalidate"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PermissionController struct {
	core.Controller
}

// List 权限列表
func (PermissionController) List(ctx *gin.Context) {
	msg := "PermissionController.List"
	params := &admin_proto.ReqAdminPermissionList{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := govalidate.ValidateWithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminPermission.List(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Add 添加权限
func (PermissionController) Add(ctx *gin.Context) {
	msg := "PermissionController.Add"
	params := &admin_proto.ReqAdminPermissionAdd{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := govalidate.ValidateWithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.Add(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
}

// Info 详情
func (PermissionController) Info(ctx *gin.Context) {
	msg := "PermissionController.Info"
	params := &admin_proto.ReqAdminPermissionInfo{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := govalidate.ValidateWithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	info, err := logic.H.AdminPermission.Info(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(info)
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Edit 编辑权限
func (PermissionController) Edit(ctx *gin.Context) {
	msg := "PermissionController.Edit"
	params := &admin_proto.ReqAdminPermissionEdit{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := govalidate.ValidateWithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.Edit(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Enable 启用禁用
func (PermissionController) Enable(ctx *gin.Context) {
	msg := "PermissionController.Enable"
	params := &admin_proto.ReqAdminPermissionEnable{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := govalidate.ValidateWithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.Enable(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Delete 删除权限
func (PermissionController) Delete(ctx *gin.Context) {
	msg := "PermissionController.Delete"
	params := &admin_proto.ReqAdminPermissionDelete{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := govalidate.ValidateWithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.Delete(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// BindAPI 绑定权限接口
func (PermissionController) BindAPI(ctx *gin.Context) {
	msg := "PermissionController.BindAPI"
	params := &admin_proto.ReqAdminPermissionBindApis{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := govalidate.ValidateWithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.BindAPI(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// UnBindAPI 解绑权限接口
func (PermissionController) UnBindAPI(ctx *gin.Context) {
	msg := "PermissionController.UnBindAPI"
	params := &admin_proto.ReqAdminPermissionUnBindApi{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := govalidate.ValidateWithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.UnBindAPI(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// AddMenuPermissions 指定菜单创建查看编辑删除权限
func (PermissionController) AddMenuPermissions(ctx *gin.Context) {
	msg := "PermissionController.AddMenuPermissions"
	params := &admin_proto.ReqAdminPermissionBindMenu{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := govalidate.ValidateWithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.AddMenuPermissions(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}
