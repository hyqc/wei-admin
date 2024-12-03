package controller

import (
	"admin/code"
	"admin/global"
	"admin/pkg/core"
	"admin/proto/code_proto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DemoController struct {
	core.Controller
}

func (d DemoController) Demo(ctx *gin.Context) {
	result := code.NewCode(code_proto.ErrorCode_Success)
	global.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	code.JSON(ctx, result)
	return
}
