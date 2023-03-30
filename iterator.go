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

func (s *SliceIterator[T]) UpdateIndex() {
	if s.reverse {
		s.index--
	} else {
		s.index++
	}
}

func (s *SliceIterator[T]) Next() bool {
	if s.HasNext() {
		s.item = s.Elements[s.index]
		s.UpdateIndex()
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
