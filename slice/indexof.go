package slice

// 返回指定值在切片中第一次出现的下标. 没有则返回-1
// returns the index of the first occurrent of the specified value in the slice. If not found, returns -1
func IndexOf[T comparable](sl []T, val T) int {

	var res int = -1

	for i := 0; i < len(sl); i++ {
		if sl[i] == val {
			return i
		}
	}

	return res
}

// 返回指定值在切片中最后一次出现的下标. 没有则返回-1
// returns the index of the last occurrent of the specified value in the slice. If not found, returns -1
func LastIndexOf[T comparable](sl []T, val T) int {

	var res int = -1

	for i := len(sl) - 1; i >= 0; i-- {
		if sl[i] == val {
			return i
		}
	}

	return res
}
