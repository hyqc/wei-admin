package validate

import (
	"admin/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{}) url.Values

func Validate(c *gin.Context, params interface{}, handler ValidatorFunc) (err error) {
	// 1. 解析请求，支持 JSON 数据、表单请求和 URL Query
	if err := c.ShouldBind(params); err != nil {
		// 请求解析失败
		return err
	}
	utils.PrintfLn("request params: %+v", params)
	// 2. 表单验证
	res := handler(params)
	errs := map[string][]string(res)
	// 3. 判断验证是否通过
	if len(errs) > 0 {
		for _, v := range errs {
			if len(v) == 0 {
				continue
			}
			return errors.New(v[0])
		}
	}
	return
}

func GetValidateJSONTag(field string) string {
	return fmt.Sprintf("%s,omitempty", field)
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}
	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
