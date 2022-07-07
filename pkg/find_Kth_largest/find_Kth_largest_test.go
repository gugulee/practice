package findkthlargest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindKthLargest(t *testing.T) {
	in := []int{8, 10, 2, 3, 6, 1, 5}
	k := 5
	out := findKthLargest(in, 0, len(in)-1, k)
	require.Equal(t, 6, out)
}
