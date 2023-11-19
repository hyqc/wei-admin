package code

type Language int

const (
	ZH Language = iota
	EN
)

var (
	language = ZH
	msg      = map[Language]map[int]string{
		ZH: zhMsg,
		EN: enMsg,
	}
)

const (
	Success = 0
	Error   = 1

	ReadContextRequestBodyFailed = 100001

	AuthTokenFailed         = 200001 // 鉴权
	AuthTokenInvalid        = 200002 // BearerToken无效
	AuthTokenInspectInvalid = 200003 // BearerToken无效
	AuthTokenInfoInvalid    = 200004 // Token信息无效

	RequestBodyInvalid   = 300001 // 请求参数无效
	RequestQueryInvalid  = 300002 // 请求参数无效
	RequestParamsInvalid = 300003 // 请求参数无效

	AdminAccountPasswordInvalid = 400001 // 密码错误
)

var zhMsg = map[int]string{
	Success:                      "成功",
	Error:                        "请求失败",
	ReadContextRequestBodyFailed: "读取请求体数据失败",
	AuthTokenFailed:              "未登录或登录令牌已过期",
	AuthTokenInvalid:             "登录令牌无效",
	AuthTokenInspectInvalid:      "登录令牌检查失败",
	AuthTokenInfoInvalid:         "登录令牌信息无效",

	RequestBodyInvalid:   "请求体参数无效",
	RequestQueryInvalid:  "查询参数无效",
	RequestParamsInvalid: "请求参数无效",

	AdminAccountPasswordInvalid: "密码错误",
}

var enMsg = map[int]string{
	Success:                      "success",
	Error:                        "error",
	ReadContextRequestBodyFailed: "read context request body failed",
	AuthTokenFailed:              "token expired or not login",
	AuthTokenInvalid:             "token invalid",
	AuthTokenInspectInvalid:      "token inspect invalid",
	AuthTokenInfoInvalid:         "token info invalid",

	RequestBodyInvalid:   "request body invalid",
	RequestQueryInvalid:  "request query invalid",
	RequestParamsInvalid: "request params invalid",

	AdminAccountPasswordInvalid: "password invalid",
}

func SetLanguage(c Language) {
	language = c
}

func Msg(code int) string {
	return msg[language][code]
}
