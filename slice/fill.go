package slice

import "fmt"

// Fill 使用value填充slice中start到end位置的元素
// Fill uses the value replace elements in the slice from start to end
func Fill[T interface{}](sl []T, value T, start int, end int) error {

	if start > end {
		return fmt.Errorf("start cannot be genater than end")
	}

	if end > len(sl) {
		return fmt.Errorf("end greater than sl length")
	}

	if start < 0 || end < 0 {
		return fmt.Errorf("error end or start")
	}

	for i := start; i <= end; i++ {
		sl[i] = value
	}

	return nil
}
