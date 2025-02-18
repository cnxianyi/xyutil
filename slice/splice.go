package slice

// Splice 从指定位置删除指定数量的元素. 支持传入替换的值. 修改原数组
// Splice removes a specified number of elements from a given position in the slice. It also allows inserting replacement values. Modifies the original slice.
func Splice[T interface{}](sl *[]T, start int, deleteCount int, arg ...T) {
	*sl = append(append((*sl)[0:start], arg...), (*sl)[start+deleteCount:]...)
}

// ToSplice 从指定位置删除指定数量的元素. 支持传入替换的值. 返回浅拷贝
// ToSplice removes a specified number of elements from a given position in the slice. It also allows inserting replacement values. returns a shallow copy of the slice
func ToSplice[T interface{}](sl []T, start int, deleteCount int, arg ...T) []T {
	var res = CopyWithIn(sl)
	res = append(append((res)[0:start], arg...), (res)[start+deleteCount:]...)
	return res
}
