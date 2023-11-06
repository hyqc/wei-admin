package controller

import (
	"ginweb/codes"
	"ginweb/config"
	"ginweb/pkg/core"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DemoController struct {
	core.Controller
}

func (d DemoController) Demo(ctx *gin.Context) {
	result := core.ResponseData{}
	result.Code = codes.Success
	result.Message = codes.Message[codes.Success]
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	d.ResponseOk(ctx, result)
	return
}
