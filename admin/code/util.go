package code

import "admin/pkg/response"

func NewCode(code int) response.Message {
	return response.Message{
		MessageBase: response.MessageBase{
			Code:    code,
			Message: Msg[code],
		},
	}
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
