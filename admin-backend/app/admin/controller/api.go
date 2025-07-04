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

// List 接口列表
//
//	@Summary		接口列表
//	@Description	接口列表
//	@Tags			API接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminApiList							true	"请求参数"
//	@Success		200		{object}	code.Message{data=admin_proto.RespAdminApiListData}	"desc"
//	@Router			/admin/api/list [post]
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

// All 全部有效接口列表
//
//	@Summary		全部有效接口列表
//	@Description	全部有效接口列表
//	@Tags			API接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{object}	code.Message{data=[]admin_proto.AdminApiItem}	"desc"
//	@Router			/admin/api/all [post]
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

// Add 新增接口
//
//	@Summary		新增接口
//	@Description	新增接口
//	@Tags			API接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminApiAdd	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}		"desc"
//	@Router			/admin/api/add [post]
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

// Info 接口详情
//
//	@Summary		接口详情
//	@Description	接口详情
//	@Tags			API接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminApiInfo					true	"请求参数"
//	@Success		200		{object}	code.Message{data=admin_proto.AdminApiItem}	"desc"
//	@Router			/admin/api/info [post]
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

// Edit 接口编辑
//
//	@Summary		接口编辑
//	@Description	接口编辑
//	@Tags			API接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminApiEdit	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}		"desc"
//	@Router			/admin/api/edit [post]
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

// Enable 接口启用/禁用
//
//	@Summary		接口启用/禁用
//	@Description	接口启用/禁用
//	@Tags			API接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminApiEnable	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}			"desc"
//	@Router			/admin/api/enable [post]
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

// Delete 接口删除
//
//	@Summary		接口删除
//	@Description	接口删除
//	@Tags			API接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminApiDelete	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}			"desc"
//	@Router			/admin/api/delete [post]
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
