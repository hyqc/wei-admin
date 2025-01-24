package validate

import (
	"admin/pkg/govalidate"
	"admin/proto/admin_proto"
)

var AdminPermissionReq = &AdminPermissionReqValidator{}

type AdminPermissionReqValidator struct {
}

// ListReq 权限列表参数验证
func (a *AdminPermissionReqValidator) ListReq(data interface{}) error {
	return nil
}

func (a *AdminPermissionReqValidator) AddReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminPermissionAdd{},
			Rules: map[string]string{
				"MenuId":   "required",
				"Type":     "required,trimBlank",
				"Key":      "required,trimBlank",
				"Name":     "required,trimBlank",
				"Describe": "required,trimBlank",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminPermissionReqValidator) InfoReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminPermissionInfo{},
			Rules: map[string]string{
				"Id": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminPermissionReqValidator) EditReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminPermissionEdit{},
			Rules: map[string]string{
				"Id":       "required",
				"MenuId":   "required",
				"Type":     "required,trimBlank",
				"Key":      "required,trimBlank",
				"Name":     "required,trimBlank",
				"Describe": "required,trimBlank",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminPermissionReqValidator) EnableReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminPermissionEnable{},
			Rules: map[string]string{
				"Id": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminPermissionReqValidator) DeleteReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminPermissionDelete{},
			Rules: map[string]string{
				"Id": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminPermissionReqValidator) BindAPIReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminPermissionBindApis{},
			Rules: map[string]string{
				"PermissionId": "required",
				"ApiIds":       "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminPermissionReqValidator) PermissionBindMenuReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminPermissionBindMenu{},
			Rules: map[string]string{
				"MenuId":      "required",
				"Permissions": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminPermissionReqValidator) UnBindAPIReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminPermissionUnBindApi{},
			Rules: map[string]string{
				"PermissionId": "required",
				"ApiId":        "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}
