package slice

import (
	"fmt"
	"reflect"
	"strconv"
)

// ToString 将切片转为字符串 支持自定义分隔符. 返回转换后的字符串
// ToString converts the slice to a string with a custom delimiter and returns the resulting string.
func ToString[T any](sl []T, separator ...string) string {
	var res string

	if len(separator) == 0 {
		Push(&separator, ",")
	}

	for _, v := range sl {
		switch val := any(v).(type) {
		case int:
			res += strconv.Itoa(val)
			res += separator[0]
		case string:
			res += val
			res += separator[0]
		case float64:
			res += strconv.FormatFloat(val, 'f', -1, 64)
			res += separator[0]
		case bool:
			res += strconv.FormatBool(val)
			res += separator[0]
		default:
			// 自定义类型/结构体
			if reflect.TypeOf(v).Kind() == reflect.Struct {
				res += fmt.Sprintf("%+v", v)
				res += separator[0]
			} else {
				res += fmt.Sprintf("%v", v)
				res += separator[0]
			}
		}
	}

	return res[0 : len(res)-1]
}
