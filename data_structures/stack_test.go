package data_structures

import (
	"errors"
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
	t.Run("pop from empty stack returns error", func(t *testing.T) {
		s := &Stack[int]{}
		_, err := s.Pop()

		if !errors.Is(err, ErrEmptyStack) {
			t.Errorf("expected ErrEmptyStack, got %v", err)
		}
	})

	t.Run("pop returns items in LIFO order", func(t *testing.T) {
		s := &Stack[int]{}
		for _, v := range []int{1, 2, 3} {
			s.Push(v)
		}

		expected := []int{3, 2, 1}
		for i, exp := range expected {
			item, err := s.Pop()
			if err != nil {
				t.Fatalf("unexpected error at step %d: %v", i, err)
			}
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
			if _, err := s.Pop(); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if s.Len() != expected {
				t.Errorf("expected size %d, got %d", expected, s.Len())
			}
		}
	})

	t.Run("pop all items then returns error", func(t *testing.T) {
		s := &Stack[string]{}
		for _, v := range []string{"a", "b", "c"} {
			s.Push(v)
		}

		for range 3 {
			if _, err := s.Pop(); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		}

		_, err := s.Pop()
		if !errors.Is(err, ErrEmptyStack) {
			t.Errorf("expected ErrEmptyStack after draining stack, got %v", err)
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
