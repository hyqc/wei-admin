package array

import "slices"

type Number interface {
	int | int8 | int32 | int64 | uint | uint32 | uint64 | float32 | float64
}

// Deduplicate 去重数组
// arr 要去重的数组
// removeZero 是否去除0，如果为true，则去移除0
// order 是否排序
func Deduplicate[T Number](arr []T, removeZero, order bool) []T {
	m := make(map[T]struct{}, len(arr))
	for _, v := range arr {
		m[v] = struct{}{}
	}
	res := make([]T, 0, len(m))
	for k := range m {
		if removeZero && k == 0 {
			continue
		}
		res = append(res, k)
	}
	if order {
		slices.Sort(res)
	}
	return res
}
