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
