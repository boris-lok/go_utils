package utils

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

// Reduce the elements to a single one.
// Usage:
// ```
// var numbers []int = []int {1,2,3}
// iter := IntoIterator(numbers)
// ans := Reduce(0, func(prev int, cur int) int { return prev + cur }, iter)
// ```
func Reduce[T any, U any](initial U, reducer func(U, T) U, iter Iterator[T]) U {
	var initialValue = initial

	for iter.Next() {
		initialValue = reducer(initialValue, iter.Item())
	}

	return initialValue
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
