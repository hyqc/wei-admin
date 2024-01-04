package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateAdminUserLastLoginIp(t *testing.T) {
	s, err := SetAdminUserLastLoginIp("", "")
	assert.Nil(t, err, err, "处理登录IP")
	assert.Equal(t, "", s, "返回错误")

	s, err = SetAdminUserLastLoginIp("127.0.0.2", "[\"127.0.0.1\"]")
	assert.Nil(t, err, err, "处理登录IP")
	assert.Equal(t, "[\"127.0.0.1\",\"127.0.0.2\"]", s, "返回错误")

	s, err = SetAdminUserLastLoginIp("127.0.0.3", "[\"127.0.0.1\",\"127.0.0.2\"]")
	assert.Nil(t, err, err, "处理登录IP")
	assert.Equal(t, "[\"127.0.0.2\",\"127.0.0.3\"]", s, "返回错误")
}
