package global

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
)

// 自定义验证规则注册
type customRegexpRule struct {
	Name    string `json:"name"`
	Msg     string `json:"msg"`
	Pattern string `json:"pattern"`
}

var (
	// TranslatorLocal 本地翻译器
	TranslatorLocal = zh.New()
	// ValidatorTagNameValue 验证器的表前面tag值
	ValidatorTagNameValue = "label"
	UtTranslator          ut.Translator

	// CustomRegexpPatternsValidator 自定义规则
	CustomRegexpPatternsValidator = []customRegexpRule{
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
)

func initValidator() {
	translator, ok := ut.New(TranslatorLocal, TranslatorLocal).GetTranslator(TranslatorLocal.Locale())
	if !ok {
		panic(fmt.Sprintf("lang %s locale is not exist", TranslatorLocal.Locale()))
		return
	}
	UtTranslator = translator

	if binding.Validator == nil {
		return
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册标签名
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			label := field.Tag.Get(ValidatorTagNameValue)
			if label == "" {
				return field.Name
			}
			return label
		})

		//注册默认翻译器
		if err := zhTrans.RegisterDefaultTranslations(v, UtTranslator); err != nil {
			panic(err)
		}

		//注册自定义验证规则
		if len(CustomRegexpPatternsValidator) != 0 {
			for _, item := range CustomRegexpPatternsValidator {
				err := v.RegisterValidation(item.Name, func(fl validator.FieldLevel) bool {
					return regexp.MustCompile(item.Pattern).MatchString(fl.Field().String())
				}, true)
				if err != nil {
					panic(err)
				}
				//注册自定义验证规则的翻译器
				err = v.RegisterTranslation(item.Name, UtTranslator, func(ut ut.Translator) error {
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
	}
}
