package weight

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func BenchmarkDrop(b *testing.B) {
	weights := []int{5, 10, 2, 3}
	values := Make(weights)
	for i := 0; i < b.N; i++ {
		_, err := Drop(values)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkDropN(b *testing.B) {
	weights := []int{5, 10, 2, 3}
	values := Make(weights)
	for i := 0; i < b.N; i++ {
		seed := 1 + rand.Intn(len(values))
		_, err := DropN(values, seed)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func TestMake(t *testing.T) {
	weights := []int{5, 10, 2, 3}
	values := Make(weights)
	fmt.Println(values)
}

func TestDrop(t *testing.T) {
	weights := []int{5, 10, 2, 3}
	values := Make(weights)
	for i := 0; i < 100; i++ {
		index, err := Drop(values)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(index)
	}
}

func TestDropN(t *testing.T) {
	weights := []int{5, 10, 2, 3}
	values := Make(weights)
	_, err := DropN(values, 5)
	assert.NotNil(t, err, "判断错误", err)
	_, err = DropN([]Elem[int]{}, 5)
	assert.NotNil(t, err, "参数错误", err)

	for i := 0; i < 100; i++ {
		seed := 1 + rand.Intn(len(weights))
		indexes, err := DropN(values, seed)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(fmt.Sprintf("values: %+v, rest: %+v", values, indexes))
	}
}

func TestArraySplice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr[0:0])
	fmt.Println(arr[10:])
	arr = append(arr[0:2], arr[2+1:]...)
	fmt.Println(arr)
}
