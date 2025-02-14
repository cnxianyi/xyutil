package slice

// Concat 合并多个slice
// Concat concat more slice
func Concat[T interface{}](sls [][]T) []T {

	var res []T

	for i := 0; i < len(sls); i++ {
		for j := 0; j < len(sls[i]); j++ {
			res = append(res, sls[i][j])
		}
	}

	return res
}
