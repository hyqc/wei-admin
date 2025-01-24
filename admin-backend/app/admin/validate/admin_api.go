package validate

import (
	"admin/pkg/govalidate"
	"admin/proto/admin_proto"
	"net/url"
)

var AdminApiReq = &AdminApiReqValidator{}

type AdminApiReqValidator struct {
}

// ListReq 接口列表参数验证
func (a *AdminApiReqValidator) ListReq(data interface{}) url.Values {
	return url.Values{}
}

func (a *AdminApiReqValidator) AddReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminApiAdd{},
			Rules: map[string]string{
				"Path":     "required,adminApiPath",
				"Key":      "required,adminApiKey",
				"Name":     "required,trimBlank",
				"Describe": "required,trimBlank",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminApiReqValidator) InfoReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminApiInfo{},
			Rules: map[string]string{
				"Id": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminApiReqValidator) EditReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminApiEdit{},
			Rules: map[string]string{
				"Id":   "required",
				"Path": "required,adminApiPath",
				"Key":  "required,adminApiKey",
				"Name": "required,trimBlank",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminApiReqValidator) EnableReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminApiEnable{},
			Rules: map[string]string{
				"Id": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminApiReqValidator) DeleteReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminApiDelete{},
			Rules: map[string]string{
				"Id": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}
