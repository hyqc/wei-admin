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
	if err := validator.Validate(ctx, params, validate.ValidateAccount.LoginReq); err != nil {
		result.SetCodeError(code.RequestParamsInvalid, err)
		config.AppLoggerSugared.Debugw("info", zap.Any("msg", result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	data, err := accountLogic.Login(ctx, params, ctx.ClientIP())
	if err != nil {
		res := code.GetCodeMsg(err)
		if res == nil {
			result.SetCodeError(code.RequestParamsInvalid, err)
			config.AppLoggerSugared.Debugw("info", zap.Any("msg", result), zap.Any("error", err))
			code.JSON(ctx, result)
			return
		}

		config.AppLoggerSugared.Debugw("info", zap.Any("msg", res), zap.Any("error", err))
		code.JSON(ctx, res)
		return
	}
	result.SetData(data)

	config.AppLoggerSugared.Debugw("info", zap.Any("msg", result))
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
			config.AppLoggerSugared.Debugw("info", zap.Any("msg", result), zap.Any("error", err))
			code.JSON(ctx, result)
			return
		}

		config.AppLoggerSugared.Debugw("info", zap.Any("msg", res), zap.Any("error", err))
		code.JSON(ctx, res)
		return
	}

	result.SetData(data)
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	code.JSON(ctx, result)
	return
}

// Edit 编辑账号
func (AccountController) Edit(ctx *gin.Context) {
	params := &adminproto.AccountEditReq{}
	result := code.NewCode(code.Success)
	if err := validator.Validate(ctx, params, validate.ValidateAccount.AccountEditReq); err != nil {
		result.SetCode(code.RequestParamsInvalid)
		config.AppLoggerSugared.Debugw("info", zap.Any("msg", result), zap.Any("error", err))
		code.JSON(ctx, result)
		return
	}
	err := accountLogic.Edit(ctx, params)
	if err != nil {
		res := code.GetCodeMsg(err)
		if res == nil {
			result.SetCode(code.RequestParamsInvalid)
			config.AppLoggerSugared.Debugw("info", zap.Any("msg", result), zap.Any("error", err))
			code.JSON(ctx, result)
			return
		}

		config.AppLoggerSugared.Debugw("info", zap.Any("msg", res), zap.Any("error", err))
		code.JSON(ctx, res)
		return
	}
	config.AppLoggerSugared.Debugw("info", zap.Any("msg", result))
	code.JSON(ctx, result)
	return
}

func (AccountController) Password(ctx *gin.Context) {
	result := code.NewCode(code.Success)
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	code.JSON(ctx, result)
	return
}

func (AccountController) Menu(ctx *gin.Context) {
	result := code.NewCode(code.Success)
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	code.JSON(ctx, result)
	return
}

func (AccountController) Permission(ctx *gin.Context) {
	result := code.NewCode(code.Success)
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	code.JSON(ctx, result)
	return
}
