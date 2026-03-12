package list

import "github.com/practice/pkg/list/single"

type Stack struct {
	data *single.LinkList

	// the top of the stack,
	// the top of the stack is the head of the list
	top *single.ListNode
}

func NewStack() *Stack {
	return &Stack{
		data: single.NewLinkList(),
		top:  nil,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Push(value interface{}) {
	s.data.InsertHead(value)
	// the top of the stack is the head of the list
	s.top = s.data.Head().Next()
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	v := s.top.Value()
	pNode := s.data.Head().PNext()
	*pNode = (*pNode).Next()
	s.top = s.data.Head().Next()

	return v
}

func (s *Stack) Print() []interface{} {
	if s.IsEmpty() {
		return []interface{}{}
	}

	var out []interface{}

	for node := s.top; nil != node; node = node.Next() {
		out = append(out, node.Value())
	}

	return out
}

func (s *Stack) Flush() {
	s.data = nil
	s.top = nil
}
