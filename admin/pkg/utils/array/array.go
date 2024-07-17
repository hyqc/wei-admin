package array

import "sort"

type Number interface {
	int | int8 | int32 | int64 | uint | uint32 | uint64 | float32 | float64
}

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
		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})
	}
	return res
}
