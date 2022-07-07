package kmp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kmp1(t *testing.T) {
	in := "ababaeababacdb"
	p := "ababacd"
	out := kmp1(in, p)
	require.True(t, out)
}

func Test_kmp(t *testing.T) {
	in := "ababaeababacdb"
	p := "ababacd"
	out := kmp(in, p)
	require.True(t, out)
}

func Test_getNext1(t *testing.T) {
	in := "ababacd"
	next := getNext1(in)
	require.Equal(t, []int{-1, -1, 0, 1, 2, -1}, next)

	in = "ababacdd"
	next = getNext1(in)
	require.Equal(t, []int{-1, -1, 0, 1, 2, -1, -1}, next)

	in = "bcaebcaed"
	next = getNext1(in)
	require.Equal(t, []int{-1, -1, -1, -1, 0, 1, 2, 3}, next)
}

func Test_getNext(t *testing.T) {
	in := "ababacd"
	next := getNext(in)
	require.Equal(t, []int{-1, -1, 0, 1, 2, -1, -1}, next)

	in = "ababacdd"
	next = getNext(in)
	require.Equal(t, []int{-1, -1, 0, 1, 2, -1, -1, -1}, next)

	in = "bcaebcaedbc"
	next = getNext(in)
	require.Equal(t, []int{-1, -1, -1, -1, 0, 1, 2, 3, -1, 0, 1}, next)
}
