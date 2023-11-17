package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Message struct {
	MessageBase
	Data interface{} `json:"data"`
}

func JSON(ctx *gin.Context, data Message) {
	base := MessageBase{
		Code:    data.Code,
		Message: data.Message,
	}
	if data.Data == nil {
		ctx.AbortWithStatusJSON(http.StatusOK, base)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, Message{
		base, data.Data,
	})
	return
}
