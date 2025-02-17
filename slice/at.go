package slice

import (
	"xyutil"
)

// At 返回切片中指定索引位置的值.
// At Return the element at the index in the slice.
func At[T xyutil.RNumber](sl []T, at int) T {
	if at > 0 {
		return sl[at]
	} else if at < 0 {
		return sl[len(sl)+at]
	} else {
		return sl[0]
	}
}
