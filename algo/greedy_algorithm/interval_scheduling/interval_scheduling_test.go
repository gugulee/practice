package intervalscheduling

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sortDyadicSlice(t *testing.T) {
	in := [][]int{{3, 6}, {1, 7}, {2, 4}, {3, 7}}
	sortDyadicSlice(in)
	require.Equal(t, [][]int{{2, 4}, {3, 6}, {1, 7}, {3, 7}}, in)
}

func Test_intervalScheduling(t *testing.T) {
	in := [][]int{{3, 6}, {1, 3}, {2, 4}, {7, 14}}
	out := intervalScheduling(in)
	require.Equal(t, [][]int{{1, 3}, {3, 6}, {7, 14}}, out)

	in = [][]int{{3, 6}, {1, 3}, {2, 4}, {3, 5}, {5, 14}}
	out = intervalScheduling(in)
	require.Equal(t, [][]int{{1, 3}, {3, 5}, {5, 14}}, out)

	in = [][]int{{4, 6}, {3, 6}, {1, 3}, {2, 4}, {7, 14}}
	out = intervalScheduling(in)
	require.Equal(t, [][]int{{1, 3}, {4, 6}, {7, 14}}, out)
}
