package data_structures

type node[T any] struct {
	Item T
	Next *node[T]
}
