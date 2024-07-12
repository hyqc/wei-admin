package validate

import (
	"admin/pkg/validator"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

var AdminPermissionReq = &AdminPermissionReqValidator{}

type AdminPermissionReqValidator struct {
}

// ListReq 权限列表参数验证
func (a *AdminPermissionReqValidator) ListReq(data interface{}) url.Values {
	return url.Values{}
}

func (a *AdminPermissionReqValidator) AddReq(data interface{}) url.Values {
	rules := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("menuId"):   []string{"required", "min:1"},
		validator.GetValidateJsonOmitemptyTag("key"):      []string{"required", fmt.Sprintf("regex:%s", PatternTrimBlankStringRule)},
		validator.GetValidateJsonOmitemptyTag("name"):     []string{"required", fmt.Sprintf("regex:%s", PatternTrimBlankStringRule)},
		validator.GetValidateJsonOmitemptyTag("describe"): []string{fmt.Sprintf("regex:%s", PatternTrimBlankStringRule)},
		validator.GetValidateJsonOmitemptyTag("type"):     []string{"required", fmt.Sprintf("regex:%s", PatternTrimBlankStringRule)},
	}
	messages := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("menuId"):   []string{"required:所属页面ID不能为空", "min:所属页面ID无效"},
		validator.GetValidateJsonOmitemptyTag("key"):      []string{"required:权限名称不能为空", PatternTrimBlankStringMsg},
		validator.GetValidateJsonOmitemptyTag("name"):     []string{"required:权限名称不能为空", PatternTrimBlankStringMsg},
		validator.GetValidateJsonOmitemptyTag("describe"): []string{PatternTrimBlankStringMsg},
		validator.GetValidateJsonOmitemptyTag("type"):     []string{"required:权限类型不能为空", PatternTrimBlankStringMsg},
	}
	opts := govalidator.Options{
		Data:     data,
		Rules:    rules,
		Messages: messages,
	}
	return govalidator.New(opts).ValidateStruct()
}

func (a *AdminPermissionReqValidator) InfoReq(data interface{}) url.Values {
	rules := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("id"): []string{"required", "min:1"},
	}
	messages := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("id"): []string{"required:权限ID不能为空", "min:权限ID无效"},
	}
	opts := govalidator.Options{
		Data:     data,
		Rules:    rules,
		Messages: messages,
	}
	return govalidator.New(opts).ValidateStruct()
}

func (a *AdminPermissionReqValidator) EditReq(data interface{}) url.Values {
	rules := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("id"):       []string{"required", "min:1"},
		validator.GetValidateJsonOmitemptyTag("menuId"):   []string{"required", "min:1"},
		validator.GetValidateJsonOmitemptyTag("key"):      []string{"required", fmt.Sprintf("regex:%s", PatternTrimBlankStringRule)},
		validator.GetValidateJsonOmitemptyTag("name"):     []string{"required", fmt.Sprintf("regex:%s", PatternTrimBlankStringRule)},
		validator.GetValidateJsonOmitemptyTag("describe"): []string{fmt.Sprintf("regex:%s", PatternTrimBlankStringRule)},
		validator.GetValidateJsonOmitemptyTag("type"):     []string{"required", fmt.Sprintf("regex:%s", PatternTrimBlankStringRule)},
	}
	messages := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("id"):       []string{"required:ID不能为空", "min:ID无效"},
		validator.GetValidateJsonOmitemptyTag("menuId"):   []string{"required:所属页面ID不能为空", "min:所属页面ID无效"},
		validator.GetValidateJsonOmitemptyTag("key"):      []string{"required:权限名称不能为空", PatternTrimBlankStringMsg},
		validator.GetValidateJsonOmitemptyTag("name"):     []string{"required:权限名称不能为空", PatternTrimBlankStringMsg},
		validator.GetValidateJsonOmitemptyTag("describe"): []string{PatternTrimBlankStringMsg},
		validator.GetValidateJsonOmitemptyTag("type"):     []string{"required:权限类型不能为空", PatternTrimBlankStringMsg},
	}
	opts := govalidator.Options{
		Data:     data,
		Rules:    rules,
		Messages: messages,
	}
	return govalidator.New(opts).ValidateStruct()
}

func (a *AdminPermissionReqValidator) EnableReq(data interface{}) url.Values {
	rules := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("id"):      []string{"required", "min:1"},
		validator.GetValidateJsonOmitemptyTag("enabled"): []string{"bool"},
	}
	messages := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("id"):      []string{"required:ID不能为空", "min:ID无效"},
		validator.GetValidateJsonOmitemptyTag("enabled"): []string{"bool:类型错误"},
	}
	opts := govalidator.Options{
		Data:     data,
		Rules:    rules,
		Messages: messages,
	}
	return govalidator.New(opts).ValidateStruct()
}

func (a *AdminPermissionReqValidator) DeleteReq(data interface{}) url.Values {
	rules := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("id"): []string{"required", "min:1"},
	}
	messages := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("id"): []string{"required:ID不能为空", "min:ID无效"},
	}
	opts := govalidator.Options{
		Data:     data,
		Rules:    rules,
		Messages: messages,
	}
	return govalidator.New(opts).ValidateStruct()
}
