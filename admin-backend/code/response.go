package code

import (
	"admin/proto/code_proto"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageBase struct {
	Code   code_proto.ErrorCode `json:"code"`   // 错误码
	Msg    string               `json:"msg"`    // 错误码描述
	Reason string               `json:"reason"` // 错误具体信息
}

type Message struct {
	MessageBase
	Data any `json:"data"`
}

type IMessage interface {
	SetCode(code code_proto.ErrorCode)
	GetCode() code_proto.ErrorCode
	SetMsg(msg string)
	GetMsg() string
	GetError() string
	SetCodeMsg(code code_proto.ErrorCode, msg error)
	SetCodeError(code code_proto.ErrorCode, err error)
	SetData(data any)
	GetData() any
	Error() string
	SetError(err error)
	SetMessage(err IMessage)
	String() string
}

func NewMessage() IMessage {
	return &Message{
		MessageBase: MessageBase{},
		Data:        nil,
	}
}

func (m *Message) SetCode(code code_proto.ErrorCode) {
	m.Code = code
	m.Msg = GetMsgByCode(code)
	m.Reason = m.Msg
}

func (m *Message) GetCode() code_proto.ErrorCode {
	return m.Code
}

func (m *Message) SetMsg(message string) {
	m.Msg = message
	m.Reason = m.Msg
}

func (m *Message) GetMsg() string {
	return m.Msg
}

func (m *Message) GetError() string {
	return m.Reason
}

func (m *Message) SetCodeMsg(code code_proto.ErrorCode, err error) {
	m.Code = code
	m.Msg = err.Error()
	m.Reason = m.Msg
}
func (m *Message) SetCodeError(code code_proto.ErrorCode, err error) {
	m.Code = code
	m.Msg = GetMsgByCode(code)
	m.Reason = err.Error()
}

func (m *Message) SetData(data any) {
	m.Data = data
}

func (m *Message) GetData() any {
	return m.Data
}

func (m *Message) Error() string {
	if m.Reason != "" {
		return m.Reason
	}
	return m.Msg
}

func (m *Message) SetError(err error) {
	m.Reason = err.Error()
}

func (m *Message) SetMessage(err IMessage) {
	m.Code = err.GetCode()
	m.Msg = err.GetMsg()
	m.Reason = err.GetError()
	m.Data = err.GetData()
}

func (m *Message) String() string {
	body, _ := json.Marshal(m)
	return string(body)
}

func JSON(ctx *gin.Context, data IMessage) {
	base := MessageBase{
		Code: data.GetCode(),
		Msg:  data.GetMsg(),
	}
	if debug {
		base.Reason = data.GetError()
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
