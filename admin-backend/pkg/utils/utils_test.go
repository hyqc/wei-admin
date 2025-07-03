package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBeanCopy(t *testing.T) {
	type a struct {
		Id   int
		Name string
		Val  map[string]string
	}
	type b struct {
		Id   int32
		Name string
		Val  map[string]string
	}
	av := a{
		Id:   1,
		Name: "aaaa",
		Val:  map[string]string{"a": "a", "b": "b"},
	}
	bv := b{}
	err := BeanCopy(&bv, &av)
	assert.Nil(t, err, err)
	assert.Equal(t, b{
		Id:   0,
		Name: "aaaa",
		Val:  map[string]string{"a": "a", "b": "b"},
	}, bv)
}

func TestGetConfigEnv(t *testing.T) {
	fmt.Println(GetConfigEnv("mode"))
}

func TestCamelToSnake(t *testing.T) {
	arr := []string{
		"abcAbc",
		"hideInMenu",
	}
	for _, v := range arr {
		fmt.Println(CamelToSnake(v))
	}
}

func TestGetOutBoundIP(t *testing.T) {
	ip, err := GetOutBoundIP()
	assert.Nil(t, err, "错误1")
	fmt.Println("ip: ", ip)

	lip, err := GetLocalBoundIP()
	assert.Nil(t, err, "错误1")
	fmt.Println("lip: ", lip)
}
