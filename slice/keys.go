package slice

// Keys 返回所有索引组成的切片
// Keys returns a new slice containing all the index in the slice.
func Keys[T any](sl []T) []int {
	var res []int = make([]int, len(sl))

	for i := 0; i < len(sl); i++ {
		res[i] = i
	}

	return res
}
