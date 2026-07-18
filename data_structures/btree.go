package data_structures

type treeNode[T any] struct {
	Item  T
	Left  *treeNode[T]
	Right *treeNode[T]
}

type BTree[T any] struct {
	root *treeNode[T]
	size int
}

func (t *BTree[T]) LevelOrder() []T {
	if t.IsEmpty() {
		return []T{}
	}

	result := make([]T, 0, t.size)

	queue := &Queue[*treeNode[T]]{}
	queue.Enqueue(t.root)

	for !queue.IsEmpty() {
		current := queue.Dequeue()
		result = append(result, current.Item)

		if current.Left != nil {
			queue.Enqueue(current.Left)
		}

		if current.Right != nil {
			queue.Enqueue(current.Right)
		}
	}

	return result
}

func (t *BTree[T]) PreOrder() []T {
	if t.IsEmpty() {
		return []T{}
	}

	result := make([]T, 0, t.size)
	stack := &Stack[*treeNode[T]]{}

	stack.Push(t.root)

	for !stack.IsEmpty() {
		curr := stack.Pop()
		result = append(result, curr.Item)

		if curr.Right != nil {
			stack.Push(curr.Right)
		}

		if curr.Left != nil {
			stack.Push(curr.Left)
		}
	}

	return result
}

func (t *BTree[T]) InOrder() []T {
	if t.IsEmpty() {
		return []T{}
	}

	result := make([]T, 0, t.size)
	stack := &Stack[*treeNode[T]]{}

	current := t.root

	for current != nil || !stack.IsEmpty() {
		if current != nil {
			stack.Push(current)
			current = current.Left
			continue
		}

		current = stack.Pop()
		result = append(result, current.Item)
		current = current.Right
	}

	return result
}

func (t *BTree[T]) Insert(item T) {
	if t.root == nil {
		t.root = &treeNode[T]{Item: item}
		t.size++
		return
	}

	queue := make([]*treeNode[T], 0, t.size)
	queue = append(queue, t.root)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Left == nil {
			current.Left = &treeNode[T]{Item: item}
			break
		} else {
			queue = append(queue, current.Left)
		}

		if current.Right == nil {
			current.Right = &treeNode[T]{Item: item}
			break
		} else {
			queue = append(queue, current.Right)
		}
	}

	t.size++
}

func (t *BTree[T]) IsEmpty() bool {
	return t.size == 0
}

func (t *BTree[T]) Size() int {
	return t.size
}
