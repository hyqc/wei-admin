package pwd

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncode(t *testing.T) {
	a := "123456"
	b, err := Encode(a)
	fmt.Println("123456的密码:", b)
	assert.Nil(t, err, err, "加密失败")
	assert.True(t, Matches(a, b), "密码不匹配")
}
