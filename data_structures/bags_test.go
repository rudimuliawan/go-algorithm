package data_structures

import (
	"testing"
)

func TestBagAdd(t *testing.T) {
	t.Run("add single item", func(t *testing.T) {
		bag := &Bag[string]{}
		bag.Add("hello")

		items := collectAll(bag)
		if len(items) != 1 {
			t.Fatalf("expected 1 item, got %d", len(items))
		}

		if items[0] != "hello" {
			t.Errorf("expected 'hello', got %v", items[0])
		}
	})

	t.Run("add multiple items returns in LIFO order", func(t *testing.T) {
		bag := &Bag[string]{}
		bag.Add("first")
		bag.Add("second")
		bag.Add("third")

		items := collectAll(bag)
		expected := []string{"third", "second", "first"}

		if len(items) != len(expected) {
			t.Fatalf("expected %d items, got %d", len(expected), len(items))
		}
		for i, item := range items {
			if item != expected[i] {
				t.Errorf("at index %d: expected %v, got %v", i, expected[i], item)
			}
		}
	})

	t.Run("add integer items", func(t *testing.T) {
		bag := &Bag[int]{}
		bag.Add(1)
		bag.Add(2)
		bag.Add(3)

		items := collectAll(bag)
		expected := []int{3, 2, 1}

		if len(items) != len(expected) {
			t.Fatalf("expected %d items, got %d", len(expected), len(items))
		}
		for i, item := range items {
			if item != expected[i] {
				t.Errorf("at index %d: expected %d, got %d", i, expected[i], item)
			}
		}
	})
}

func TestBagIterator(t *testing.T) {
	t.Run("empty bag yields nothing", func(t *testing.T) {
		bag := &Bag[string]{}
		items := collectAll(bag)
		if len(items) != 0 {
			t.Errorf("expected 0 items, got %d", len(items))
		}
	})

	t.Run("iterates all items", func(t *testing.T) {
		bag := &Bag[int]{}
		bag.Add(1)
		bag.Add(2)
		bag.Add(3)

		items := collectAll(bag)
		if len(items) != 3 {
			t.Fatalf("expected 3 items, got %d", len(items))
		}
	})

	t.Run("early exit stops iteration", func(t *testing.T) {
		bag := &Bag[int]{}
		bag.Add(1)
		bag.Add(2)
		bag.Add(3)

		count := 0
		for range bag.All() {
			count++
			break
		}
		if count != 1 {
			t.Errorf("expected iterator to stop after break, got %d iterations", count)
		}
	})

	t.Run("multiple iterations are independent", func(t *testing.T) {
		bag := &Bag[string]{}
		bag.Add("a")
		bag.Add("b")
		bag.Add("c")

		first := collectAll(bag)
		second := collectAll(bag)

		if len(first) != len(second) {
			t.Fatalf("expected same length on repeated iteration: %d vs %d", len(first), len(second))
		}
		for i := range first {
			if first[i] != second[i] {
				t.Errorf("at index %d: first=%v, second=%v", i, first[i], second[i])
			}
		}
	})

	t.Run("iterate preserves insertion order (LIFO)", func(t *testing.T) {
		bag := &Bag[int]{}
		for i := range 5 {
			bag.Add(i)
		}

		items := collectAll(bag)
		for i, item := range items {
			expected := 4 - i
			if item != expected {
				t.Errorf("at index %d: expected %d, got %v", i, expected, item)
			}
		}
	})
}

func collectAll[T any](b *Bag[T]) []T {
	var result []T
	for item := range b.All() {
		result = append(result, item)
	}

	return result
}
