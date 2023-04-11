package utils

type Iterator[T any] interface {
	Next() bool
	Item() T
	Rev() Iterator[T]
}

type SliceIterator[T any] struct {
	Elements []T
	index    int
	item     T
	reverse  bool
}

func (s *SliceIterator[T]) HasNext() bool {
	return s.index < len(s.Elements) && s.index >= 0
}

func (s *SliceIterator[T]) updateIndex() {
	if s.reverse {
		s.index--
	} else {
		s.index++
	}
}

func (s *SliceIterator[T]) Next() bool {
	if s.HasNext() {
		s.item = s.Elements[s.index]
		s.updateIndex()
		return true
	}
	return false
}

func (s *SliceIterator[T]) Item() T {
	return s.item
}

func (s *SliceIterator[T]) Rev() Iterator[T] {
	if s.reverse {
		s.index = 0
	} else {
		s.index = len(s.Elements) - 1
	}
	s.reverse = !s.reverse
	return s
}

// IntoIterator Turn an array into an iterator
// Usage:
// ```
// var numbers []int = []int{1,2,3}
// iter := IntoIterator(numbers)
// ```
func IntoIterator[T any](elements []T) Iterator[T] {
	return &SliceIterator[T]{
		Elements: elements,
	}
}

// Collect Turn an iterator into a array
// Usage:
// ```
// var numbers []int = []int {1,2,3}
// iter := IntoIterator(numbers)
// arr := Collect(iter)
// ```
func Collect[T any](iter Iterator[T]) []T {
	var res []T

	for iter.Next() {
		res = append(res, iter.Item())
	}

	return res
}

// Fold the elements to a single one.
// Usage:
// ```
// var numbers []int = []int {1,2,3}
// iter := IntoIterator(numbers)
// ans := Fold(0, func(prev int, cur int) int { return prev + cur }, iter)
// ```
func Fold[T any, U any](initial U, reducer func(U, T) U, iter Iterator[T]) U {
	var initialValue = initial

	for iter.Next() {
		initialValue = reducer(initialValue, iter.Item())
	}

	return initialValue
}

// Reduce the elements to a single one.
// Usage:
// ```
// var numbers []int = []int {1,2,3}
// iter := IntoIterator(numbers)
// ans := Reduce(func(prev int, cur int) int { return prev + cur }, iter)
// ```
func Reduce[T any](reducer func(T, T) T, iter Iterator[T]) T {
	var res T

	if iter.Next() {
		res = iter.Item()

		for iter.Next() {
			res = reducer(res, iter.Item())
		}
	}

	return res
}

// FirstWhere find the first element which is match the predicate function
// Usage:
// ```
// var numbers []int = []int {1,2,3}
// iter := IntoIterator(numbers)
// res := FirstWhere(func (x int) bool { return x % 2 == 1 }, iter)
// ```
func FirstWhere[T any](pred func(T) bool, iter Iterator[T]) (T, error) {
	var res T

	for iter.Next() {
		if pred(iter.Item()) {
			return iter.Item(), nil
		}
	}

	return res, NotFound
}

// Contain check the element is exists in the list
// Usage:
// ```
// var numbers []int = []int {1,2,3}
// iter := IntoIterator(numbers)
// isExists := Contain(2, iter)
// ```
func Contain[T comparable](target T, iter Iterator[T]) bool {
	for iter.Next() {
		if target == iter.Item() {
			return true
		}
	}
	return false
}
