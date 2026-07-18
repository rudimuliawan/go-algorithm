package data_structures

import (
	"reflect"
	"testing"
)

func TestBinaryTreeInsert(t *testing.T) {
	t.Run("insert into empty tree sets root", func(t *testing.T) {
		tree := &BTree[int]{}
		tree.Insert(1)

		if tree.root == nil {
			t.Fatal("expected root to be set")
		}
		if tree.root.Item != 1 {
			t.Errorf("expected root item 1, got %d", tree.root.Item)
		}
		if tree.Size() != 1 {
			t.Errorf("expected size 1, got %d", tree.Size())
		}
	})

	t.Run("second insert becomes left child of root", func(t *testing.T) {
		tree := &BTree[int]{}
		tree.Insert(1)
		tree.Insert(2)

		if tree.root.Left == nil {
			t.Fatal("expected root.Left to be set")
		}
		if tree.root.Left.Item != 2 {
			t.Errorf("expected root.Left item 2, got %d", tree.root.Left.Item)
		}
		if tree.root.Right != nil {
			t.Errorf("expected root.Right to remain nil, got %v", tree.root.Right.Item)
		}
	})

	t.Run("third insert becomes right child of root", func(t *testing.T) {
		tree := &BTree[int]{}
		tree.Insert(1)
		tree.Insert(2)
		tree.Insert(3)

		if tree.root.Right == nil {
			t.Fatal("expected root.Right to be set")
		}
		if tree.root.Right.Item != 3 {
			t.Errorf("expected root.Right item 3, got %d", tree.root.Right.Item)
		}
	})

	t.Run("insert does not duplicate item into both children", func(t *testing.T) {
		tree := &BTree[int]{}
		tree.Insert(1)
		tree.Insert(2)

		if tree.root.Right != nil {
			t.Fatalf("expected root.Right to still be nil after a single Insert, got %v", tree.root.Right.Item)
		}
		if tree.Size() != 2 {
			t.Errorf("expected size 2 after two inserts, got %d", tree.Size())
		}
	})

	t.Run("insert maintains level-order placement across two levels", func(t *testing.T) {
		tree := &BTree[int]{}
		for _, v := range []int{1, 2, 3, 4, 5, 6, 7} {
			tree.Insert(v)
		}

		root := tree.root
		if root == nil || root.Item != 1 {
			t.Fatal("expected root item 1")
		}
		if root.Left == nil || root.Left.Item != 2 {
			t.Fatal("expected root.Left item 2")
		}
		if root.Right == nil || root.Right.Item != 3 {
			t.Fatal("expected root.Right item 3")
		}
		if root.Left.Left == nil || root.Left.Left.Item != 4 {
			t.Error("expected root.Left.Left item 4")
		}
		if root.Left.Right == nil || root.Left.Right.Item != 5 {
			t.Error("expected root.Left.Right item 5")
		}
		if root.Right.Left == nil || root.Right.Left.Item != 6 {
			t.Error("expected root.Right.Left item 6")
		}
		if root.Right.Right == nil || root.Right.Right.Item != 7 {
			t.Error("expected root.Right.Right item 7")
		}

		if tree.Size() != 7 {
			t.Errorf("expected size 7, got %d", tree.Size())
		}
	})
}

func TestBinaryTreeSize(t *testing.T) {
	t.Run("empty tree has size zero", func(t *testing.T) {
		tree := &BTree[int]{}
		if tree.Size() != 0 {
			t.Errorf("expected size 0, got %d", tree.Size())
		}
	})

	t.Run("size reflects number of inserts", func(t *testing.T) {
		tree := &BTree[int]{}
		for i, v := range []int{10, 20, 30, 40} {
			tree.Insert(v)

			if tree.Size() != i+1 {
				t.Errorf("after inserting %d: expected size %d, got %d", v, i+1, tree.Size())
			}
		}
	})
}

func TestBinaryTreeLevelOrder(t *testing.T) {
	t.Run("level-order on empty tree returns empty slice", func(t *testing.T) {
		tree := &BTree[int]{}
		got := tree.LevelOrder()

		if len(got) != 0 {
			t.Errorf("expected empty slice, got %v", got)
		}
	})

	t.Run("level-order on single node tree returns that item", func(t *testing.T) {
		tree := &BTree[int]{}
		tree.Insert(1)

		got := tree.LevelOrder()
		expected := []int{1}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("level-order returns items in level-order for a full two-level tree", func(t *testing.T) {
		tree := &BTree[int]{}
		for _, v := range []int{1, 2, 3, 4, 5, 6, 7} {
			tree.Insert(v)
		}

		got := tree.LevelOrder()
		expected := []int{1, 2, 3, 4, 5, 6, 7}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("level-order reflects a partially filled last level", func(t *testing.T) {
		tree := &BTree[int]{}
		for _, v := range []int{1, 2, 3, 4, 5} {
			tree.Insert(v)
		}

		got := tree.LevelOrder()
		expected := []int{1, 2, 3, 4, 5}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}

func TestBinaryTreeInOrder(t *testing.T) {
	t.Run("in-order on empty tree returns empty slice", func(t *testing.T) {
		tree := &BTree[int]{}
		got := tree.InOrder()

		if len(got) != 0 {
			t.Errorf("expected empty slice, got %v", got)
		}
	})

	t.Run("in-order on single node tree returns that item", func(t *testing.T) {
		tree := &BTree[int]{}
		tree.Insert(1)

		got := tree.InOrder()
		expected := []int{1}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("in-order returns items in in-order for a full two-level tree", func(t *testing.T) {
		tree := &BTree[int]{}
		for _, v := range []int{1, 2, 3, 4, 5, 6, 7} {
			tree.Insert(v)
		}

		got := tree.InOrder()
		expected := []int{4, 2, 5, 1, 6, 3, 7}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("in-order reflects a partially filled last level", func(t *testing.T) {
		tree := &BTree[int]{}
		for _, v := range []int{1, 2, 3, 4, 5} {
			tree.Insert(v)
		}

		got := tree.InOrder()
		expected := []int{4, 2, 5, 1, 3}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}

func TestBinaryTreePreOrder(t *testing.T) {
	t.Run("pre-order on empty tree returns empty slice", func(t *testing.T) {
		tree := &BTree[int]{}
		got := tree.PreOrder()

		if len(got) != 0 {
			t.Errorf("expected empty slice, got %v", got)
		}
	})

	t.Run("pre-order on single node tree returns that item", func(t *testing.T) {
		tree := &BTree[int]{}
		tree.Insert(1)

		got := tree.PreOrder()

		expected := []int{1}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("pre-order returns items in pre-order for a full two-level tree", func(t *testing.T) {
		tree := &BTree[int]{}
		for _, v := range []int{1, 2, 3, 4, 5, 6, 7} {
			tree.Insert(v)
		}

		got := tree.PreOrder()

		expected := []int{1, 2, 4, 5, 3, 6, 7}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("pre-order reflects a partially filled last level", func(t *testing.T) {
		tree := &BTree[int]{}
		for _, v := range []int{1, 2, 3, 4, 5} {
			tree.Insert(v)
		}

		got := tree.PreOrder()

		expected := []int{1, 2, 4, 5, 3}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}

func TestBinaryTreeIsEmpty(t *testing.T) {
	t.Run("new tree is empty", func(t *testing.T) {
		tree := &BTree[int]{}
		if !tree.IsEmpty() {
			t.Error("expected new tree to be empty")
		}
	})

	t.Run("tree with items is not empty", func(t *testing.T) {
		tree := &BTree[int]{}
		tree.Insert(1)

		if tree.IsEmpty() {
			t.Error("expected tree with items to not be empty")
		}
	})
}
