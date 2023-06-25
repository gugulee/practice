package quick

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuickSort(t *testing.T) {
	in := []int{8, 10, 2, 3, 6, 1, 5}
	quickSort(in, 0, len(in)-1)
	require.Equal(t, []int{1, 2, 3, 5, 6, 8, 10}, in)
}

func Test_quickSortReview(t *testing.T) {
	in := []int{8, 10, 2, 3, 6, 1, 5}
	quickSortReview(in, 0, len(in)-1)
	require.Equal(t, []int{1, 2, 3, 5, 6, 8, 10}, in)
}
