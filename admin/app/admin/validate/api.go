package validate

import (
	"net/url"
)

var APIReq = &APIReqValidator{}

type APIReqValidator struct {
}

// ListReq 接口列表参数验证
func (a *APIReqValidator) ListReq(data interface{}) url.Values {
	return url.Values{}
}
