package array

import (
	"fmt"
	"testing"
)

func TestDistinct(t *testing.T) {
	a := []int32{1, 0, 2, 3, 4, 5}
	fmt.Println(Deduplicate(a, true, false))
}
