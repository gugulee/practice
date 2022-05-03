package translation

import (
	"testing"

	"github.com/practice/pkg/heap"
	"github.com/stretchr/testify/require"
)

func generateTestData() (result [][]int) {
	result = make([][]int, 3)

	result[0] = make([]int, 3)
	result[0][0] = 8
	result[0][1] = 6
	result[0][2] = 3

	result[1] = make([]int, 4)
	result[1][0] = 6
	result[1][1] = 5
	result[1][2] = 4
	result[1][3] = 2

	result[2] = make([]int, 2)
	result[2][0] = 9
	result[2][1] = 6

	return result
}

func Test_translation(t *testing.T) {
	data := generateTestData()
	topk := 5
	minHeap = heap.NewMin(topk + 1)
	translation(data, topk, 0, 0)
	require.ElementsMatch(t, []int{20, 19, 22, 23, 21}, minHeap.Slice())
}
