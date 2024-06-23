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
	Code    int         `json:"code_proto"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func handleResponseResult(data ResponseData) map[string]interface{} {
	result := make(map[string]interface{})
	result["code_proto"] = data.Code
	result["message"] = data.Message
	result["error"] = data.Error
	if data.Data != nil {
		result["data"] = data.Data
	}
	return result
}

func ResponseOk(ctx *gin.Context, data ResponseData) {
	ctx.AbortWithStatusJSON(http.StatusOK, handleResponseResult(data))
}

func ResponseStatus(ctx *gin.Context, status int, data ResponseData) {
	ctx.AbortWithStatusJSON(status, handleResponseResult(data))
}

func (b Controller) ResponseOk(ctx *gin.Context, data ResponseData) {
	ctx.AbortWithStatusJSON(http.StatusOK, handleResponseResult(data))
}

func (b Controller) ResponseStatus(ctx *gin.Context, status int, data ResponseData) {
	ctx.AbortWithStatusJSON(status, handleResponseResult(data))
}
