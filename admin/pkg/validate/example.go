package validate

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

type exampleData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Web      string `json:"web"`
}

// 自定义验证器
func exampleValidatorFunc(data interface{}) url.Values {
	rules := govalidator.MapData{
		"username": []string{"required", "between:3,5"},
		"email":    []string{"required", "min:4", "max:20", "email"},
		"web":      []string{"url"},
	}
	opts := govalidator.Options{
		Data:  data,
		Rules: rules,
	}
	return govalidator.New(opts).ValidateStruct()
}

func example(ctx *gin.Context) {
	req := &exampleData{}
	if err := Validate(ctx, req, exampleValidatorFunc); err != nil {
		fmt.Println(err)
		return
	}
}
