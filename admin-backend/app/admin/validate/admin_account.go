package validate

import (
	"admin/pkg/govalidate"
	"admin/proto/admin_proto"
)

var AdminAccountReq = &AdminAccountReqValidator{}

type AdminAccountReqValidator struct {
}

// LoginReq 登录参数验证
func (a *AdminAccountReqValidator) LoginReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqLogin{},
			Rules: map[string]string{
				"Username": "required,adminname",
				"Password": "required,adminpwd",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminAccountReqValidator) AccountEditReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAccountEdit{},
			Rules: map[string]string{
				"Nickname": "required,min=1,max=32",
				"Avatar":   "url",
				"Email":    "email",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}

func (a *AdminAccountReqValidator) AccountEditPasswordReq(data interface{}) error {
	rules := govalidate.Rules{
		{
			Type: admin_proto.ReqAccountPasswordEdit{},
			Rules: map[string]string{
				"Password":        "required,adminpwd",
				"ConfirmPassword": "required,adminpwd,eqfield=Password",
				"OldPassword":     "required,adminpwd",
			},
		},
	}
	return govalidate.ValidateStructWithRules(data, rules)
}
