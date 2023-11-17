package controller

import (
	"admin/code"
	"admin/config"
	"admin/pkg/core"
	"admin/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AccountController struct {
	core.Controller
}

func (a AccountController) Login(ctx *gin.Context) {
	result := code.NewCode(code.Success)
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	response.JSON(ctx, result)
	return
}

func (a AccountController) Register(ctx *gin.Context) {
	result := code.NewCode(code.Success)
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	response.JSON(ctx, result)
	return
}
