package validate

import (
	"admin/pkg/validator"
	"admin/proto/admin_proto"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

var AdminUserReq = &AdminUserReqValidator{}

type AdminUserReqValidator struct {
}

// ListReq 接口列表参数验证
func (a *AdminUserReqValidator) ListReq(data interface{}) url.Values {
	return url.Values{}
}

func (a *AdminUserReqValidator) AddReq(data interface{}) url.Values {
	rules := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("username"):        []string{"required", fmt.Sprintf("regex:%s", PatternAdminUsernameRule)},
		validator.GetValidateJsonOmitemptyTag("nickname"):        []string{"required", fmt.Sprintf("regex:%s", PatternTrimBlankStringRule)},
		validator.GetValidateJsonOmitemptyTag("password"):        []string{"required", fmt.Sprintf("regex:%s", PatternAdminPasswordRule)},
		validator.GetValidateJsonOmitemptyTag("confirmPassword"): []string{"required", fmt.Sprintf("regex:%s", PatternAdminPasswordRule)},
		validator.GetValidateJsonOmitemptyTag("email"):           []string{"email"},
		validator.GetValidateJsonOmitemptyTag("avatar"):          []string{"url"},
	}
	messages := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("username"):        []string{"required:账号不能为空", fmt.Sprintf("regex:%s", PatternAdminUsernameMsg)},
		validator.GetValidateJsonOmitemptyTag("nickname"):        []string{"required:昵称不能为空", fmt.Sprintf("regex:%s", PatternTrimBlankStringMsg)},
		validator.GetValidateJsonOmitemptyTag("password"):        []string{"required:密码不能为空", PatternAdminPasswordMsg},
		validator.GetValidateJsonOmitemptyTag("confirmPassword"): []string{"required:密码不能为空", PatternAdminPasswordMsg},
		validator.GetValidateJsonOmitemptyTag("email"):           []string{"email:邮箱格式错误"},
		validator.GetValidateJsonOmitemptyTag("avatar"):          []string{"url:无效链接地址"},
	}
	opts := govalidator.Options{
		Data:     data,
		Rules:    rules,
		Messages: messages,
	}
	res := govalidator.New(opts).ValidateStruct()
	errs := map[string][]string(res)
	if len(errs) > 0 {
		return res
	}
	tmp := data.(*admin_proto.ReqAdminUserAdd)
	if tmp.Password != tmp.ConfirmPassword {
		res["confirmPassword"] = []string{"两次输入的密码不一致"}
	}
	return res
}

func (a *AdminUserReqValidator) InfoReq(data interface{}) url.Values {
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

func (a *AdminUserReqValidator) EditReq(data interface{}) url.Values {
	rules := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("id"):              []string{"required", "min:1"},
		validator.GetValidateJsonOmitemptyTag("username"):        []string{"required", fmt.Sprintf("regex:%s", PatternAdminUsernameRule)},
		validator.GetValidateJsonOmitemptyTag("nickname"):        []string{"required", fmt.Sprintf("regex:%s", PatternTrimBlankStringRule)},
		validator.GetValidateJsonOmitemptyTag("password"):        []string{"required", fmt.Sprintf("regex:%s", PatternAdminPasswordRule)},
		validator.GetValidateJsonOmitemptyTag("confirmPassword"): []string{"required", fmt.Sprintf("regex:%s", PatternAdminPasswordRule)},
		validator.GetValidateJsonOmitemptyTag("email"):           []string{"email"},
		validator.GetValidateJsonOmitemptyTag("avatar"):          []string{"url"},
	}
	messages := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("id"):              []string{"required:ID不能为空", "min:ID无效"},
		validator.GetValidateJsonOmitemptyTag("username"):        []string{"required:账号不能为空", fmt.Sprintf("regex:%s", PatternAdminUsernameMsg)},
		validator.GetValidateJsonOmitemptyTag("nickname"):        []string{"required:昵称不能为空", fmt.Sprintf("regex:%s", PatternTrimBlankStringMsg)},
		validator.GetValidateJsonOmitemptyTag("password"):        []string{"required:密码不能为空", PatternAdminPasswordMsg},
		validator.GetValidateJsonOmitemptyTag("confirmPassword"): []string{"required:密码不能为空", PatternAdminPasswordMsg},
		validator.GetValidateJsonOmitemptyTag("email"):           []string{"email:邮箱格式错误"},
		validator.GetValidateJsonOmitemptyTag("avatar"):          []string{"url:无效链接地址"},
	}
	opts := govalidator.Options{
		Data:     data,
		Rules:    rules,
		Messages: messages,
	}
	res := govalidator.New(opts).ValidateStruct()
	errs := map[string][]string(res)
	if len(errs) > 0 {
		return res
	}
	tmp := data.(*admin_proto.ReqAdminUserAdd)
	if tmp.Password != "" && tmp.Password != tmp.ConfirmPassword {
		res["confirmPassword"] = []string{"两次输入的密码不一致"}
	}
	return res
}

func (a *AdminUserReqValidator) EnableReq(data interface{}) url.Values {
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

func (a *AdminUserReqValidator) DeleteReq(data interface{}) url.Values {
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

func (a *AdminUserReqValidator) BindRolesReq(data interface{}) url.Values {
	rules := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("adminId"): []string{"required", "min:1"},
		validator.GetValidateJsonOmitemptyTag("roleIds"): []string{"required", "min:1"},
	}
	messages := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("adminId"): []string{"required:账号ID不能为空", "min:账号ID无效"},
		validator.GetValidateJsonOmitemptyTag("roleIds"): []string{"required:角色配置不能为空", "min:角色配置不能为空"},
	}
	opts := govalidator.Options{
		Data:     data,
		Rules:    rules,
		Messages: messages,
	}
	return govalidator.New(opts).ValidateStruct()
}
