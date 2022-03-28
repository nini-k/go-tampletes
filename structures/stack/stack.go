package stack

import "github.com/nini-k/go-tampletes/structures/list"

type Stack struct {
	list list.LinkedList
}

func (s *Stack) Push(value interface{}) {
	s.list.AddAtHead(value)
}

func (s *Stack) Pop() interface{} {
	val, ok := s.list.Head()
	if ok {
		s.list.RemoveAtHead()
	}
	return val
}

func (s *Stack) Peek() interface{} {
	val, _ := s.list.Head()
	return val
}

func (s *Stack) Size() int {
	return s.list.Size()
}

func (s *Stack) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *Stack) ConvertToSlice() []interface{} {
	return s.list.ConvertToSlice()
}

func New() Stack {
	return Stack{list: list.New()}
}
