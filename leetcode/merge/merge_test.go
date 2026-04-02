package merge

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
			[]int{1, 2, 2, 3, 5, 6},
		},
		{
			[]int{1},
			1,
			make([]int, 1),
			0,
			[]int{1},
		},
		{
			make([]int, 1),
			0,
			[]int{1},
			1,
			[]int{1},
		},
	}

	for _, tt := range tests {
		merge(tt.nums1, tt.m, tt.nums2, tt.n)
		require.ElementsMatch(t, tt.nums1, tt.want)
	}
}
