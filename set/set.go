package set

// Set接口
type Set[T comparable] interface {
	Add(key T)
	Delete(key T)
	Has(key T) bool
	Keys() []T
	Clear()
}

// MapSet结构体
type MapSet[T comparable] struct {
	value map[T]struct{}
}

// MapSet构造函数
func NewMapSet[T comparable]() *MapSet[T] {
	return &MapSet[T]{
		value: map[T]struct{}{},
	}
}

// Add 新增 key
// Add add key
func (s *MapSet[T]) Add(key T) {
	s.value[key] = struct{}{}
}

// Delete 删除 key
// Delete delete key
func (s *MapSet[T]) Delete(key T) {
	delete(s.value, key)
}

// Has 如果存在则返回true
// Has if key is exist. return true
func (s *MapSet[T]) Has(key T) bool {
	if _, ok := s.value[key]; ok {
		return true
	}
	return false
}

// Keys 返回 所有key组成的 切片.
// Keys return slice by all keys
func (s *MapSet[T]) Keys() []T {
	var res = make([]T, 0, len(s.value))
	for v := range s.value {
		res = append(res, v)
	}

	return res
}

// Clear 清空set
// Clear Clear all element
func (s *MapSet[T]) Clear() {
	s.value = make(map[T]struct{})
}
