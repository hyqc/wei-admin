package controller

import (
	"admin/app/admin/logic"
	"admin/app/admin/validate"
	"admin/code"
	"admin/config"
	"admin/constant"
	"admin/pkg/core"
	"admin/pkg/validator"
	adminproto "admin/proto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AccountController struct {
	core.Controller
}

var (
	accountLogic = logic.NewAdminUserLogic()
)

func (AccountController) Register(ctx *gin.Context) {
	result := code.NewCode(code.Success)
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
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
// @Router /admin/login [post]
func (AccountController) Login(ctx *gin.Context) {
	params := &adminproto.LoginReq{}
	result := code.NewCode(code.Success)
	if err := validator.Validate(ctx, params, validate.Account.LoginReq); err != nil {
		result.SetCodeError(code.RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw("Login", zap.Any("msg", result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := accountLogic.Login(ctx, params, ctx.ClientIP())
	if err != nil {
		res := code.GetCodeMsg(err)
		if res == nil {
			result.SetCodeError(code.RequestParamsInvalid, err)
			config.AppLoggerSugared.Debugw("Login", zap.Any("msg", result), zap.Any("error", err))
			code.JSON(ctx, result)
			return
		}

		config.AppLoggerSugared.Debugw("Login", zap.Any("msg", res), zap.Any("error", err))
		code.JSON(ctx, res)
		return
	}
	result.SetData(data)

	config.AppLoggerSugared.Debugw("Login", zap.Any("msg", result))
	code.JSON(ctx, result)
	return
}

// Info 管理员账号详情
func (AccountController) Info(ctx *gin.Context) {
	refreshToken := ctx.GetBool("refreshToken")
	result := code.NewCode(code.Success)
	data, err := accountLogic.Info(ctx, constant.GetCustomClaims(ctx).AdminID, refreshToken, 3600)
	if err != nil {
		res := code.GetCodeMsg(err)
		if res == nil {
			result.SetCodeError(code.RequestParamsInvalid, err)
			config.AppLoggerSugared.Debugw("Info", zap.Any("msg", result), zap.Any("error", err))
			code.JSON(ctx, result)
			return
		}

		config.AppLoggerSugared.Debugw("Info", zap.Any("msg", res), zap.Any("error", err))
		code.JSON(ctx, res)
		return
	}

	result.SetData(data)
	config.AppLogger.Sugar().Debugw("Info", zap.Any("msg", result))
	code.JSON(ctx, result)
	return
}

// Edit 编辑账号
func (AccountController) Edit(ctx *gin.Context) {
	params := &adminproto.AccountEditReq{}
	result := code.NewCode(code.Success)
	if err := validator.Validate(ctx, params, validate.Account.AccountEditReq); err != nil {
		result.SetCodeError(code.RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw("Edit", zap.Any("msg", result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	err := accountLogic.Edit(ctx, constant.GetCustomClaims(ctx).AdminID, params)
	if err != nil {
		res := code.GetCodeMsg(err)
		if res == nil {
			result.SetCodeError(code.RequestParamsInvalid, err)
			config.AppLoggerSugared.Debugw("Edit", zap.Any("msg", result), zap.Any("error", err))
			code.JSON(ctx, result)
			return
		}

		config.AppLoggerSugared.Debugw("Edit", zap.Any("msg", res), zap.Any("error", err))
		code.JSON(ctx, res)
		return
	}
	config.AppLoggerSugared.Debugw("Edit", zap.Any("msg", result))
	code.JSON(ctx, result)
	return
}

// Password 修改密码
func (AccountController) Password(ctx *gin.Context) {
	params := &adminproto.AccountPasswordEditReq{}
	result := code.NewCode(code.Success)
	if err := validator.Validate(ctx, params, validate.Account.AccountEditPasswordReq); err != nil {
		result.SetCodeError(code.RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw("Password", zap.Any("msg", result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	err := accountLogic.EditPassword(ctx, constant.GetCustomClaims(ctx).AdminID, params)
	if err != nil {
		res := code.GetCodeMsg(err)
		if res == nil {
			result.SetCodeError(code.RequestParamsInvalid, err)
			config.AppLoggerSugared.Debugw("Password", zap.Any("msg", result), zap.Any("error", err))
			code.JSON(ctx, result)
			return
		}

		config.AppLoggerSugared.Debugw("Password", zap.Any("msg", res), zap.Any("error", err))
		code.JSON(ctx, res)
		return
	}
	config.AppLoggerSugared.Debugw("Password", zap.Any("msg", result))
	code.JSON(ctx, result)
	return
}

// Menu 登录用户可访问的菜单
func (AccountController) Menu(ctx *gin.Context) {
	result := code.NewCode(code.Success)
	data, err := accountLogic.MyMenus(ctx, constant.GetCustomClaims(ctx).AdminID)
	if err != nil {
		res := code.GetCodeMsg(err)
		if res == nil {
			result.SetCodeError(code.RequestParamsInvalid, err)
			config.AppLoggerSugared.Debugw("Menu", zap.Any("msg", result), zap.Any("error", err))
			code.JSON(ctx, result)
			return
		}

		config.AppLoggerSugared.Debugw("Menu", zap.Any("msg", res), zap.Any("error", err))
		code.JSON(ctx, res)
		return
	}
	result.SetData(data)
	config.AppLoggerSugared.Debugw("Menu", zap.Any("msg", result))
	code.JSON(ctx, result)
	return
}

func (AccountController) Permission(ctx *gin.Context) {
	result := code.NewCode(code.Success)
	data, err := accountLogic.MyPermission(ctx, constant.GetCustomClaims(ctx).AdminID)
	if err != nil {
		res := code.GetCodeMsg(err)
		if res == nil {
			result.SetCodeError(code.RequestParamsInvalid, err)
			config.AppLoggerSugared.Debugw("Permission", zap.Any("msg", result), zap.Any("error", err))
			code.JSON(ctx, result)
			return
		}

		config.AppLoggerSugared.Debugw("Permission", zap.Any("msg", res), zap.Any("error", err))
		code.JSON(ctx, res)
		return
	}
	result.SetData(data)
	config.AppLoggerSugared.Debugw("Permission", zap.Any("msg", result))
	code.JSON(ctx, result)
	return
}
