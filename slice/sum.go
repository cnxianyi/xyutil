package slice

import "xyutil"

// Sum 计算传入切片所有 Number 的总和.
// Sum calculates the sum of all elements in the slice.
func Sum[T xyutil.RNumber](sl []T) T {
	var res T
	for _, v := range sl {
		res += v
	}

	return res
}
