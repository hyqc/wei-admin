package validate

import (
	"admin/pkg/govalidate"
	"admin/proto/admin_proto"
)

var AdminMenuReq = &AdminMenuReqValidator{}

type AdminMenuReqValidator struct {
}

func (a *AdminMenuReqValidator) AddReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminMenuAdd{},
			Rules: map[string]string{
				"ParentId": "min=0",
				"Path":     "required,adminApiPath",
				"Key":      "required,adminApiKey",
				"Name":     "required,trimBlank",
				"Describe": "required,trimBlank",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminMenuReqValidator) InfoReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminMenuInfo{},
			Rules: map[string]string{
				"MenuId": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminMenuReqValidator) EditReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminMenuEdit{},
			Rules: map[string]string{
				"Id":       "required",
				"ParentId": "min=0",
				"Path":     "required,adminApiPath",
				"Key":      "required,adminApiKey",
				"Name":     "required,trimBlank",
				"Describe": "required,trimBlank",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminMenuReqValidator) EnableReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminMenuEnable{},
			Rules: map[string]string{
				"MenuId": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminMenuReqValidator) ShowReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminMenuShow{},
			Rules: map[string]string{
				"MenuId": "required",
				"Field":  "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminMenuReqValidator) DeleteReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminMenuDelete{},
			Rules: map[string]string{
				"MenuId": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminMenuReqValidator) PermissionsReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminMenuPermissions{},
			Rules: map[string]string{
				"MenuId": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminMenuReqValidator) PagesReq(data interface{}) error {
	return nil
}
