package merge

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	in := []int{1, 2, 4, 6, 7}
	var in1 []int
	out := merge(in, in1)
	assert.Equal(t, in, out)

	in = []int{}
	in1 = []int{1, 2, 4, 6, 7}
	out = merge(in, in1)
	assert.Equal(t, in1, out)

	in = []int{1, 2, 4, 6, 7}
	in1 = []int{0, 3, 5, 8, 9}
	out = merge(in, in1)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, out)

	in = []int{0, 3, 5, 8, 9}
	in1 = []int{1, 2, 4, 6, 7}
	out = merge(in, in1)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, out)

	in = []int{0, 3, 5, 8, 9}
	in1 = []int{1, 2, 3, 6, 7}
	out = merge(in, in1)
	assert.Equal(t, []int{0, 1, 2, 3, 3, 5, 6, 7, 8, 9}, out)
}

func TestMerge1(t *testing.T) {
	in := []int{2, 4, 6, 1}
	merge1(in, 2, 2, 3)
	require.Equal(t, []int{2, 4, 1, 6}, in)

	in = []int{1, 4, 9, 13, 5, 6, 20}
	merge1(in, 0, 3, 6)
	require.Equal(t, []int{1, 4, 5, 6, 9, 13, 20}, in)

	in = []int{1, 4, 9, 13, 5, 6, 2}
	merge1(in, 0, 3, 5)
	require.Equal(t, []int{1, 4, 5, 6, 9, 13, 2}, in)
}

func TestMergeSort(t *testing.T) {
	in := []int{}
	out := MergeSort(in)
	assert.Nil(t, out)

	in = []int{7}
	out = MergeSort(in)
	assert.Equal(t, in, out)

	in = []int{7, 6, 2, 4, 1, 9, 3, 8, 0, 5}
	out = MergeSort(in)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, out)

	in = []int{7, 6, 2, 3, 1, 9, 3, 8, 0, 5}
	out = MergeSort(in)
	assert.Equal(t, []int{0, 1, 2, 3, 3, 5, 6, 7, 8, 9}, out)

	in = []int{3434, 3356, 67, 12334, 878667, 387}
	out = MergeSort(in)
	assert.Equal(t, []int{67, 387, 3356, 3434, 12334, 878667}, out)
}

func TestMergeSort1(t *testing.T) {
	in := []int{}
	MergeSort1(in, -1, len(in)-1)
	assert.Equal(t, []int{}, in)

	in = []int{7}
	MergeSort1(in, 0, len(in)-1)
	assert.Equal(t, []int{7}, in)

	in = []int{2, 4, 1, 6}
	MergeSort1(in, 0, len(in)-1)
	assert.Equal(t, []int{1, 2, 4, 6}, in)

	in = []int{11, 8, 3, 9, 7, 1, 2, 5}
	MergeSort1(in, 0, len(in)-1)
	assert.Equal(t, []int{1, 2, 3, 5, 7, 8, 9, 11}, in)
}
