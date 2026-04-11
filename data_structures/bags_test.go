package data_structures

import (
	"slices"
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

	t.Run("add multiple items contains all items", func(t *testing.T) {
		bag := &Bag[string]{}
		bag.Add("first")
		bag.Add("second")
		bag.Add("third")

		items := collectAll(bag)
		expected := []string{"first", "second", "third"}

		if len(items) != len(expected) {
			t.Fatalf("expected %d items, got %d", len(expected), len(items))
		}
		for _, e := range expected {
			if !slices.Contains(items, e) {
				t.Errorf("expected item %v to be in bag", e)
			}
		}
	})

	t.Run("add integer items contains all items", func(t *testing.T) {
		bag := &Bag[int]{}
		bag.Add(1)
		bag.Add(2)
		bag.Add(3)

		items := collectAll(bag)
		expected := []int{1, 2, 3}

		if len(items) != len(expected) {
			t.Fatalf("expected %d items, got %d", len(expected), len(items))
		}
		for _, e := range expected {
			if !slices.Contains(items, e) {
				t.Errorf("expected item %d to be in bag", e)
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

	t.Run("iterate contains all inserted items", func(t *testing.T) {
		bag := &Bag[int]{}
		for i := range 5 {
			bag.Add(i)
		}

		items := collectAll(bag)
		if len(items) != 5 {
			t.Fatalf("expected 5 items, got %d", len(items))
		}
		for i := range 5 {
			if !slices.Contains(items, i) {
				t.Errorf("expected item %d to be in bag", i)
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
