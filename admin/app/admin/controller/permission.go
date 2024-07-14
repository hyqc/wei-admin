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

type PermissionController struct {
	core.Controller
}

// List 权限列表
func (PermissionController) List(ctx *gin.Context) {
	msg := "PermissionController.List"
	params := &admin_proto.PermissionListReq{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminPermissionReq.ListReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminPermission.List(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Add 添加权限
func (PermissionController) Add(ctx *gin.Context) {
	msg := "PermissionController.Add"
	params := &admin_proto.PermissionAddReq{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminPermissionReq.AddReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.Add(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
}

// Info 详情
func (PermissionController) Info(ctx *gin.Context) {
	msg := "PermissionController.AccountInfo"
	params := &admin_proto.PermissionInfoReq{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminPermissionReq.InfoReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	info, err := logic.H.AdminPermission.Info(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(info)
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Edit 编辑权限
func (PermissionController) Edit(ctx *gin.Context) {
	msg := "PermissionController.AccountEdit"
	params := &admin_proto.PermissionEditReq{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminPermissionReq.EditReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.Edit(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Enable 启用禁用
func (PermissionController) Enable(ctx *gin.Context) {
	msg := "PermissionController.Enable"
	params := &admin_proto.PermissionEnableReq{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminPermissionReq.EnableReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.Enable(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Delete 删除权限
func (PermissionController) Delete(ctx *gin.Context) {
	msg := "PermissionController.Delete"
	params := &admin_proto.PermissionDeleteReq{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminPermissionReq.DeleteReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.Delete(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// BindAPI 绑定权限接口
func (PermissionController) BindAPI(ctx *gin.Context) {
	msg := "PermissionController.BindAPI"
	params := &admin_proto.PermissionBindApisReq{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminPermissionReq.BindAPIReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.BindAPI(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// AddMenuPermissions 指定菜单创建查看编辑删除权限
func (PermissionController) AddMenuPermissions(ctx *gin.Context) {
	msg := "PermissionController.AddMenuPermissions"
	params := &admin_proto.PermissionBindMenuReq{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminPermissionReq.PermissionBindMenuReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminPermission.AddMenuPermissions(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}
