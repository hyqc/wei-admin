package validate

import (
	"admin/pkg/validator"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

var ValidateAccount = AccountValidator{}

type AccountValidator struct {
}

// LoginReq 登录参数验证
func (AccountValidator) LoginReq(data interface{}) url.Values {
	rules := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("username"): []string{"required", "between:1,32"},
		validator.GetValidateJsonOmitemptyTag("password"): []string{"required", "between:6,64"},
	}
	messages := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("username"): []string{"required:管理员名称不能为空", "between:管理员名称长度为1-32个字符"},
		validator.GetValidateJsonOmitemptyTag("password"): []string{"required:密码不能为空", "between:密码长度6-64个字符"},
	}
	opts := govalidator.Options{
		Data:     data,
		Rules:    rules,
		Messages: messages,
	}
	return govalidator.New(opts).ValidateStruct()
}

func (v AccountValidator) AccountEditReq(data interface{}) url.Values {
	rules := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("nickname"): []string{"required", "between:1,32"},
		validator.GetValidateJsonOmitemptyTag("avatar"):   []string{"url"},
	}
	messages := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("nickname"): []string{"required:昵称不能为空", "between:昵称长度为1-32个字符"},
		validator.GetValidateJsonOmitemptyTag("avatar"):   []string{"url:不是有效的地址"},
	}
	opts := govalidator.Options{
		Data:     data,
		Rules:    rules,
		Messages: messages,
	}
	return govalidator.New(opts).ValidateStruct()
}
