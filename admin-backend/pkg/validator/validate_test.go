package validator

import (
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"net/url"
	"testing"
)

type user struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Web      string `json:"web"`
	Age      int    `json:"age,omitempty"`
}

// 自定义验证器
func userValidatorFunc(data interface{}) url.Values {
	rules := govalidator.MapData{
		//"username": []string{"required", "between:3,5"},
		//"email":    []string{"required", "min:4", "max:20", "email"},
		//"web":      []string{"url"},
		"age": []string{"required"},
	}
	opts := govalidator.Options{
		Data:  data,
		Rules: rules,
	}
	return govalidator.New(opts).ValidateStruct()
}

func TestValidate(t *testing.T) {
	data := &user{
		Username: "123",
		Email:    "dfsdf",
		Web:      "dfsdf",
		Age:      1,
	}
	res := userValidatorFunc(data)
	errs := map[string][]string(res)
	// 3. 判断验证是否通过
	if len(errs) > 0 {
		for _, v := range errs {
			if len(v) == 0 {
				continue
			}
			fmt.Println(v[0])
		}
	}
}
