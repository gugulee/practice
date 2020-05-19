package list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	s := NewStack()
	out := s.IsEmpty()
	require.Equal(t, true, out)
}

func TestPush(t *testing.T) {
	s := NewStack()
	for _, v := range []int{0, 1, 2, 3, 4, 5, 6} {
		s.Push(v)
	}

	require.Equal(t, "0->1->2->3->4->5->6", s.data.Print())
	require.Equal(t, 7, s.top)
}


func TestPrint(t *testing.T) {
	s := NewStack(5)
	for _, v := range []int{0, 1, 2, 3, 4, 5, 6} {
		s.Push(v)
	}

	out := s.Print()

	require.Equal(t, []interface{}{6, 5, 4, 3, 2, 1, 0}, out)

	s = NewStack(5)
	out = s.Print()
	require.Equal(t, []interface{}{}, out)
}
/*
func TestPop(t *testing.T) {
	s := NewStack(5)
	for _, v := range []int{0, 1, 2, 3, 4, 5, 6} {
		s.Push(v)
	}

	out := s.Pop()
	require.Equal(t, interface{}(6), out)
	require.Equal(t, []interface{}{0, 1, 2, 3, 4, 5}, s.data[0:s.top+1])

	out = s.Pop()
	require.Equal(t, interface{}(5), out)
	require.Equal(t, []interface{}{0, 1, 2, 3, 4}, s.data[0:s.top+1])

	out = s.Pop()
	require.Equal(t, interface{}(4), out)
	require.Equal(t, []interface{}{0, 1, 2, 3}, s.data[0:s.top+1])

	out = s.Pop()
	require.Equal(t, interface{}(3), out)
	require.Equal(t, []interface{}{0, 1, 2}, s.data[0:s.top+1])

	out = s.Pop()
	require.Equal(t, interface{}(2), out)
	require.Equal(t, []interface{}{0, 1}, s.data[0:s.top+1])

	out = s.Pop()
	require.Equal(t, interface{}(1), out)
	require.Equal(t, []interface{}{0}, s.data[0:s.top+1])

	out = s.Pop()
	require.Equal(t, interface{}(0), out)
	require.Equal(t, []interface{}{}, s.data[0:s.top+1])

	out = s.Pop()
	require.Equal(t, nil, out)
	require.Equal(t, -1, s.top)

	out = s.Pop()
	require.Equal(t, nil, out)
	require.Equal(t, -1, s.top)
}
*/
