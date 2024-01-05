package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray2Set(t *testing.T) {
	arr := []int{1, 1, 3, 4, 5, 5, 6}
	res := []int{1, 3, 4, 5, 6}
	assert.Equal(t, res, Array2Set(arr), "去重失败")
}
