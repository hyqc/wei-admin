package govalidate

import (
	"fmt"
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
	Validator  = validator.New(validator.WithRequiredStructEnabled())
	Translator ut.Translator
)

type custom struct {
	lang         locales.Translator
	translator   CustomTranslator
	tagNameValue string
}

// ValidatorFunc 验证函数类型，data 为结构体类型的数值
type ValidatorFunc func(data any) error

// Rules 验证规则
type Rules []*StructRule

// StructRule 验证规则
type StructRule struct {
	// Type  interface{}
	// 示例：
	// type Foo struct {
	// 	 Name string
	//   Age int
	// }
	// 则Type为Foo{}
	Type any

	// Rules map[string]string
	// key为结构体字段名，value为验证规则，示例：
	// rules := map[string]string{
	//		"Name": "required,min=6,max=32",
	//		"Age": "required,min=18,max=100",
	// }
	Rules map[string]string
}

type Option func(c *custom)

// CustomTranslator 自定义翻译器注册方法
type CustomTranslator func(validate *validator.Validate, trans ut.Translator) error

// WithTagNameValue 设置字段标签名的值，示例label
//
//	示例：
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

func (c *custom) initCustom() error {
	var ok bool
	// 返回一个多语言翻译器实例
	uni := ut.New(c.lang, c.lang)
	// 返回中文翻译器
	Translator, ok = uni.GetTranslator(c.lang.Locale())
	if !ok {
		return fmt.Errorf("lang locale is not exist")
	}
	if err := c.translator(Validator, Translator); err != nil {
		return err
	}
	Validator.RegisterTagNameFunc(func(field reflect.StructField) string {
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
	return newCustom(lang, translator, opts...).initCustom()
}
