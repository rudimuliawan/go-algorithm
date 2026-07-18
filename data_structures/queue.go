package data_structures

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

func (q *Queue[T]) Dequeue() T {
	var zero T
	if q.IsEmpty() {
		return zero
	}

	item := q.first.Item
	q.first = q.first.Next
	if q.first == nil {
		q.last = nil
	}

	q.size--

	return item
}

func (q *Queue[T]) First() T {
	var zero T
	if q.IsEmpty() {
		return zero
	}

	return q.first.Item
}

func (q *Queue[T]) Len() int {
	return q.size
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}
