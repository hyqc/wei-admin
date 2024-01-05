package validator

import (
	"admin/pkg/validate"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

type AccountValidator struct {
}

func (AccountValidator) Login(data interface{}) url.Values {
	rules := govalidator.MapData{
		validate.GetValidateJSONTag("username"): []string{"required", "between:1,32"},
		validate.GetValidateJSONTag("pwd"):      []string{"required", "between:6,64"},
	}
	messages := govalidator.MapData{
		validate.GetValidateJSONTag("username"): []string{"required:管理员名称不能为空", "between:管理员名称长度为1-32个字符"},
		validate.GetValidateJSONTag("pwd"):      []string{"required:密码不能为空", "between:密码长度6-64个字符"},
	}
	opts := govalidator.Options{
		Data:     data,
		Rules:    rules,
		Messages: messages,
	}
	return govalidator.New(opts).ValidateStruct()
}
