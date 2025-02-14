package slice

// Every 返回 切片中所有参数是否都能通过测试
// Every returns whether all element in the slice pass the test
func Every[T interface{}, F func(T) bool](sl []T, fc F) bool {
	var res = true

	for i := 0; i < len(sl); i++ {
		if fc(sl[i]) == false {
			res = false
		}
	}

	return res
}
