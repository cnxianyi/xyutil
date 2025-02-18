package slice

import (
	"xyutil"
)

// SortByMath 根据数字大小 插入排序
// SortByMath Sorts the slice in ascending order based on numeric values using insertion sort.
func SortByMath[T xyutil.RNumber](sl *[]T) {

	for i := range *sl {
		preIndex := i - 1   // 上一个值索引
		current := (*sl)[i] // 当前值

		// 循环检查 每个前值是否大于当前值
		for preIndex >= 0 && (*sl)[preIndex] > current {
			(*sl)[preIndex+1] = (*sl)[preIndex] // 将前值后移一位
			preIndex -= 1                       // 前值索引 -= 1
		}

		(*sl)[preIndex+1] = current // 当前值没有大于当前值后. 将最后一个前值设置为当前值
	}

}

// 根据自定义规则插入排序数组. 修改原数组. 自定义规则需要传入 前值和当前值,返回bool
// Sorts the slice based on custom rules, modifying the original slice. The custom rule function takes two arguments (previous value and current value) and returns a boolean.
func Sort[T interface{}](sl *[]T, fc func(a T, b T) bool) {
	for i := range *sl {
		preIndex := i - 1
		current := (*sl)[i]

		for preIndex >= 0 && fc((*sl)[preIndex], current) {
			(*sl)[preIndex+1] = (*sl)[preIndex]
			preIndex -= 1
		}

		(*sl)[preIndex+1] = current
	}
}

// ToSortByMath 根据数字大小 插入排序. 返回浅拷贝
// ToSortByMath returns a shallow copy of the slice, sorted in ascending order based on numeric values using insertion sort.
func ToSortedByMath[T xyutil.RNumber](sl []T) []T {

	var res []T = CopyWithIn(sl)

	for i := range sl {
		preIndex := i - 1   // 上一个值索引
		current := (res)[i] // 当前值

		// 循环检查 每个前值是否大于当前值
		for preIndex >= 0 && (res)[preIndex] > current {
			(res)[preIndex+1] = (res)[preIndex] // 将前值后移一位
			preIndex -= 1                       // 前值索引 -= 1
		}

		(res)[preIndex+1] = current // 当前值没有大于当前值后. 将最后一个前值设置为当前值
	}

	return res
}

// ToSort 根据自定义规则插入排序数组. 返回浅拷贝. 自定义规则需要传入 前值和当前值,返回bool
// ToSort the slice based on custom rules, returns a shallow copy of the slice. The custom rule function takes two arguments (previous value and current value) and returns a boolean.
func ToSort[T interface{}](sl []T, fc func(a T, b T) bool) []T {
	var res []T = CopyWithIn(sl)

	for i := range res {
		preIndex := i - 1
		current := (res)[i]

		for preIndex >= 0 && fc((res)[preIndex], current) {
			(res)[preIndex+1] = (res)[preIndex]
			preIndex -= 1
		}

		(res)[preIndex+1] = current
	}

	return res
}
