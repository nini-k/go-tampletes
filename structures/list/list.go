package list

type linkedListNode struct {
	value interface{}
	next  *linkedListNode
}

type LinkedList struct {
	tail *linkedListNode
	head *linkedListNode
	size int
}

func (l *LinkedList) AddAtTail(value interface{}) {
	tmp := &linkedListNode{value: value}
	if l.tail == nil {
		l.tail = tmp
		l.head = l.tail
	} else {
		l.tail.next = tmp
		l.tail = l.tail.next
	}
	l.size++
}

func (l *LinkedList) AddAtHead(value interface{}) {
	tmp := &linkedListNode{value: value}
	if l.head == nil {
		l.head = tmp
		l.tail = l.head
	} else {
		tmp.next = l.head
		l.head = tmp
	}
	l.size++
}

func (l *LinkedList) Head() (interface{}, bool) {
	if l.head == nil {
		return nil, false
	}
	return l.head.value, true
}

func (l *LinkedList) Tail() (interface{}, bool) {
	if l.tail == nil {
		return nil, false
	}
	return l.tail.value, true
}

func (l *LinkedList) RemoveAtHead() {
	if l.head == nil {
		return
	}
	head := l.head
	l.head = l.head.next
	head.next = nil
	if head == l.tail {
		l.tail = l.head
	}
	l.size--
}

func (l *LinkedList) ForEach(foo func(value interface{})) {
	cur := l.head
	for cur != nil {
		foo(cur.value)
		cur = cur.next
	}
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.head == nil && l.tail == nil
}

func (l *LinkedList) ConvertToSlice() []interface{} {
	out, cur := make([]interface{}, 0, l.size), l.head
	for cur != nil {
		out = append(out, cur.value)
		cur = cur.next
	}
	return out
}

func New() LinkedList {
	return LinkedList{}
}
