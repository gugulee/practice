package selection

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelectionSort(t *testing.T) {
	in := []int{6, 4, 1, 5, 2, 7, 3}
	selectionSort(in)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, in)
}

func Test_selectionSortReview(t *testing.T) {
	in := []int{6, 4, 1, 5, 2, 7, 3}
	selectionSortReview(in)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, in)
}
