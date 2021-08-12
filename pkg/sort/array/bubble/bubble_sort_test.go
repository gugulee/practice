package bubble

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubbleSort(t *testing.T) {
	in := []int{6, 4, 1, 5, 2, 7, 3}
	bubbleSort(in)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, in)
}
