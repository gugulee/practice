package negativesequencedegree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	in := []int{1, 5, 6, 2, 3, 4}
	out := merge(in, 0, 2, 5)
	require.Equal(t, 6, out)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6}, in)
}

func Test_nsd(t *testing.T) {
	in := []int{6, 1, 5, 4, 3, 2}
	out := nsd(in, 0, 5)
	require.Equal(t, 11, out)
}
