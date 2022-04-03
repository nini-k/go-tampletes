package stack

import "github.com/nini-k/go-tampletes/structures/list"

type Stack[T any] struct {
	list list.LinkedList[T]
}

func New[T any]() Stack[T] {
	return Stack[T]{list: list.New[T]()}
}

func (s *Stack[T]) Push(value T) {
	s.list.AddAtHead(value)
}

func (s *Stack[T]) Pop() T {
	val, ok := s.list.Head()
	if ok {
		s.list.RemoveAtHead()
	}
	return val
}

func (s *Stack[T]) Peek() T {
	val, _ := s.list.Head()
	return val
}

func (s *Stack[T]) Size() int {
	return s.list.Size()
}

func (s *Stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *Stack[T]) ConvertToSlice() []T {
	return s.list.ConvertToSlice()
}
