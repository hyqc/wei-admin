package code

import (
	"admin/pkg/core"
	"encoding/json"
	"errors"
)

type Err error

func NewCode(code ErrCode) IMessage {
	return &Message{
		MessageBase: newCodeBase(code),
	}
}

func newCodeBase(code ErrCode) MessageBase {
	s := GetMsgByCode(code)
	return MessageBase{
		Code:    code,
		Message: s,
		Error:   s,
	}
}

func NewCodeError(code ErrCode) Err {
	msg := newCodeBase(code)
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return errors.New(string(body))
}

func GetCodeMsg(err error) IMessage {
	var e Err
	switch {
	case errors.As(err, &e):
		data := &Message{}
		if e := json.Unmarshal([]byte(err.Error()), data); e != nil {
			return nil
		}
		return data
	}
	return nil
}

func NewCodeMsg(code ErrCode, msg string) IMessage {
	return &Message{
		MessageBase: MessageBase{
			Code:    code,
			Message: msg,
		},
	}
}

func NewCodeMsgData(code ErrCode, msg string, data interface{}) IMessage {
	return &Message{
		MessageBase: MessageBase{
			Code:    code,
			Message: msg,
		},
		Data: data,
	}
}

func ResponseData(code ErrCode, data interface{}) core.ResponseData {
	msg := GetMsgByCode(code)
	return core.ResponseData{
		Code:    int(code),
		Message: msg,
		Error:   msg,
		Data:    data,
	}
}
