package slice

// Some 返回 切片中是否存在一个值满足测试
// Some checks if at least one element in the slice passes the test and returns true if found, otherwise false.
func Some[T interface{}](sl []T, fc func(T) bool) bool {
	for i := 0; i < len(sl); i++ {
		if fc(sl[i]) {
			return true
		}
	}

	return false
}
