package slice

// ForEach 为切片中每个元素执行一遍给定函数
// ForEach executes the givin function for each elementt in the slice.
func ForEach[T interface{}](sl []T, fc func(T)) {
	for i := 0; i < len(sl); i++ {
		fc(sl[i])
	}
}
