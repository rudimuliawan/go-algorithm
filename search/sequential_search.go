package search

import (
	"errors"
	"fmt"
)

type node[T comparable, U comparable] struct {
	key   T
	value U
	next  *node[T, U]
}

type SequentialSearchST[T comparable, U comparable] struct {
	first *node[T, U]
}

func (s *SequentialSearchST[T, U]) Get(key T) (U, bool) {
	for curr := s.first; curr != nil; curr = curr.next {
		if curr.key == key {
			return curr.value, true
		}
	}

	var empty U
	return empty, false
}

func (s *SequentialSearchST[T, U]) Put(key T, value U) {
	for curr := s.first; curr != nil; curr = curr.next {
		if curr.key == key {
			curr.value = value
			return
		}
	}

	newFirst := node[T, U]{key: key, value: value}
	newFirst.next = s.first
	s.first = &newFirst
}

func (s *SequentialSearchST[T, U]) Delete(key T) error {
	prev := s.first

	for curr := s.first; curr != nil; curr = curr.next {
		if curr.key == key {
			if curr == s.first {
				s.first = s.first.next
				return nil
			}

			prev.next = curr.next
			return nil
		}

		prev = curr
	}

	return errors.New(fmt.Sprintf("Item with key %T not fount", key))
}
