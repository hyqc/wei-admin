package validate

import (
	"admin/global"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidatorFunc func(data any) error

// TranslateError 翻译错误信息为注册的语言
func TranslateError(err error) error {
	if err == nil {
		return nil
	}
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}
	em := errs.Translate(global.UtTranslator)
	for _, e := range em {
		return errors.New(e)
	}
	return err
}

// WithCtx 执行验证器
func WithCtx(ctx *gin.Context, data any, call ...ValidatorFunc) error {
	if err := ctx.ShouldBind(data); err != nil {
		// 请求解析失败
		return TranslateError(err)
	}
	if call == nil || len(call) == 0 {
		return nil
	}
	return Validate(data, call...)
}

// Validate 执行验证器
func Validate(data any, call ...ValidatorFunc) error {
	for _, handler := range call {
		if err := handler(data); err != nil {
			return TranslateError(err)
		}
	}
	return nil
}
