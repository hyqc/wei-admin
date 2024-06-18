package code

type Language int

const (
	ZH Language = iota
	EN
)

type ErrCode int

var (
	language      = ZH
	debug         = true
	defaultErrMsg = "失败"
	msg           = map[Language]map[ErrCode]string{
		ZH: zhMsg,
		EN: enMsg,
	}
)

const (
	Success ErrCode = 0
	Error   ErrCode = 1

	ReadContextRequestBodyFailed ErrCode = 100001

	AuthTokenFailed         ErrCode = 200001 // 鉴权
	AuthTokenInvalid        ErrCode = 200002 // BearerToken无效
	AuthTokenInspectInvalid ErrCode = 200003 // BearerToken无效
	AuthTokenInfoInvalid    ErrCode = 200004 // Token信息无效

	RequestBodyInvalid   ErrCode = 300001 // 请求参数无效
	RequestQueryInvalid  ErrCode = 300002 // 请求参数无效
	RequestParamsInvalid ErrCode = 300003 // 请求参数无效

	AdminAccountPasswordInvalid ErrCode = 400001 // 密码错误
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

func GetMsgByCode(code ErrCode) string {
	if val, ok := msg[language]; ok {
		if v, ok := val[code]; ok {
			return v
		}
	}
	return defaultErrMsg
}
