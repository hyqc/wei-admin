package validate

import (
	"admin/pkg/govalidate"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"regexp"
)

func init() {
	err := govalidate.Init(zh.New(), func(valid *validator.Validate, trans ut.Translator) error {
		return zhTrans.RegisterDefaultTranslations(valid, trans)
	})
	if err != nil {
		panic(err)
	}

	type CustomRule struct {
		Name    string `json:"name"`
		Msg     string `json:"msg"`
		Pattern string `json:"pattern"`
	}

	pattenrs := []CustomRule{
		{
			Name:    "adminname",
			Msg:     "{0}必须是数字字母组合",
			Pattern: PatternAdminUsernameRule,
		},
		{
			Name:    "adminpwd",
			Msg:     "{0}必须是6-32位非空字符串",
			Pattern: PatternAdminPasswordRule,
		},
		{
			Name:    "adminApiKey",
			Msg:     "{0}不是有效路由键名",
			Pattern: PatternAdminApiKeyRule,
		},
		{
			Name:    "adminApiPath",
			Msg:     "{0}不是有效路由",
			Pattern: PatternAdminApiPathRule,
		},
		{
			Name:    "phone",
			Msg:     "{0}手机号格式错误",
			Pattern: PatternPhoneRule,
		},
		{
			Name:    "trimBlank",
			Msg:     "{0}不是有效字符串",
			Pattern: PatternTrimBlankStringRule,
		},
	}

	for _, val := range pattenrs {
		item := val
		err = govalidate.Validator.RegisterValidation(item.Name, func(fl validator.FieldLevel) bool {
			return regexp.MustCompile(item.Pattern).MatchString(fl.Field().String())
		}, true)
		if err != nil {
			panic(err)
		}

		err = govalidate.Validator.RegisterTranslation(item.Name, govalidate.Translator, func(ut ut.Translator) error {
			return ut.Add(item.Name, item.Msg, true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(item.Name, fe.Field(), fe.Param())
			return t
		})
		if err != nil {
			panic(err)
		}
	}
}
