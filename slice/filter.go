package slice

// Filter 返回 满足切片中满足测试的元素的 浅拷贝
// Filter returns a shallow copy of elements in the slice that pass the test
func Filter[T any](sl []T, fc func(T) bool) []T {

	var res []T

	for i := 0; i < len(sl); i++ {
		if fc(sl[i]) {
			res = append(res, sl[i])
		}
	}

	return res
}
