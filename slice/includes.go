package slice

// Includes 查找切片中是否存在指定值. 存在返回true
// Includes check if the specified value exists in the slice, return true else false
func Includes[T comparable](sl []T, val T) bool {
	for i := 0; i < len(sl); i++ {
		if sl[i] == val {
			return true
		}
	}
	return false
}
