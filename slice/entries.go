package slice

// Entries 返回 切片转换为map的 可迭代对象
// Entries returns a map created from the slice
func Entries[T interface{}](sl []T) map[int]T {
	res := make(map[int]T)

	for i := 0; i < len(sl); i++ {
		res[i] = sl[i]
	}

	return res
}
