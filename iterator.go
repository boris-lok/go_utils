package iter

type Iterator[T any] interface {
	Next() bool
	Item() T
	Filter(pred func(T) bool) Iterator[T]
	Collect() []T
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

func (s *SliceIterator[T]) Filter(pred func(T) bool) Iterator[T] {
	return &FilterIterator[T]{
		source:    s,
		predicate: pred,
	}
}

func (s *SliceIterator[T]) Collect() []T {
	var res []T

	for s.Next() {
		res = append(res, s.Item())
	}

	return res
}

type FilterIterator[T any] struct {
	source    Iterator[T]
	predicate func(T) bool
}

func (f *FilterIterator[T]) Next() bool {
	for f.source.Next() {
		if f.predicate(f.source.Item()) {
			return true
		}
	}
	return false
}

func (f *FilterIterator[T]) Item() T {
	return f.source.Item()
}

func (f *FilterIterator[T]) Filter(pred func(T) bool) Iterator[T] {
	return &FilterIterator[T]{
		source:    f,
		predicate: pred,
	}
}

func (f *FilterIterator[T]) Collect() []T {
	var res []T

	for f.Next() {
		res = append(res, f.Item())
	}

	return res
}

func IntoIterator[T any](elements []T) Iterator[T] {
	return &SliceIterator[T]{
		Elements: elements,
	}
}
