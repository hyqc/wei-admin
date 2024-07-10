package validate

import "net/url"

var AdminPermissionReq = &AdminPermissionReqValidator{}

type AdminPermissionReqValidator struct {
}

// ListReq 权限列表参数验证
func (a *AdminPermissionReqValidator) ListReq(data interface{}) url.Values {
	return url.Values{}
}
