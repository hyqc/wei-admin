package weight

import (
	"fmt"
	"math/rand"
	"time"
)

type Number interface {
	~int8 | ~int | ~int32 | ~int64 | ~uint8 | ~uint | ~uint32 | ~uint64
}

type Elem[T Number] struct {
	Index  int
	Weight T
}

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Make 权重数组生成权重随机参数
func Make[T Number](weights []T) []Elem[T] {
	values := make([]Elem[T], 0, len(weights))
	for i, val := range weights {
		values = append(values, Elem[T]{
			Index:  i,
			Weight: val,
		})
	}
	return values
}

// Drop 根据权重掉落，返回索引
func Drop[T Number](weights []Elem[T]) (int, error) {
	if len(weights) == 0 {
		return -1, fmt.Errorf("权重参数为空")
	}
	var total T = 0
	for _, w := range weights {
		total += w.Weight
	}

	if total <= 0 {
		return -1, fmt.Errorf("权重配置错误:%v", weights)
	}

	rate := T(1 + rand.Intn(int(total))) // [1,total]
	var acc T
	for _, w := range weights {
		acc += w.Weight
		if rate <= acc {
			return w.Index, nil
		}
	}
	return -1, fmt.Errorf("随机失败")
}

// DropN 根据权重掉落，返回索引列表
func DropN[T Number](weights []Elem[T], n int) ([]int, error) {
	if len(weights) == 0 {
		return nil, fmt.Errorf("权重参数为空")
	}
	if n <= 0 || n > len(weights) {
		return nil, fmt.Errorf("权重随机数量无效")
	}
	var total T = 0
	for _, w := range weights {
		total += w.Weight
	}

	if total <= 0 {
		return nil, fmt.Errorf("权重配置错误:%v", weights)
	}

	indexes := make([]int, 0, n)

	if len(weights) == n {
		for i := 0; i < n; i++ {
			indexes = append(indexes, i)
		}
		return indexes, nil
	}

loop:
	{
		rate := T(1 + rand.Intn(int(total))) // [1,total]
		var acc T = 0
		for i, w := range weights {
			acc += w.Weight
			if rate <= acc {
				indexes = append(indexes, w.Index)
				if len(indexes) == n {
					return indexes, nil
				}
				weights = append(weights[0:i], weights[i+1:]...)
				total -= w.Weight
				goto loop
			}
		}
	}
	return nil, fmt.Errorf("随机失败")
}
