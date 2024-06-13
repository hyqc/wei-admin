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

type IMessage interface {
	SetCode(code int)
	GetCode() int
	SetMessage(msg string)
	GetMessage() string
	SetData(data interface{})
	GetData() interface{}
}

func (m *Message) SetCode(code int) {
	m.Code = code
}

func (m *Message) GetCode() int {
	return m.Code
}

func (m *Message) SetMessage(message string) {
	m.Message = message
}

func (m *Message) GetMessage() string {
	return m.Message
}

func (m *Message) SetData(data interface{}) {
	m.Data = data
}

func (m *Message) GetData() interface{} {
	return m.Data
}

func JSON(ctx *gin.Context, data IMessage) {
	base := MessageBase{
		Code:    data.GetCode(),
		Message: data.GetMessage(),
	}
	if data.GetData() == nil {
		ctx.AbortWithStatusJSON(http.StatusOK, base)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, Message{
		base, data.GetData(),
	})
	return
}
