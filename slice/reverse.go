package slice

// Reverse 反转切片. 会修改原切片
// Reverse reverses the slice in place, modifying the original slice.
func Reverse[T interface{}](sl *[]T) *[]T {
	var cache []T

	for i := len(*sl) - 1; i >= 0; i-- {
		cache = append(cache, (*sl)[i])
	}

	*sl = cache
	return sl
}
