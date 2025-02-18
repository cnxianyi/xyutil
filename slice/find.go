package slice

import "fmt"

// Find 返回切片中第一个满足测试的值. 如果没有则返回 error undefined
// Find returns the first element in the slice that passes the test. if none is found, returns an error "undefined"
func Find[T any](sl []T, fc func(T) bool) (T, error) {

	var res T
	var ok bool = false

	for i := 0; i < len(sl); i++ {
		if fc(sl[i]) {
			res = sl[i]
			ok = true
			break
		}
	}

	if ok {
		return res, nil
	} else {
		return res, fmt.Errorf("undefined")
	}

}

// FindIndex 返回切片中满足测试的第一个元素的索引. 没有则返回 -1
// FindIndex returns the index of the first element in the slice that passes the test. if none is found, returns -1
func FindIndex[T any](sl []T, fc func(T) bool) int {
	var res int = -1

	for i := 0; i < len(sl); i++ {
		if fc(sl[i]) {
			res = i
			break
		}
	}

	return res
}

// FindLast 返回切片中倒数第一个满足测试的值. 如果没有则返回 error undefined
// FindLast returns the last element in the slice that passes the test. if none is found, returns an error "undefined"
func FindLast[T any](sl []T, fc func(T) bool) (T, error) {

	var res T
	var ok bool = false

	for i := len(sl) - 1; i >= 0; i-- {
		if fc(sl[i]) {
			res = sl[i]
			ok = true
			break
		}
	}

	if ok {
		return res, nil
	} else {
		return res, fmt.Errorf("undefined")
	}

}

// FindLastIndex 返回切片中满足测试的倒数第一个元素的索引. 没有则返回 -1
// FindLastIndex returns the index of the last element in the slice that passes the test. if none is found, returns -1
func FindLastIndex[T any](sl []T, fc func(T) bool) int {
	var res int = -1

	for i := len(sl) - 1; i >= 0; i-- {
		if fc(sl[i]) {
			res = i
			break
		}
	}

	return res
}
