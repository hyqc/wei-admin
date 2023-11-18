package controller

import (
	"admin/app/admin/validater"
	"admin/code"
	"admin/config"
	"admin/pkg/core"
	"admin/pkg/response"
	"admin/pkg/validate"
	"admin/proto/admin_account"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AccountController struct {
	core.Controller
}

// Login 登录
// @Summary 登录后台
// @Description 登录后台
// @Tags 账号相关接口
// @Accept application/json
// @Produce application/json
// @Param object query models.ParamPostList true "请求参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /admin/login [post]
func (AccountController) Login(ctx *gin.Context) {
	params := &admin_account.LoginReq{}
	if err := validate.Validate(ctx, params, validater.AccountValidatorHandler{}.Login); err != nil {
		result := code.NewCode(code.RequestParamsInvalid)
		config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
		response.JSON(ctx, result)
		return
	}
	result := code.NewCode(code.Success)
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	response.JSON(ctx, result)
	return
}

func (AccountController) Register(ctx *gin.Context) {
	result := code.NewCode(code.Success)
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	response.JSON(ctx, result)
	return
}

func (AccountController) Detail(ctx *gin.Context) {
	result := code.NewCode(code.Success)
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	response.JSON(ctx, result)
	return
}

func (AccountController) Edit(ctx *gin.Context) {
	result := code.NewCode(code.Success)
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	response.JSON(ctx, result)
	return
}

func (AccountController) Password(ctx *gin.Context) {
	result := code.NewCode(code.Success)
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	response.JSON(ctx, result)
	return
}

func (AccountController) Menu(ctx *gin.Context) {
	result := code.NewCode(code.Success)
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	response.JSON(ctx, result)
	return
}

func (AccountController) Permission(ctx *gin.Context) {
	result := code.NewCode(code.Success)
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	response.JSON(ctx, result)
	return
}
