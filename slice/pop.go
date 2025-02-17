package slice

// Pop 弹出slice尾部元素,并返回. 修改原数组
// Pop removes and returns the last element of the slice, modifying the original slice.
func Pop[T interface{}](sl *[]T) T {
	var res T
	l := len(*sl)
	res = (*sl)[l-1]
	*sl = (*sl)[:l-1]
	return res
}

// Push 在切片末尾添加元素. 返回更新后的切片长度
// Push add element in the last at slice. returns new slice's length
func Push[T interface{}](sl *[]T, val ...T) int {
	var cache []T = val
	*sl = append(*sl, cache...)
	return len(*sl)
}

// Shift 删除切片中第一个元素. 修改原切片. 返回被删除的元素
// Shift removes and returns the first element of the slice, modifying the original slice.
func Shift[T interface{}](sl *[]T) T {
	var res T = (*sl)[0]
	*sl = (*sl)[1:]
	return res
}

// Unshift 在切片头部添加内容. 返回更新后的切片长度
// Unshift adds elements to the beginning of the slice and returns the updated length of the slice.
func Unshift[T interface{}](sl *[]T, val ...T) int {
	var cache []T
	cache = append(cache, val...)
	cache = append(cache, *sl...)

	*sl = cache

	return len(*sl)
}
