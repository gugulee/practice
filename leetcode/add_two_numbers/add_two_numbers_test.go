package atn

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func prepareData(vals ...int) *ListNode {
	var l *ListNode
	for _, v := range vals {
		if nil == l {
			l = NewNode(v)
		} else {
			InsertTail(l, v)
		}
	}

	return l
}

func Test_prepareData(t *testing.T) {
	/* prepare test data*/
	l1 := prepareData(2, 4, 3)
	l2 := prepareData(5, 6, 4)

	require.Equal(t, []int{2, 4, 3}, List(l1))
	require.Equal(t, []int{5, 6, 4}, List(l2))
}

func Test_addTwoNumbers(t *testing.T) {
	l1 := prepareData(2, 4, 3)
	l2 := prepareData(5, 6, 4)
	out := addTwoNumbers(l1, l2)
	require.NotNil(t, out)
	require.Equal(t, []int{7, 0, 8}, List(out))

	l1 = prepareData(1, 2)
	l2 = prepareData(1, 9, 1)
	out = addTwoNumbers(l1, l2)
	require.NotNil(t, out)
	require.Equal(t, []int{2, 1, 2}, List(out))

	l1 = prepareData(6, 9)
	l2 = prepareData(5, 0, 0, 1)
	out = addTwoNumbers(l1, l2)
	require.NotNil(t, out)
	require.Equal(t, []int{1, 0, 1, 1}, List(out))

	l1 = prepareData(1, 2)
	l2 = prepareData(1, 9, 9)
	out = addTwoNumbers(l1, l2)
	require.NotNil(t, out)
	require.Equal(t, []int{2, 1, 0, 1}, List(out))

	l1 = prepareData(9, 9, 9, 9, 9, 9, 9)
	l2 = prepareData(9, 9, 9, 9)
	out = addTwoNumbers(l1, l2)
	require.NotNil(t, out)
	require.Equal(t, []int{8, 9, 9, 9, 0, 0, 0, 1}, List(out))
}
