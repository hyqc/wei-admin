package validate

import (
	"admin/pkg/govalidate"
	"admin/proto/admin_proto"
)

var AdminRoleReq = &AdminRoleReqValidator{}

type AdminRoleReqValidator struct {
}

// ListReq 接口列表参数验证
func (a *AdminRoleReqValidator) ListReq(data interface{}) error {
	return nil
}

func (a *AdminRoleReqValidator) AddReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminRoleAdd{},
			Rules: map[string]string{
				"Name":     "required,trimBlank",
				"Describe": "omitempty,trimBlank",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminRoleReqValidator) InfoReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminRoleInfo{},
			Rules: map[string]string{
				"Id": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminRoleReqValidator) EditReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminRoleEdit{},
			Rules: map[string]string{
				"Id":       "required",
				"Name":     "required,trimBlank",
				"Describe": "omitempty,trimBlank",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminRoleReqValidator) EnableReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminRoleEnable{},
			Rules: map[string]string{
				"Id": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminRoleReqValidator) DeleteReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminRoleDelete{},
			Rules: map[string]string{
				"Id": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminRoleReqValidator) RolePermissionsReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminRolePermissions{},
			Rules: map[string]string{
				"Id": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminRoleReqValidator) RoleBindPermissionsReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminRoleBindPermissions{},
			Rules: map[string]string{
				"Id":            "required",
				"PermissionIds": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}
