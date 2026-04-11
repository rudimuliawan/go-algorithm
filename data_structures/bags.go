package data_structures

import (
	"iter"
)

type Bag[T any] struct {
	first *node[T]
	size  int
}

func (b *Bag[T]) Add(item T) {
	b.first = &node[T]{Item: item, Next: b.first}
	b.size++
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

func (b *Bag[T]) Size() int {
	return b.size
}

func (b *Bag[T]) IsEmpty() bool {
	return b.size == 0
}
