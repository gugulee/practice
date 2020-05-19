package list

import "github.com/practice/pkg/list/single"

type Stack struct {
	data *single.LinkList

	// the top of the stack
	top int
}

func NewStack() *Stack {
	return &Stack{
		data: single.NewLinkList(),
		top:  0,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack) Push(value interface{}) {
	s.top++
	s.data.InsertTail(value)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	v := s.data[s.top]
	s.top--
	return v
}

func (s *Stack) Print() []interface{} {
	if s.IsEmpty() {
		return []interface{}{}
	}

	var out []interface{}
	for i := s.top; i > 0; i-- {
		out = append(out, s.data[i])
	}

	return out
}

/*
func (s *Stack) Flush() {
	s.top = -1
}
*/
