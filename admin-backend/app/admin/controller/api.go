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

type APIController struct {
	core.Controller
}

func (APIController) List(ctx *gin.Context) {
	msg := "APIController.List"
	params := &admin_proto.ReqAdminApiList{Base: common.NewListBaseReq()}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminAPI.List(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
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
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

func (APIController) Add(ctx *gin.Context) {
	msg := "APIController.Add"
	params := &admin_proto.ReqAdminApiAdd{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminAPI.Add(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

func (APIController) Info(ctx *gin.Context) {
	msg := "APIController.Info"
	params := &admin_proto.ReqAdminApiInfo{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	info, err := logic.H.AdminAPI.Info(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(info)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

func (APIController) Edit(ctx *gin.Context) {
	msg := "APIController.Edit"
	params := &admin_proto.ReqAdminApiEdit{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminAPI.Edit(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

func (APIController) Enable(ctx *gin.Context) {
	msg := "APIController.Enable"
	params := &admin_proto.ReqAdminApiEnable{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminAPI.Enable(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

func (APIController) Delete(ctx *gin.Context) {
	msg := "APIController.Delete"
	params := &admin_proto.ReqAdminApiDelete{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminAPI.Delete(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}
