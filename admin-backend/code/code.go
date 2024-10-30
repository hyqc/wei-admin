package code

import "admin/proto/code_proto"

type Language int

const (
	ZH Language = iota
	EN
)

var (
	language      = ZH
	debug         = true
	defaultErrMsg = "失败"
	msg           = map[Language]map[code_proto.ErrorCode]string{
		ZH: zhMsg,
		EN: enMsg,
	}
)

func SetLanguage(c Language) {
	language = c
}

func SetDebug(d bool) {
	debug = d
}

func SetDefaultErrMsg(msg string) {
	defaultErrMsg = msg
}

func GetMsgByCode(code code_proto.ErrorCode) string {
	if val, ok := msg[language]; ok {
		if v, ok := val[code]; ok {
			return v
		}
	}
	return defaultErrMsg
}
