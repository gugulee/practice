package insertion

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertionSort(t *testing.T) {
	in := []int{4, 5, 6, 1, 3, 2}
	insertionSort(in)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6}, in)
}
