package data_structures

import (
	"errors"
)

var ErrEmptyStack = errors.New("stack is empty")

type Stack[T any] struct {
	top  *node[T]
	size int
}

func (s *Stack[T]) Push(item T) {
	s.top = &node[T]{Item: item, Next: s.top}
	s.size++
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if s.top == nil {
		return zero, ErrEmptyStack
	}

	item := s.top.Item
	s.top = s.top.Next
	s.size--

	return item, nil
}

func (s *Stack[T]) Len() int {
	return s.size
}
