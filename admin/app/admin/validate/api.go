package validate

import (
	"admin/pkg/validator"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

var APIRequest = &APIRequestValidator{}

type APIRequestValidator struct {
}

// ListReq 接口列表参数验证
func (a *APIRequestValidator) ListReq(data interface{}) url.Values {
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
