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

type UserController struct {
	core.Controller
}

// List 管理员列表
//
//	@Summary		管理员列表
//	@Description	管理员列表
//	@Tags			管理员账号接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminUserList							true	"请求参数"
//	@Success		200		{object}	code.Message{data=admin_proto.RespAdminUserListData}	"desc"
//	@Router			/admin/user/list [post]
func (UserController) List(ctx *gin.Context) {
	msg := "UserController.List"
	params := &admin_proto.ReqAdminUserList{Base: common.NewListBaseReq()}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminUser.List(ctx, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Add 添加管理员
//
//	@Summary		添加管理员
//	@Description	添加管理员
//	@Tags			管理员账号接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminUserAdd	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}		"desc"
//	@Router			/admin/user/add [post]
func (UserController) Add(ctx *gin.Context) {
	msg := "UserController.Add"
	params := &admin_proto.ReqAdminUserAdd{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminUser.Add(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
}

// Info 管理员详情
//
//	@Summary		管理员详情
//	@Description	管理员详情
//	@Tags			管理员账号接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminUserInfo						true	"请求参数"
//	@Success		200		{object}	code.Message{data=admin_proto.RespAccountInfoData}	"desc"
//	@Router			/admin/user/info [post]
func (UserController) Info(ctx *gin.Context) {
	msg := "UserController.Info"
	params := &admin_proto.ReqAdminUserInfo{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	info, err := logic.H.AdminUser.AccountInfo(ctx, params.AdminId, false, constant.AdminTokenExpireSeconds)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(info)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
}

// Edit 编辑管理员
//
//	@Summary		编辑管理员
//	@Description	编辑管理员
//	@Tags			管理员账号接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminUserEdit	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}			"desc"
//	@Router			/admin/user/edit [post]
func (UserController) Edit(ctx *gin.Context) {
	msg := "UserController.Edit"
	params := &admin_proto.ReqAdminUserEdit{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminUser.Edit(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
}

// EditPassword 编辑管理员密码
//
//	@Summary		编辑管理员密码
//	@Description	编辑管理员密码
//	@Tags			管理员账号接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminUserEditPassword	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}					"desc"
//	@Router			/admin/user/edit_pwd [post]
func (UserController) EditPassword(ctx *gin.Context) {
	msg := "UserController.EditPassword"
	params := &admin_proto.ReqAdminUserEditPassword{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminUser.EditPassword(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
}

// Enable 启用|禁用管理员账号
//
//	@Summary		启用|禁用管理员账号
//	@Description	启用|禁用管理员账号
//	@Tags			管理员账号接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminUserEnabled	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}			"desc"
//	@Router			/admin/user/enable [post]
func (UserController) Enable(ctx *gin.Context) {
	msg := "UserController.Enable"
	params := &admin_proto.ReqAdminUserEnabled{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminUser.Enable(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
}

// Delete 删除管理员
//
//	@Summary		删除管理员
//	@Description	删除管理员
//	@Tags			管理员账号接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminUserDelete	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}			"desc"
//	@Router			/admin/user/delete [post]
func (UserController) Delete(ctx *gin.Context) {
	msg := "UserController.Delete"
	params := &admin_proto.ReqAdminUserDelete{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminUser.Delete(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
}

// BindRoles 管理员绑定角色
//
//	@Summary		管理员绑定角色
//	@Description	管理员绑定角色
//	@Tags			管理员账号接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAdminUserBindRoles	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}				"desc"
//	@Router			/admin/user/bind_roles [post]
func (UserController) BindRoles(ctx *gin.Context) {
	msg := "UserController.BindRoles"
	params := &admin_proto.ReqAdminUserBindRoles{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	if err := logic.H.AdminUser.BindRoles(ctx, params); err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
}
