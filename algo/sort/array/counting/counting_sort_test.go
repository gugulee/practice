package counting

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountingSort(t *testing.T) {
	in := []int{2, 5, 3, 0, 2, 3, 0, 3}
	countingSort(in)
	require.Equal(t, []int{0, 0, 2, 2, 3, 3, 3, 5}, in)
}
