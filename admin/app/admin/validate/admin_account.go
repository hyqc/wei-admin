package validate

import (
	"admin/pkg/validator"
	"admin/proto/admin_proto"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

var AdminAccountReq = &AdminAccountReqValidator{}

type AdminAccountReqValidator struct {
}

// LoginReq 登录参数验证
func (a *AdminAccountReqValidator) LoginReq(data interface{}) url.Values {
	rules := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("username"): []string{"required", fmt.Sprintf("regex:%s", PatternAdminUsernameRule)},
		validator.GetValidateJsonOmitemptyTag("password"): []string{"required", fmt.Sprintf("regex:%s", PatternAdminPasswordRule)},
	}
	messages := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("username"): []string{"required:管理员名称不能为空", PatternAdminUsernameMsg},
		validator.GetValidateJsonOmitemptyTag("password"): []string{"required:密码不能为空", PatternAdminPasswordMsg},
	}
	opts := govalidator.Options{
		Data:     data,
		Rules:    rules,
		Messages: messages,
	}
	return govalidator.New(opts).ValidateStruct()
}

func (a *AdminAccountReqValidator) AccountEditReq(data interface{}) url.Values {
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

func (a *AdminAccountReqValidator) AccountEditPasswordReq(data interface{}) url.Values {
	rules := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("password"):        []string{"required", "between:6,32"},
		validator.GetValidateJsonOmitemptyTag("oldPassword"):     []string{"required", "between:6,32"},
		validator.GetValidateJsonOmitemptyTag("confirmPassword"): []string{"required", "between:6,32"},
	}
	messages := govalidator.MapData{
		validator.GetValidateJsonOmitemptyTag("password"):        []string{"required:新密码不能为空", "between:密码长度为6-32个字符"},
		validator.GetValidateJsonOmitemptyTag("oldPassword"):     []string{"required:旧密码不能为空", "between:密码长度为6-32个字符"},
		validator.GetValidateJsonOmitemptyTag("confirmPassword"): []string{"required:确认密码不能为空", "between:密码长度为6-32个字符"},
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
	tmp := data.(*admin_proto.AccountPasswordEditReq)
	if tmp.Password != tmp.ConfirmPassword {
		res["confirmPassword"] = []string{"两次输入的密码不一致"}
	}
	if tmp.Password == tmp.OldPassword {
		res["confirmPassword"] = []string{"新旧密码不能一致"}
	}
	return res
}
