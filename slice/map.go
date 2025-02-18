package slice

// Map 返回一个新切片 其中每个元素都由原切片中对应位置的元素调用一次指定方法的返回值组成
// Map returns a new slice where each element is the result of applying the given function to the corresponding element in the original slice
func Map[T any](sl []T, fc func(T) T) []T {
	var res []T = make([]T, len(sl))

	for i := 0; i < len(sl); i++ {
		res[i] = fc(sl[i])
	}

	return res
}
