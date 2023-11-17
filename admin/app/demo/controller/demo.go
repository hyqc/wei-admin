package controller

import (
	"admin/code"
	"admin/config"
	"admin/pkg/core"
	"admin/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DemoController struct {
	core.Controller
}

func (d DemoController) Demo(ctx *gin.Context) {
	result := code.NewCode(code.Success)
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	response.JSON(ctx, result)
	return
}
