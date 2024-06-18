package code

var enMsg = map[ErrCode]string{
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

	AdminAccountPasswordInvalid: "pwd invalid",
}
