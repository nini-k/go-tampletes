package list

type linkedListNode[T any] struct {
	value T
	next  *linkedListNode[T]
}

type LinkedList[T any] struct {
	tail *linkedListNode[T]
	head *linkedListNode[T]
	size int
}

func New[T any]() LinkedList[T] {
	return LinkedList[T]{}
}

func (l *LinkedList[T]) AddAtTail(value T) {
	tmp := &linkedListNode[T]{value: value}
	if l.tail == nil {
		l.tail = tmp
		l.head = l.tail
	} else {
		l.tail.next = tmp
		l.tail = l.tail.next
	}
	l.size++
}

func (l *LinkedList[T]) AddAtHead(value T) {
	tmp := &linkedListNode[T]{value: value}
	if l.head == nil {
		l.head = tmp
		l.tail = l.head
	} else {
		tmp.next = l.head
		l.head = tmp
	}
	l.size++
}

func (l *LinkedList[T]) Head() (out T, ok bool) {
	if l.head == nil {
		return
	}
	return l.head.value, true
}

func (l *LinkedList[T]) Tail() (out T, ok bool) {
	if l.tail == nil {
		return
	}
	return l.tail.value, true
}

func (l *LinkedList[T]) RemoveAtHead() {
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

func (l *LinkedList[T]) ForEach(fn func(value T)) {
	cur := l.head
	for cur != nil {
		fn(cur.value)
		cur = cur.next
	}
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.head == nil && l.tail == nil
}

func (l *LinkedList[T]) ConvertToSlice() []T {
	out, cur := make([]T, 0, l.size), l.head
	for cur != nil {
		out = append(out, cur.value)
		cur = cur.next
	}
	return out
}
