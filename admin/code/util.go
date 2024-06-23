package code

import (
	"admin/pkg/core"
	"admin/proto/code_proto"
	"errors"
)

func newCodeBase(code code_proto.ErrorCode) MessageBase {
	return MessageBase{
		Code: code,
		Msg:  GetMsgByCode(code),
	}
}

func NewCode(code code_proto.ErrorCode) IMessage {
	return &Message{
		MessageBase: newCodeBase(code),
	}
}

func NewCodeError(code code_proto.ErrorCode, err error) IMessage {
	msg := NewCode(code)
	if err != nil {
		msg.SetError(err)
	} else {
		msg.SetError(errors.New(msg.GetMsg()))
	}
	return msg
}

func ResponseData(code code_proto.ErrorCode, data interface{}, err error) core.ResponseData {
	msg := core.ResponseData{
		Code:    int(code),
		Message: GetMsgByCode(code),
		Error:   err.Error(),
		Data:    data,
	}
	if err != nil {
		msg.Error = err.Error()
	}
	return msg
}
