package controller

import (
	"admin/app/admin/logic"
	"admin/code"
	"admin/config"
	"admin/pkg/core"
	adminproto "admin/proto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type APIController struct {
	core.Controller
}

var (
	apiLogic = logic.NewAdminAPILogic()
)

func (APIController) List(ctx *gin.Context) {
	params := &adminproto.ApiListReq{}
	result := code.NewCode(code.Success)
	data, err := apiLogic.List(ctx, params)
	if err != nil {
		res := code.GetCodeMsg(err)
		if res == nil {
			result.SetCodeError(code.RequestParamsInvalid, err)
			config.AppLoggerSugared.Debugw("List", zap.Any("msg", result), zap.Any("error", err))
			code.JSON(ctx, result)
			return
		}

		config.AppLoggerSugared.Debugw("List", zap.Any("msg", res), zap.Any("error", err))
		code.JSON(ctx, res)
		return
	}
	result.SetData(data)
	config.AppLoggerSugared.Debugw("List", zap.Any("msg", result))
	code.JSON(ctx, result)
	return
}

func (APIController) All(ctx *gin.Context) {

}

func (APIController) Add(ctx *gin.Context) {

}

func (APIController) Info(ctx *gin.Context) {

}

func (APIController) Edit(ctx *gin.Context) {

}

func (APIController) Enable(ctx *gin.Context) {

}

func (APIController) Delete(ctx *gin.Context) {

}
