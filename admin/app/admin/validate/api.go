package validate

import (
	"admin/pkg/validator"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

var APIReq = &APIReqValidator{}

type APIReqValidator struct {
}

// ListReq 接口列表参数验证
func (a *APIReqValidator) ListReq(data interface{}) url.Values {
	return url.Values{}
}

func (a *APIReqValidator) AddReq(data interface{}) url.Values {
	rules := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("path"):     []string{"required", fmt.Sprintf("regex:%s", PatternAdminApiPathRule)},
		validator.GetValidateJsonOmitemptyTag("key"):      []string{"required", fmt.Sprintf("regex:%s", PatternAdminApiKeyRule)},
		validator.GetValidateJsonOmitemptyTag("name"):     []string{"required", fmt.Sprintf("regex:%s", PatternTrimBlankStringRule)},
		validator.GetValidateJsonOmitemptyTag("describe"): []string{"required", fmt.Sprintf("regex:%s", PatternTrimBlankStringRule)},
	}
	messages := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("path"):     []string{"required:接口路由不能为空", PatternAdminApiPathMsg},
		validator.GetValidateJsonOmitemptyTag("key"):      []string{"required:接口路由键名不能为空", PatternAdminApiKeyMsg},
		validator.GetValidateJsonOmitemptyTag("name"):     []string{"required:接口名称不能为空", PatternTrimBlankStringMsg},
		validator.GetValidateJsonOmitemptyTag("describe"): []string{PatternTrimBlankStringMsg},
	}
	opts := govalidator.Options{
		Data:     data,
		Rules:    rules,
		Messages: messages,
	}
	return govalidator.New(opts).ValidateStruct()
}

func (a *APIReqValidator) InfoReq(data interface{}) url.Values {
	rules := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("id"): []string{"required", "min:1"},
	}
	messages := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("id"): []string{"required:接口ID不能为空", "min:接口ID不能小于1"},
	}
	opts := govalidator.Options{
		Data:     data,
		Rules:    rules,
		Messages: messages,
	}
	return govalidator.New(opts).ValidateStruct()
}
