package slice

// CopyWithIn 返回 浅拷贝数组.
// CopyWithIn returns a shallow copy of elements in the slice.
func CopyWithIn[T any](sl []T) []T {
	var res []T

	for _, v := range sl {
		res = append(res, v)
	}

	return res
}
