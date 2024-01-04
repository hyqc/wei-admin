package utils

import (
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

func TestArray2Set(t *testing.T) {
	arr := []int{1, 1, 3, 4, 5, 5, 6}
	res := []int{1, 3, 4, 5, 6}
	assert.Equal(t, res, Array2Set(arr), "去重失败")
}
