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

type APIController struct {
	core.Controller
}

var (
	apiLogic = logic.NewAdminAPILogic()
)

func (APIController) List(ctx *gin.Context) {
	msg := "APIController.List"
	params := &admin_proto.ApiListReq{Base: common.NewListBaseReq()}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := apiLogic.List(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

func (APIController) All(ctx *gin.Context) {
	msg := "APIController.All"
	result := code.NewCode(code_proto.ErrorCode_Success)
	data, err := apiLogic.AllValid(ctx)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

func (APIController) Add(ctx *gin.Context) {
	msg := "APIController.Add"
	params := &admin_proto.ApiAddReq{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.APIReq.AddReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := apiLogic.Add(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

func (APIController) Info(ctx *gin.Context) {
	msg := "APIController.Info"
	params := &admin_proto.ApiInfoReq{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.APIReq.InfoReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	info, err := apiLogic.Info(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(info)
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

func (APIController) Edit(ctx *gin.Context) {
	msg := "APIController.Edit"
	params := &admin_proto.ApiEditReq{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.APIReq.EditReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := apiLogic.Edit(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

func (APIController) Enable(ctx *gin.Context) {
	msg := "APIController.Enable"
	params := &admin_proto.ApiEnableReq{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.APIReq.EnableReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := apiLogic.Enable(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

func (APIController) Delete(ctx *gin.Context) {
	msg := "APIController.Delete"
	params := &admin_proto.ApiDeleteReq{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.APIReq.DeleteReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := apiLogic.Delete(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}
