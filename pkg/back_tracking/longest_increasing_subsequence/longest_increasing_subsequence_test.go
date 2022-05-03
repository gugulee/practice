package lis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isIncreasing(t *testing.T) {
	a := []int{2, 9, 3, 6}
	out := isIncreasing(a)
	require.False(t, out)

	a = []int{2, 3, 6}
	out = isIncreasing(a)
	require.True(t, out)

	a = []int{10, 2, 3, 6}
	out = isIncreasing(a)
	require.False(t, out)

	a = []int{2, 3, 6, 1}
	out = isIncreasing(a)
	require.False(t, out)
}

func Test_longestIncreasingSubsequence(t *testing.T) {
	// a := []int{2, 9, 3, 6}
	a := []int{2, 9, 3, 6, 5, 1, 7}
	longestIncreasingSubsequence(a, []int{}, 0)
	require.Equal(t, 4, maxLis)
	require.Equal(t, []int{2, 3, 5, 7}, maxResult)
}

func Test_longestIncreasingSubsequence1(t *testing.T) {
	// a := []int{2, 9, 3, 6}
	a := []int{2, 9, 3, 6, 5, 1, 7}
	longestIncreasingSubsequence1(a, []int{}, 0)
	require.Equal(t, 4, maxLis)
	require.Equal(t, []int{2, 3, 5, 7}, maxResult)
}

func Test_longestIncreasingSubsequence2(t *testing.T) {
	// a := []int{2, 9, 3, 6}
	a := []int{2, 9, 3, 6, 5, 1, 7}
	longestIncreasingSubsequence2(a, []int{}, -1, 0)
	require.Equal(t, 4, maxLis)
	require.Equal(t, []int{2, 3, 5, 7}, maxResult)
}

func Test_longestIncreasingSubsequence3(t *testing.T) {
	a := []int{2, 9, 3, 6}
	// a := []int{2, 9, 3, 6, 5, 1, 7}
	longestIncreasingSubsequence3(a, -1, 0, 0)
	require.Equal(t, 3, maxLis)
}
