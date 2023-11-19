package code

import (
	"admin/pkg/response"
	"encoding/json"
	"errors"
)

type Err error

func NewCode(code int) response.Message {
	return response.Message{
		MessageBase: NewCodeBase(code),
	}
}

func NewCodeBase(code int) response.MessageBase {
	return response.MessageBase{
		Code:    code,
		Message: Msg(code),
	}
}

func NewCodeError(code int) Err {
	msg := NewCodeBase(code)
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return errors.New(string(body))
}

func GetCodeMsg(err error) *response.Message {
	switch err.(type) {
	case Err:
		data := &response.Message{}
		if e := json.Unmarshal([]byte(err.Error()), data); e != nil {
			return nil
		}
		return data
	}
	return nil
}

func NewCodeMsg(code int, msg string) response.Message {
	return response.Message{
		MessageBase: response.MessageBase{
			Code:    code,
			Message: msg,
		},
	}
}

func NewCodeMsgData(code int, msg string, data interface{}) response.Message {
	return response.Message{
		MessageBase: response.MessageBase{
			Code:    code,
			Message: msg,
		},
		Data: data,
	}
}
