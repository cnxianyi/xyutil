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
func Push[T interface{}](sl *[]T, val T) int {
	*sl = append(*sl, val)
	return len(*sl)
}
