package utils

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

func (f *FilterIterator[T]) Rev() Iterator[T] {
	f.s = f.s.Rev()
	return f
}
