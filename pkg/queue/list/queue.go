package list

import "github.com/practice/pkg/list/single"

// Queue is queue of list
// the tail of the queue is the head of the list
type Queue struct {
	data *single.LinkList

	// the tail of the queue
	tail *single.ListNode
}

// New return a new queue
func New() *Queue {
	l := single.NewLinkList()
	return &Queue{
		data: l,
		tail: l.Head(),
	}
}

// IsEmpty return true if the queue is empty, else return false
func (q *Queue) IsEmpty() bool {
	return nil == q.data.Head().Next()
}

// Enqueue ...
func (q *Queue) Enqueue(value interface{}) {
	q.data.InsertAfter(q.tail, value)
	q.tail = q.tail.Next()
}

// Dequeue ...
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	realHead := q.data.Head().Next()
	value := realHead.Value()

	q.data.DeleteNode(realHead)

	return value
}

// Print print the queue
func (q *Queue) Print() []interface{} {
	if q.IsEmpty() {
		return nil
	}

	var out []interface{}
	realHead := q.data.Head().Next()

	for node := realHead; nil != node; node = node.Next() {
		out = append(out, node.Value())
	}

	return out
}
