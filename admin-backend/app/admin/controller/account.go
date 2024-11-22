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

type AccountController struct {
	core.Controller
}

func (AccountController) Register(ctx *gin.Context) {
	msg := "AccountController.Register"
	result := code.NewCode(code_proto.ErrorCode_Success)
	config.AppLogger.Sugar().Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Login 登录
// @Summary 登录后台
// @Description 登录后台
// @Tags 账号相关接口
// @Accept application/json
// @Produce application/json
// @Param object query admin_proto.LoginReq true "请求参数"
// @Success 200 {object}
// @Router /admin_proto/login [post]
func (AccountController) Login(ctx *gin.Context) {
	msg := "AccountController.Login"
	params := &admin_proto.ReqLogin{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminAccountReq.LoginReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := logic.H.AdminUser.AccountLogin(ctx, params, ctx.ClientIP())
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)

	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Info 管理员账号详情
func (AccountController) Info(ctx *gin.Context) {
	msg := "AccountController.Info"
	refreshToken := ctx.GetBool("refreshToken")
	result := code.NewCode(code_proto.ErrorCode_Success)
	data, err := logic.H.AdminUser.AccountInfo(ctx, constant.GetCustomClaims(ctx).AdminID, refreshToken, constant.AdminTokenExpireSeconds)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}

	result.SetData(data)
	config.AppLogger.Sugar().Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Edit 编辑账号
func (AccountController) Edit(ctx *gin.Context) {
	msg := "AccountController.Edit"
	params := &admin_proto.ReqAccountEdit{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminAccountReq.AccountEditReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	err := logic.H.AdminUser.AccountEdit(ctx, constant.GetCustomClaims(ctx).AdminID, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Password 修改密码
func (AccountController) Password(ctx *gin.Context) {
	msg := "AccountController.Password"
	params := &admin_proto.ReqAccountPasswordEdit{}
	result := code.NewCode(code_proto.ErrorCode_Success)
	if err := validator.Validate(ctx, params, validate.AdminAccountReq.AccountEditPasswordReq); err != nil {
		result.SetCodeError(code_proto.ErrorCode_RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	err := logic.H.AdminUser.AccountEditPassword(ctx, constant.GetCustomClaims(ctx).AdminID, params)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

// Menu 登录用户可访问的菜单
func (AccountController) Menu(ctx *gin.Context) {
	msg := "AccountController.Menu"
	result := code.NewCode(code_proto.ErrorCode_Success)
	data, err := logic.H.AdminUser.MyMenus(ctx, constant.GetCustomClaims(ctx).AdminID)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}

func (AccountController) Permission(ctx *gin.Context) {
	msg := "AccountController.Permission"
	result := code.NewCode(code_proto.ErrorCode_Success)
	data, err := logic.H.AdminUser.MyPermission(ctx, constant.GetCustomClaims(ctx).AdminID)
	if err != nil {
		common.HandleLogicError(ctx, err, msg, result)
		return
	}
	result.SetData(data)
	config.AppLoggerSugared.Debugw(msg, zap.Any(constant.LogResponseMsgField, result))
	code.JSON(ctx, result)
	return
}