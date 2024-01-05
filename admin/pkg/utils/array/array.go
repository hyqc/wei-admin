package array

func Array2Set[T comparable](data []T) (result []T) {
	m := make(map[T]struct{})
	for _, val := range data {
		if _, ok := m[val]; !ok {
			m[val] = struct{}{}
			result = append(result, val)
		}
	}
	return result
}
