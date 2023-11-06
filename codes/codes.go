package codes

const (
	Success = 0
	Error   = 1
)

var Message map[int]string = map[int]string{
	Success: "success",
	Error:   "error",
}

var MessageCN map[int]string = map[int]string{
	Success: "成功",
	Error:   "请求失败",
}
