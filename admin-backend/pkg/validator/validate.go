package validator

import (
	"admin/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
)

// Func 验证函数类型
type Func func(interface{}) url.Values

func Validate(c *gin.Context, params interface{}, handlers ...Func) (err error) {
	// 1. 解析请求，支持 JSON 数据、表单请求和 URL Query
	if err := c.ShouldBindJSON(params); err != nil {
		// 请求解析失败
		return err
	}
	utils.PrintfLn("request params: %v", params)
	if len(handlers) == 0 {
		return nil
	}
	// 2. 表单验证
	for _, handler := range handlers {
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
	}
	return nil
}

func GetValidateJsonOmitemptyTag(field string) string {
	return fmt.Sprintf("%s,omitempty", field)
}

func GetValidateJsonTag(field string) string {
	return fmt.Sprintf("%s", field)
}
