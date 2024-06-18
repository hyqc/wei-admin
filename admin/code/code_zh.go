package code

var zhMsg = map[ErrCode]string{
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
