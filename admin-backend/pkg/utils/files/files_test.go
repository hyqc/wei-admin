package files

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetContent(t *testing.T) {
	path := "test_1.txt"
	exc := "abcdefg"
	body, err := GetContent(path)
	assert.Nil(t, err, err, "读取文件错误")
	assert.Equal(t, []byte(exc), body, "内容读取不对")
}

func TestOverride(t *testing.T) {
	path := "test_2.txt"
	content := "a,b,c,d,e,f"
	err := Override(path, []byte(content))
	assert.Nil(t, err, err, "覆盖写错误")
	body, err := GetContent(path)
	assert.Nil(t, err, err, "覆盖写错误")
	assert.Equal(t, []byte(content), body)
}
