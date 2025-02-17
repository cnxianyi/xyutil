package slice

// Join 返回 切片中字符与间隔符组成的字符串
// Join returns a string created by concatenating the elements of the slice with a separator
func Join(sliceList []string, separator string) string {
	var res string
	var l = len(sliceList)

	for i := 0; i < l; i++ {
		res += sliceList[i]
		if i != l-1 {
			res += separator
		}
	}

	return res
}
