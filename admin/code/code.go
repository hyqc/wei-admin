package code

const (
	Success = 0
	Error   = 1

	ReadContextRequestBodyFailed = 100001
)

var Msg = map[int]string{
	Success:                      "成功",
	Error:                        "请求失败",
	ReadContextRequestBodyFailed: "读取请求体数据失败",
}

var MsgEN = map[int]string{
	Success:                      "success",
	Error:                        "error",
	ReadContextRequestBodyFailed: "read context request body failed",
}
