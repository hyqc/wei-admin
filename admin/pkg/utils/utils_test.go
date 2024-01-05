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
	err := BeanCopy(&av, &bv)
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
