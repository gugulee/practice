package array

type Stack struct {
	data []interface{}

	// the top of the stack
	top int
}

func New() *Stack {
	return &Stack{
		data: make([]interface{}, 10),
		top:  -1,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.top < 0
}

func (s *Stack) Push(value interface{}) {
	s.top++

	if s.top > len(s.data)-1 {
		s.data = append(s.data, value)
	} else {
		s.data[s.top] = value
	}
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
	for i := s.top; 0 <= i; i-- {
		out = append(out, s.data[i])
	}

	return out
}

func (s *Stack) Flush() {
	s.top = -1
}

func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.top]
}
