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

// Filter use predicate function to check each element in the iterator
// Usage:
// ```
// var numbers []int = []int {1,2,3}
// iter := IntoIterator(numbers)
// filtered := Collect(Filter(func (x int) bool { return x % 2 == 0 }, iter)
// ```
func Filter[T any](pred func(T) bool, iter Iterator[T]) Iterator[T] {
	return &FilterIterator[T]{
		s:    iter,
		pred: pred,
	}
}
