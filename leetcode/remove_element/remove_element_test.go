package removeelement

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeElement(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		want int
	}{
		{
			[]int{3, 2, 2, 3},
			3,
			2,
		},
		{
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			5,
		},
	}

	for _, tt := range tests {
		re := removeElement(tt.nums, tt.val)
		require.Equal(t, re, tt.want)
	}
}
