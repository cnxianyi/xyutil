package slice

// Join 返回 切片中字符与间隔符组成的字符串
// Join returns a string created by concatenating the elements of the slice with a separator
func Join(sl []string, sperator string) string {
	var res string
	var l = len(sl)

	for i := 0; i < l; i++ {
		res += sl[i]
		if i != l-1 {
			res += sperator
		}
	}

	return res
}
