package data_structures

import "errors"

var ErrEmptyQueue = errors.New("queue is empty")

type Queue[T any] struct {
	first *node[T]
	last  *node[T]
	size  int
}

func (q *Queue[T]) Enqueue(item T) {
	if q.IsEmpty() {
		q.last = &node[T]{Item: item}
		q.first = q.last
	} else {
		newLast := &node[T]{Item: item}
		q.last.Next = newLast
		q.last = newLast
	}

	q.size++
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, ErrEmptyQueue
	}

	item := q.first.Item
	q.first = q.first.Next
	if q.first == nil {
		q.last = nil
	}

	q.size--

	return item, nil
}

func (q *Queue[T]) First() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, ErrEmptyQueue
	}

	return q.first.Item, nil
}

func (q *Queue[T]) Len() int {
	return q.size
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}
