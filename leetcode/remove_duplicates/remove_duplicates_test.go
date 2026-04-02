package removeduplicates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			[]int{1, 1, 2},
			2,
		},
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
		},
	}

	for _, tt := range tests {
		re := removeDuplicates(tt.nums)
		require.Equal(t, re, tt.want)
	}
}
