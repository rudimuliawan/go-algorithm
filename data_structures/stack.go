package data_structures

type Stack[T any] struct {
	top  *node[T]
	size int
}

func (s *Stack[T]) Push(item T) {
	s.top = &node[T]{Item: item, Next: s.top}
	s.size++
}

func (s *Stack[T]) Pop() T {
	var zero T
	if s.IsEmpty() {
		return zero
	}

	item := s.top.Item
	s.top = s.top.Next
	s.size--

	return item
}

func (s *Stack[T]) Top() T {
	var zero T
	if s.IsEmpty() {
		return zero
	}

	return s.top.Item
}

func (s *Stack[T]) Len() int {
	return s.size
}

func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}
