package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IController interface {
	ResponseOk(ctx *gin.Context, data ResponseData)
	ResponseStatus(ctx *gin.Context, status int, data ResponseData)
}

type Controller struct {
}

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func handleResponseResult(data ResponseData) map[string]interface{} {
	result := make(map[string]interface{})
	result["code"] = data.Code
	result["message"] = data.Message
	if data.Data != nil {
		result["data"] = data.Data
	}
	return result
}

func (b Controller) ResponseOk(ctx *gin.Context, data ResponseData) {
	ctx.AbortWithStatusJSON(http.StatusOK, handleResponseResult(data))
}

func (b Controller) ResponseStatus(ctx *gin.Context, status int, data ResponseData) {
	ctx.AbortWithStatusJSON(status, handleResponseResult(data))
}
