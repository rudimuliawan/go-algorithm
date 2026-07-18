package data_structures

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	t.Run("push single item", func(t *testing.T) {
		s := &Stack[int]{}
		s.Push(1)

		if s.Len() != 1 {
			t.Errorf("expected size 1, got %d", s.Len())
		}
	})

	t.Run("push multiple items increments size", func(t *testing.T) {
		s := &Stack[int]{}
		for _, v := range []int{1, 2, 3} {
			s.Push(v)
		}

		if s.Len() != 3 {
			t.Errorf("expected size 3, got %d", s.Len())
		}
	})
}

func TestStackPop(t *testing.T) {
	t.Run("pop from empty stack returns zero value", func(t *testing.T) {
		s := &Stack[int]{}
		got := s.Pop()

		if got != 0 {
			t.Errorf("expected zero value 0, got %d", got)
		}
	})

	t.Run("pop returns items in LIFO order", func(t *testing.T) {
		s := &Stack[int]{}
		for _, v := range []int{1, 2, 3} {
			s.Push(v)
		}

		expected := []int{3, 2, 1}
		for i, exp := range expected {
			item := s.Pop()
			if item != exp {
				t.Errorf("at step %d: expected %d, got %d", i, exp, item)
			}
		}
	})

	t.Run("pop decrements size", func(t *testing.T) {
		s := &Stack[int]{}
		for _, v := range []int{1, 2, 3} {
			s.Push(v)
		}

		for expected := 2; expected >= 0; expected-- {
			s.Pop()
			if s.Len() != expected {
				t.Errorf("expected size %d, got %d", expected, s.Len())
			}
		}
	})

	t.Run("pop all items then returns zero value", func(t *testing.T) {
		s := &Stack[string]{}
		for _, v := range []string{"a", "b", "c"} {
			s.Push(v)
		}

		for range 3 {
			s.Pop()
		}

		got := s.Pop()
		if got != "" {
			t.Errorf("expected zero value \"\" after draining stack, got %q", got)
		}
	})
}

func TestStackTop(t *testing.T) {
	t.Run("top on empty stack returns zero value", func(t *testing.T) {
		s := &Stack[int]{}
		got := s.Top()

		if got != 0 {
			t.Errorf("expected zero value 0, got %d", got)
		}
	})

	t.Run("top returns top item without removing it", func(t *testing.T) {
		s := &Stack[int]{}
		s.Push(1)
		s.Push(2)
		s.Push(3)

		item := s.Top()
		if item != 3 {
			t.Errorf("expected 3, got %d", item)
		}
		if s.Len() != 3 {
			t.Errorf("expected size to remain 3, got %d", s.Len())
		}
	})

	t.Run("top returns same value on repeated calls", func(t *testing.T) {
		s := &Stack[string]{}
		s.Push("a")
		s.Push("b")

		first := s.Top()
		second := s.Top()

		if first != second {
			t.Errorf("expected repeated Top() calls to match: %v != %v", first, second)
		}
	})

	t.Run("top reflects most recently pushed item after pop", func(t *testing.T) {
		s := &Stack[int]{}
		s.Push(1)
		s.Push(2)
		s.Pop()

		item := s.Top()
		if item != 1 {
			t.Errorf("expected 1, got %d", item)
		}
	})
}

func TestStackIsEmpty(t *testing.T) {
	t.Run("new stack is empty", func(t *testing.T) {
		s := &Stack[int]{}
		if !s.IsEmpty() {
			t.Error("expected new stack to be empty")
		}
	})

	t.Run("stack with items is not empty", func(t *testing.T) {
		s := &Stack[int]{}
		s.Push(1)

		if s.IsEmpty() {
			t.Error("expected stack with items to not be empty")
		}
	})

	t.Run("stack is empty again after draining", func(t *testing.T) {
		s := &Stack[int]{}
		s.Push(1)
		s.Pop()

		if !s.IsEmpty() {
			t.Error("expected stack to be empty after draining")
		}
	})
}

func TestStackLen(t *testing.T) {
	t.Run("empty stack has size zero", func(t *testing.T) {
		s := &Stack[int]{}
		if s.Len() != 0 {
			t.Errorf("expected size 0, got %d", s.Len())
		}
	})

	t.Run("size reflects push and pop operations", func(t *testing.T) {
		s := &Stack[int]{}
		s.Push(1)
		s.Push(2)

		if s.Len() != 2 {
			t.Errorf("expected size 2, got %d", s.Len())
		}

		s.Pop()

		if s.Len() != 1 {
			t.Errorf("expected size 1, got %d", s.Len())
		}
	})
}
