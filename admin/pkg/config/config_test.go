package config

import (
	json2 "encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResetJsonOrigin(t *testing.T) {
	type A struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Max  int    `json:"max"`
	}

	type B struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	a := &A{
		Id:   1,
		Name: "1",
		Max:  1,
	}

	b := &B{
		Id:   2,
		Name: "2",
	}

	aBody, _ := json2.Marshal(a)
	fmt.Println(string(aBody))
	bBody, _ := json2.Marshal(b)
	_ = json2.Unmarshal(bBody, a)
	assert.Equal(t, 1, a.Max)

	err := ResetJsonOrigin(bBody, a)
	fmt.Println(a)
	assert.Nil(t, err, "错误")
	assert.Equal(t, 0, a.Max)
}
