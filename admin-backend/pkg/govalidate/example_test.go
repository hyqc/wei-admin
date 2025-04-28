package govalidate

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

func ExampleTranslateError() {

	type BagItem struct {
		ID  int `json:"id" validate:"required,min=10" label:"背包ID"`
		Num int `json:"num" validate:"required,min=10" label:"背包数量"`
	}

	type Foo struct {
		Name string   `json:"name" validate:"required,max=32" label:"用户名"`
		Age  int      `json:"age" validate:"required,min=18" label:"年龄"`
		Bag  *BagItem `json:"bag" validate:"required" label:"背包"`
	}

	err := Init(zh.New(), func(valid *validator.Validate, trans ut.Translator) error {
		return zhTrans.RegisterDefaultTranslations(valid, trans)
	})
	if err != nil {
		panic(err)
	}

	foo := Foo{
		Name: "John1111111111111111111111111111111111111111111111111111111111",
		Age:  20,
		Bag: &BagItem{
			ID:  12,
			Num: 13,
		},
	}

	err = TranslateError(Validator.Struct(foo))
	fmt.Println(err)

	// Output:
	// 用户名长度不能超过32个字符
}

func ExampleValidate() {

	type BagItem struct {
		ID  int `json:"id"`
		Num int `json:"num"`
	}

	type Foo struct {
		Name string   `json:"name" `
		Age  int      `json:"age" `
		Bag  *BagItem `json:"bag" `
	}

	err := Init(zh.New(), func(valid *validator.Validate, trans ut.Translator) error {
		return zhTrans.RegisterDefaultTranslations(valid, trans)
	})
	if err != nil {
		panic(err)
	}

	foo := Foo{
		Name: "John1111111111111111111111111111111111111111111111111111111111",
		Age:  20,
		Bag: &BagItem{
			ID:  12,
			Num: 13,
		},
	}

	reqAddr := func(data any) error {
		rules := []*StructRule{
			{
				Type: Foo{},
				Rules: map[string]string{
					"ID":    "required",
					"Title": "required",
				},
			},
		}

		return ValidateStructWithRules(data, rules)
	}

	err = Validate(foo, reqAddr)
	fmt.Println(err)

	// Output:
	// <nil>
}
