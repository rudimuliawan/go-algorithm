package data_structures

import (
	"testing"
)

func TestQueueEnqueue(t *testing.T) {
	t.Run("enqueue single item", func(t *testing.T) {
		q := &Queue[int]{}
		q.Enqueue(1)

		if q.Len() != 1 {
			t.Errorf("expected size 1, got %d", q.Len())
		}
	})

	t.Run("enqueue multiple items increments size", func(t *testing.T) {
		q := &Queue[int]{}
		for _, v := range []int{1, 2, 3} {
			q.Enqueue(v)
		}

		if q.Len() != 3 {
			t.Errorf("expected size 3, got %d", q.Len())
		}
	})
}

func TestQueueDequeue(t *testing.T) {
	t.Run("dequeue from empty queue returns zero value", func(t *testing.T) {
		q := &Queue[int]{}
		got := q.Dequeue()

		if got != 0 {
			t.Errorf("expected zero value 0, got %d", got)
		}
	})

	t.Run("dequeue returns items in FIFO order", func(t *testing.T) {
		q := &Queue[int]{}
		for _, v := range []int{1, 2, 3} {
			q.Enqueue(v)
		}

		expected := []int{1, 2, 3}
		for i, exp := range expected {
			item := q.Dequeue()
			if item != exp {
				t.Errorf("at step %d: expected %d, got %d", i, exp, item)
			}
		}
	})

	t.Run("dequeue decrements size", func(t *testing.T) {
		q := &Queue[int]{}
		for _, v := range []int{1, 2, 3} {
			q.Enqueue(v)
		}

		for expected := 2; expected >= 0; expected-- {
			q.Dequeue()
			if q.Len() != expected {
				t.Errorf("expected size %d, got %d", expected, q.Len())
			}
		}
	})

	t.Run("dequeue all items then returns zero value", func(t *testing.T) {
		q := &Queue[string]{}
		for _, v := range []string{"a", "b", "c"} {
			q.Enqueue(v)
		}

		for range 3 {
			q.Dequeue()
		}

		got := q.Dequeue()
		if got != "" {
			t.Errorf("expected zero value \"\" after draining queue, got %q", got)
		}
	})

	t.Run("queue can be reused after being drained", func(t *testing.T) {
		q := &Queue[int]{}
		q.Enqueue(1)
		q.Enqueue(2)

		q.Dequeue()
		q.Dequeue()

		if !q.IsEmpty() {
			t.Fatalf("expected queue to be empty after draining")
		}

		q.Enqueue(3)
		q.Enqueue(4)

		expected := []int{3, 4}
		for i, exp := range expected {
			item := q.Dequeue()
			if item != exp {
				t.Errorf("at step %d: expected %d, got %d", i, exp, item)
			}
		}
	})
}

func TestQueueFirst(t *testing.T) {
	t.Run("first on empty queue returns zero value", func(t *testing.T) {
		q := &Queue[int]{}
		got := q.First()

		if got != 0 {
			t.Errorf("expected zero value 0, got %d", got)
		}
	})

	t.Run("first returns front item without removing it", func(t *testing.T) {
		q := &Queue[int]{}
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		item := q.First()
		if item != 1 {
			t.Errorf("expected 1, got %d", item)
		}
		if q.Len() != 3 {
			t.Errorf("expected size to remain 3, got %d", q.Len())
		}
	})

	t.Run("first returns same value on repeated calls", func(t *testing.T) {
		q := &Queue[string]{}
		q.Enqueue("a")
		q.Enqueue("b")

		first := q.First()
		second := q.First()

		if first != second {
			t.Errorf("expected repeated First() calls to match: %v != %v", first, second)
		}
	})

	t.Run("first reflects new front item after dequeue", func(t *testing.T) {
		q := &Queue[int]{}
		q.Enqueue(1)
		q.Enqueue(2)
		q.Dequeue()

		item := q.First()
		if item != 2 {
			t.Errorf("expected 2, got %d", item)
		}
	})
}

func TestQueueIsEmpty(t *testing.T) {
	t.Run("new queue is empty", func(t *testing.T) {
		q := &Queue[int]{}
		if !q.IsEmpty() {
			t.Error("expected new queue to be empty")
		}
	})

	t.Run("queue with items is not empty", func(t *testing.T) {
		q := &Queue[int]{}
		q.Enqueue(1)

		if q.IsEmpty() {
			t.Error("expected queue with items to not be empty")
		}
	})

	t.Run("queue is empty again after draining", func(t *testing.T) {
		q := &Queue[int]{}
		q.Enqueue(1)
		q.Dequeue()

		if !q.IsEmpty() {
			t.Error("expected queue to be empty after draining")
		}
	})
}

func TestQueueLen(t *testing.T) {
	t.Run("empty queue has size zero", func(t *testing.T) {
		q := &Queue[int]{}
		if q.Len() != 0 {
			t.Errorf("expected size 0, got %d", q.Len())
		}
	})

	t.Run("size reflects enqueue and dequeue operations", func(t *testing.T) {
		q := &Queue[int]{}
		q.Enqueue(1)
		q.Enqueue(2)

		if q.Len() != 2 {
			t.Errorf("expected size 2, got %d", q.Len())
		}

		q.Dequeue()

		if q.Len() != 1 {
			t.Errorf("expected size 1, got %d", q.Len())
		}
	})
}
