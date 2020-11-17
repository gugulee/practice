package kmp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kmp(t *testing.T) {
	in := "ababaeababacdb"
	p := "ababacd"
	out := kmp(in, p)
	require.True(t, out)
}

func Test_getNext(t *testing.T) {
	in := "ababacd"
	next := getNext(in)
	require.Equal(t, []int{-1, -1, 0, 1, 2, -1}, next)

	in = "ababacdd"
	next = getNext(in)
	require.Equal(t, []int{-1, -1, 0, 1, 2, -1, -1}, next)
}
