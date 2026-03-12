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

	require.Equal(t, "6->5->4->3->2->1->0", s.data.Print())
	require.Equal(t, s.data.SearchListByValue(6), s.top)
}

func TestPrint(t *testing.T) {
	s := NewStack()
	for _, v := range []int{0, 1, 2, 3, 4, 5, 6} {
		s.Push(v)
	}

	out := s.Print()

	require.Equal(t, []interface{}{6, 5, 4, 3, 2, 1, 0}, out)

	s = NewStack()
	out = s.Print()
	require.Equal(t, []interface{}{}, out)
}

func TestPop(t *testing.T) {
	s := NewStack()
	for _, v := range []int{0, 1, 2, 3, 4, 5, 6} {
		s.Push(v)
	}

	out := s.Pop()
	require.Equal(t, interface{}(6), out)
	require.Equal(t, "5->4->3->2->1->0", s.data.Print())
	require.Equal(t, s.data.SearchListByValue(5), s.top)

	out = s.Pop()
	require.Equal(t, interface{}(5), out)
	require.Equal(t, "4->3->2->1->0", s.data.Print())
	require.Equal(t, s.data.SearchListByValue(4), s.top)

	out = s.Pop()
	require.Equal(t, interface{}(4), out)
	require.Equal(t, "3->2->1->0", s.data.Print())
	require.Equal(t, s.data.SearchListByValue(3), s.top)

	out = s.Pop()
	require.Equal(t, interface{}(3), out)
	require.Equal(t, "2->1->0", s.data.Print())
	require.Equal(t, s.data.SearchListByValue(2), s.top)

	out = s.Pop()
	require.Equal(t, interface{}(2), out)
	require.Equal(t, "1->0", s.data.Print())
	require.Equal(t, s.data.SearchListByValue(1), s.top)

	out = s.Pop()
	require.Equal(t, interface{}(1), out)
	require.Equal(t, "0", s.data.Print())
	require.Equal(t, s.data.SearchListByValue(0), s.top)

	out = s.Pop()
	require.Equal(t, interface{}(0), out)
	require.Equal(t, nil, s.data.Print())
	require.Equal(t, nil, s.top)

	out = s.Pop()
	require.Equal(t, nil, out)
	require.Equal(t, nil, s.data.Print())
	require.Equal(t, nil, s.top)

	out = s.Pop()
	require.Equal(t, nil, out)
	require.Equal(t, nil, s.data.Print())
	require.Equal(t, nil, s.top)
}
