package weight

import (
	"fmt"
	"math/rand"
	"time"
)

type Number interface {
	~int | ~int32 | ~int64 | ~uint | ~uint32 | ~uint64
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
	total64 := int64(total)
	rate64 := 1 + rand.Int63n(total64)
	rate := T(rate64) // [1,total]
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

	pool := make([]Elem[T], len(weights))
	copy(pool, weights)

	for _, w := range pool {
		total += w.Weight
	}

	if total <= 0 {
		return nil, fmt.Errorf("权重配置错误:%v", pool)
	}

	indexes := make([]int, 0, n)

	if len(pool) == n {
		for i := 0; i < n; i++ {
			indexes = append(indexes, i)
		}
		return indexes, nil
	}

	//抽样主循环
	remaining := len(pool) //当前pool有效元素个数 [0, remaining-1]
	for len(indexes) < n {
		//安全生成[1,total]的随机数，避免int溢出
		total64 := int64(total)
		rate64 := T(1 + rand.Int63n(total64)) // [1,total64]
		rate := T(rate64)
		var acc T = 0
		//线性扫描命中项
		foundIndex := -1
		for i := 0; i < remaining; i++ {
			acc += pool[i].Weight
			if rate <= acc {
				foundIndex = i
				break
			}
		}
		if foundIndex == -1 {
			return nil, fmt.Errorf("随机失败")
		}
		indexes = append(indexes, pool[foundIndex].Index)
		remaining--
		pool[foundIndex], pool[remaining] = pool[remaining], pool[foundIndex]
		total -= pool[remaining].Weight
	}
	return indexes, nil
}
