package queue

import "github.com/nini-k/go-tampletes/structures/list"

type Queue[T any] struct {
	list list.LinkedList[T]
}

func New[T any]() Queue[T] {
	return Queue[T]{list: list.New[T]()}
}

func (q *Queue[T]) Push(value T) {
	q.list.AddAtTail(value)
}

func (q *Queue[T]) Pop() T {
	val, ok := q.list.Head()
	if ok {
		q.list.RemoveAtHead()
	}
	return val
}

func (q *Queue[T]) Pick() T {
	val, _ := q.list.Head()
	return val
}

func (q *Queue[T]) Size() int {
	return q.list.Size()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Queue[T]) ConvertToSlice() []T {
	return q.list.ConvertToSlice()
}
