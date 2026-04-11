package data_structures

import (
	"iter"
)

type node[T any] struct {
	Item T
	Next *node[T]
}

type Bag[T any] struct {
	first *node[T]
}

func (b *Bag[T]) Add(item T) {
	b.first = &node[T]{Item: item, Next: b.first}
}

func (b *Bag[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		current := b.first
		for current != nil {
			if !yield(current.Item) {
				return
			}

			current = current.Next
		}
	}
}
