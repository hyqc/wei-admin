package govalidate

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
)

const (
	// DefaultTagNameValue 结构体的字段名的标签值，示例：label
	//	type Foo struct {
	//	    Age  int    `json:"age" validate:"required,min=18" label:"年龄"`
	//	    Name string `json:"name" validate:"required,max=32" label:"用户名"`
	//	}
	DefaultTagNameValue = "label"
)

var (
	// Validator 验证器实例
	Validator = validator.New(validator.WithRequiredStructEnabled())
	trans     ut.Translator
)

type custom struct {
	lang         locales.Translator
	translator   CustomTranslator
	tagNameValue string
}

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{}) error

type Rules struct {
	Items []*RulesItem
}

type RulesItem struct {
	Type  interface{} //结构体
	Rules map[string]string
}

type Option func(c *custom)

// CustomTranslator 自定义翻译器注册方法
type CustomTranslator func(validate *validator.Validate, trans ut.Translator) error

// WithTagNameValue 设置字段标签名的值，示例label
//
//	type Foo struct {
//	    Age  int    `json:"age" validate:"required,min=18" label:"年龄"`
//	    Name string `json:"name" validate:"required,max=32" label:"用户名"`
//	}
func WithTagNameValue(val string) Option {
	return func(c *custom) {
		c.tagNameValue = val
	}
}

func newCustom(lang locales.Translator, translator CustomTranslator, opts ...Option) *custom {
	c := &custom{
		lang:         lang,
		translator:   translator,
		tagNameValue: DefaultTagNameValue,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *custom) initCustom(v *validator.Validate) error {
	if v == nil {
		v = Validator
	}
	var ok bool
	// 返回一个多语言翻译器实例
	uni := ut.New(c.lang, c.lang)
	// 返回中文翻译器
	trans, ok = uni.GetTranslator(c.lang.Locale())
	if !ok {
		return fmt.Errorf("lang locale is not exist")
	}
	if err := c.translator(v, trans); err != nil {
		return err
	}
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get(c.tagNameValue)
		if label == "" {
			return field.Name
		}
		return label
	})
	return nil
}

// Init 验证器初始化配置
func Init(lang locales.Translator, translator CustomTranslator, opts ...Option) error {
	return newCustom(lang, translator, opts...).initCustom(nil)
}

func GinInitTrans(lang locales.Translator, translator CustomTranslator, opts ...Option) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		return newCustom(lang, translator, opts...).initCustom(v)
	}
	return nil
}

func GetTrans() ut.Translator {
	return trans
}
