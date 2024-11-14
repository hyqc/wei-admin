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

func (APIController) List(ctx *gin.Context) {
	msg := "APIController.List"
	params := &admin_proto.ReqApiList{Base: common.NewListBaseReq()}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminAPI.List(ctx, params)
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
	data, err := logic.H.AdminAPI.AllValid(ctx)
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
	params := &admin_proto.ReqApiAdd{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminApiReq.AddReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminAPI.Add(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

func (APIController) Info(ctx *gin.Context) {
	msg := "APIController.Info"
	params := &admin_proto.ReqApiInfo{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminApiReq.InfoReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	info, err := logic.H.AdminAPI.Info(ctx, params)
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
	params := &admin_proto.ReqApiEdit{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminApiReq.EditReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminAPI.Edit(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

func (APIController) Enable(ctx *gin.Context) {
	msg := "APIController.Enable"
	params := &admin_proto.ReqApiEnable{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminApiReq.EnableReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminAPI.Enable(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

func (APIController) Delete(ctx *gin.Context) {
	msg := "APIController.Delete"
	params := &admin_proto.ReqApiDelete{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminApiReq.DeleteReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminAPI.Delete(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}
