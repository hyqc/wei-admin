package govalidate

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// TranslateError 翻译错误信息为注册的语言
func TranslateError(err error) error {
	if err == nil {
		return nil
	}
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}
	em := errs.Translate(trans)
	for _, e := range em {
		return errors.New(e)
	}
	return err
}

// ValidateWithCtx 执行验证器
func ValidateWithCtx(ctx *gin.Context, data any, call ...ValidatorFunc) error {
	if err := ctx.ShouldBind(data); err != nil {
		// 请求解析失败
		return TranslateError(err)
	}
	return Validate(data, call...)
}

// Validate 执行验证器
func Validate(data any, call ...ValidatorFunc) error {
	for _, handler := range call {
		if err := handler(data); err != nil {
			return err
		}
	}
	return nil
}
