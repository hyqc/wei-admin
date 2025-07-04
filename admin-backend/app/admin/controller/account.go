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

type AccountController struct {
	core.Controller
}

func (AccountController) Register(ctx *gin.Context) {
	msg := "AccountController.Register"
	result := code.NewCode(code_proto.ErrorCode_Success)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Login 登录
//
//	@Summary		登录后台
//	@Description	登录后台
//	@Tags			账号相关接口
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqLogin							true	"请求参数"
//	@Success		200		{object}	code.Message{data=admin_proto.RespLoginData}	"desc"
//	@Router			/admin/account/login [post]
func (AccountController) Login(ctx *gin.Context) {
	msg := "AccountController.Login"
	params := &admin_proto.ReqLogin{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminUser.AccountLogin(ctx, params, ctx.ClientIP())
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)

	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Info 管理员账号详情
//
//	@Summary		账号详情
//	@Description	账号详情
//	@Tags			账号相关接口
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{object}	code.Message{data=admin_proto.RespAccountInfoData}	"desc"
//	@Router			/admin/account/info [post]
func (AccountController) Info(ctx *gin.Context) {
	msg := "AccountController.Info"
	refreshToken := ctx.GetBool("refreshToken")
	result := code.NewCode(code_proto.ErrorCode_Success)
	data, err := logic.H.AdminUser.AccountInfo(ctx, constant.GetCustomClaims(ctx).AccountId, refreshToken, constant.AdminTokenExpireSeconds)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}

	result.SetData(data)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Edit 编辑账号
//
//	@Summary		编辑账号
//	@Description	编辑账号
//	@Tags			账号相关接口
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAccountEdit	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}		"desc"
//	@Router			/admin/account/edit [post]
func (AccountController) Edit(ctx *gin.Context) {
	msg := "AccountController.Edit"
	params := &admin_proto.ReqAccountEdit{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	err := logic.H.AdminUser.AccountEdit(ctx, constant.GetCustomClaims(ctx).AccountId, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Password 修改密码
//
//	@Summary		修改密码
//	@Description	修改密码
//	@Tags			账号相关接口
//	@Accept			application/json
//	@Produce		application/json
//	@Param			object	query		admin_proto.ReqAccountPasswordEdit	true	"请求参数"
//	@Success		200		{object}	code.Message{data=nil}				"desc"
//	@Router			/admin/account/password [post]
func (AccountController) Password(ctx *gin.Context) {
	msg := "AccountController.Password"
	params := &admin_proto.ReqAccountPasswordEdit{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validate.WithCtx(ctx, params); err != nil {
		result.SetCodeMsg(code_proto.ErrorCode_RequestParamsInvalid, err)
		global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	err := logic.H.AdminUser.AccountEditPassword(ctx, constant.GetCustomClaims(ctx).AccountId, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Menu 登录用户可访问的菜单
//
//	@Summary		登录用户可访问的菜单
//	@Description	登录用户可访问的菜单
//	@Tags			账号相关接口
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{object}	code.Message{data=[]admin_proto.MenuItem}	"desc"
//	@Router			/admin/account/menu [post]
func (AccountController) Menu(ctx *gin.Context) {
	msg := "AccountController.Menu"
	result := code.NewCode(code_proto.ErrorCode_Success)
	data, err := logic.H.AdminUser.MyMenus(ctx, constant.GetCustomClaims(ctx).AccountId)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Permission 我的权限
//
//	@Summary		我的权限
//	@Description	我的权限
//	@Tags			账号相关接口
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{object}	code.Message{data=map[string]string}	"desc"
//	@Router			/admin/account/permission [post]
func (AccountController) Permission(ctx *gin.Context) {
	msg := "AccountController.Permission"
	result := code.NewCode(code_proto.ErrorCode_Success)
	data, err := logic.H.AdminUser.MyPermission(ctx, constant.GetCustomClaims(ctx).AccountId)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	global.LogSugar.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}
