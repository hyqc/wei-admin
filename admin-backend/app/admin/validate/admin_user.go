package validate

import (
	"admin/pkg/govalidate"
	"admin/proto/admin_proto"
)

var AdminUserReq = &AdminUserReqValidator{}

type AdminUserReqValidator struct {
}

// ListReq 接口列表参数验证
func (a *AdminUserReqValidator) ListReq(data any) error {
	return nil
}

func (a *AdminUserReqValidator) AddReq(data any) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminUserAdd{},
			Rules: map[string]string{
				"Username":        "required,adminname",
				"Nickname":        "required,trimBlank",
				"Password":        "required,adminpwd",
				"ConfirmPassword": "required,eqfield=Password,adminpwd",
				"Email":           "omitempty,email",
				"Avatar":          "omitempty,url",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminUserReqValidator) InfoReq(data any) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminUserInfo{},
			Rules: map[string]string{
				"AdminId": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminUserReqValidator) EditReq(data any) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminUserEdit{},
			Rules: map[string]string{
				"AdminId":  "required",
				"Username": "required,adminname",
				"Nickname": "required,trimBlank",
				"Email":    "omitempty,email",
				"Avatar":   "omitempty,url",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminUserReqValidator) EditPasswordReq(data any) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminUserEditPassword{},
			Rules: map[string]string{
				"AdminId":         "required",
				"Password":        "required,adminpwd",
				"ConfirmPassword": "required,eqfield=Password,adminpwd",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminUserReqValidator) EnableReq(data any) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminUserEnabled{},
			Rules: map[string]string{
				"AdminId": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminUserReqValidator) DeleteReq(data any) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminUserDelete{},
			Rules: map[string]string{
				"AdminId": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminUserReqValidator) BindRolesReq(data any) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAdminUserBindRoles{},
			Rules: map[string]string{
				"AdminId": "required",
				"RoleIds": "required",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}
