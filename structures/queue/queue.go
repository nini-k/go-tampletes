package queue

import "github.com/nini-k/go-tampletes/structures/list"

type Queue struct {
	list list.LinkedList
}

func (q *Queue) Push(value interface{}) {
	q.list.AddAtTail(value)
}

func (q *Queue) Pop() interface{} {
	val, ok := q.list.Head()
	if ok {
		q.list.RemoveAtHead()
	}
	return val
}

func (q *Queue) Pick() interface{} {
	val, _ := q.list.Head()
	return val
}

func (q *Queue) Size() int {
	return q.list.Size()
}

func (q *Queue) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Queue) ConvertToSlice() []interface{} {
	return q.list.ConvertToSlice()
}

func New() Queue {
	return Queue{list: list.New()}
}
