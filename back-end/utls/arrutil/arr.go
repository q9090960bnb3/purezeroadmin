package arrutil

func Contains[T comparable](arr []T, v T) bool {
	for _, item := range arr {
		if item == v {
			return true
		}
	}
	return false
}

// UniqueConcat 是一个泛型函数，用于拼接两个数组并去除重复值
func UniqueConcat[T comparable](a, b []T) []T {
	seen := make(map[T]bool)
	result := []T{}

	// 添加第一个数组的元素到结果中
	for _, item := range a {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	// 添加第二个数组的元素到结果中
	for _, item := range b {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}
