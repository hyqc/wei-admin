package code

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageBase struct {
	Code    ErrCode `json:"code"`    // 错误码
	Message string  `json:"message"` // 错误码描述
	Error   string  `json:"error"`   // 错误具体信息
}

type Message struct {
	MessageBase
	Data interface{} `json:"data"`
}

type IMessage interface {
	SetCode(code ErrCode)
	GetCode() ErrCode
	SetMessage(msg string)
	GetMessage() string
	GetError() string
	SetCodeMessage(code ErrCode, msg string)
	SetCodeError(code ErrCode, err error)
	SetData(data interface{})
	GetData() interface{}
}

func (m *Message) SetCode(code ErrCode) {
	m.Code = code
	m.Message = GetMsgByCode(code)
	m.Error = m.Message
}

func (m *Message) GetCode() ErrCode {
	return m.Code
}

func (m *Message) SetMessage(message string) {
	m.Message = message
	m.Error = m.Message
}

func (m *Message) GetMessage() string {
	return m.Message
}

func (m *Message) GetError() string {
	return m.Error
}

func (m *Message) SetCodeMessage(code ErrCode, message string) {
	m.Code = code
	m.Message = message
	m.Error = m.Message
}
func (m *Message) SetCodeError(code ErrCode, err error) {
	m.Code = code
	m.Message = GetMsgByCode(code)
	m.Error = err.Error()
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
	if debug {
		base.Error = data.GetError()
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
