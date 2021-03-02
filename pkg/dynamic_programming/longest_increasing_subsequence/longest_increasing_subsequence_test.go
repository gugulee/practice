package lis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestIncreasingSubsequence(t *testing.T) {
	// a := []int{2, 9, 3, 6}
	a := []int{4, 10, 4, 3, 8, 9}
	// a := []int{1, 2, 3, 4, 6, 7, 4}
	// a := []int{2, 9, 3, 6, 5, 1, 7}
	out := longestIncreasingSubsequence(a)
	require.Equal(t, 3, out)
}