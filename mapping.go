package utils

type MappingIterator[T any, U any] struct {
	s      Iterator[T]
	mapper func(T) U
}

func (m *MappingIterator[T, U]) Next() bool {
	return m.s.Next()
}

func (m *MappingIterator[T, U]) Item() U {
	return m.mapper(m.s.Item())
}

func (m *MappingIterator[T, U]) Rev() Iterator[U] {
	m.s = m.s.Rev()
	return m
}

// Map transforms the iterator into another, and then apply the mapper function to each element
// Usage:
// ```
// var numbers []int = []int{1,2,3}
// iter := IntoIterator(numbers)
// mapped := Collect(Map(func (x int) int { return x * x }, iter))
// ```
func Map[T any, U any](mapper func(T) U, iter Iterator[T]) Iterator[U] {
	return &MappingIterator[T, U]{
		s:      iter,
		mapper: mapper,
	}
}
