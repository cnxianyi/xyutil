package xyutil

// RNumber 实数
type RNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

// Integer 整数
type Integer interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Number 所有数字类型 包括complex
type Number interface {
	RNumber | ~complex64 | ~complex128
}
