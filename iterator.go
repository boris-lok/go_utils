package iter

type Iterator[T any] interface {
	Next() bool
	Item() T
}

type SliceIterator[T any] struct {
	Elements []T
	index    int
	item     T
}

func (s *SliceIterator[T]) Next() bool {
	if s.index < len(s.Elements) {
		s.item = s.Elements[s.index]
		s.index++
		return true
	}
	return false
}

func (s *SliceIterator[T]) Item() T {
	return s.item
}

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

type FilterIterator[T any] struct {
	s    Iterator[T]
	pred func(T) bool
}

func (f *FilterIterator[T]) Next() bool {
	for f.s.Next() {
		if f.pred(f.s.Item()) {
			return true
		}
	}
	return false
}

func (f *FilterIterator[T]) Item() T {
	return f.s.Item()
}

func IntoIterator[T any](elements []T) Iterator[T] {
	return &SliceIterator[T]{
		Elements: elements,
	}
}

func Map[T any, U any](mapper func(T) U, iter Iterator[T]) Iterator[U] {
	return &MappingIterator[T, U]{
		s:      iter,
		mapper: mapper,
	}
}

func Filter[T any](pred func(T) bool, iter Iterator[T]) Iterator[T] {
	return &FilterIterator[T]{
		s:    iter,
		pred: pred,
	}
}

func Collect[T any](iter Iterator[T]) []T {
	var res []T

	for iter.Next() {
		res = append(res, iter.Item())
	}

	return res
}

func Reduce[T any, U any](initial U, reducer func(U, T) U, iter Iterator[T]) U {
	var initialValue = initial

	for iter.Next() {
		initialValue = reducer(initialValue, iter.Item())
	}

	return initialValue
}
