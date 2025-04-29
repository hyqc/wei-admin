package controller

import (
	"admin/app/admin/logic"
	"admin/app/admin/validate"
	"admin/app/common"
	"admin/code"
	"admin/constant"
	"admin/global"
	"admin/pkg/core"
	"admin/proto/admin_proto"
	"admin/proto/code_proto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RoleController struct {
	core.Controller
}

// List 角色列表
func (RoleController) List(ctx *gin.Context) {
	msg := "RoleController.List"
	params := &admin_proto.ReqAdminRoleList{Base: common.NewListBaseReq()}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminRole.List(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// All 全部有效角色
func (RoleController) All(ctx *gin.Context) {
	msg := "RoleController.All"
	result := code.NewCode(code_proto.ErrorCode_Success)
	data, err := logic.H.AdminRole.All(ctx)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Add 添加角色
func (RoleController) Add(ctx *gin.Context) {
	msg := "RoleController.Add"
	params := &admin_proto.ReqAdminRoleAdd{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminRole.Add(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Info 详情
func (RoleController) Info(ctx *gin.Context) {
	msg := "RoleController.Info"
	params := &admin_proto.ReqAdminRoleInfo{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	info, err := logic.H.AdminRole.Info(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(info)
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Edit 编辑角色
func (RoleController) Edit(ctx *gin.Context) {
	msg := "RoleController.Edit"
	params := &admin_proto.ReqAdminRoleEdit{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminRole.Edit(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Enable 启用禁用
func (RoleController) Enable(ctx *gin.Context) {
	msg := "RoleController.Enable"
	params := &admin_proto.ReqAdminRoleEnable{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminRole.Enable(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Delete 删除角色
func (RoleController) Delete(ctx *gin.Context) {
	msg := "RoleController.Delete"
	params := &admin_proto.ReqAdminRoleDelete{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminRole.Delete(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// BindPermissions 角色绑定权限
func (RoleController) BindPermissions(ctx *gin.Context) {
	msg := "RoleController.Delete"
	params := &admin_proto.ReqAdminRoleBindPermissions{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminRole.RoleBindPermissions(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Permissions 角色权限列表
func (RoleController) Permissions(ctx *gin.Context) {
	msg := "RoleController.Permissions"
	params := &admin_proto.ReqAdminRolePermissions{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminRole.RolePermissions(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	global.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}
