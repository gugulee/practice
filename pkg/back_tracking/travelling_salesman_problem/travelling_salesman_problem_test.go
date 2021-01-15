package tsp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func generateTestData() (result [][]int) {
	result = make([][]int, 4)
	for i := range result {
		result[i] = make([]int, 4)
	}

	result[0][1] = 30
	result[0][2] = 6
	result[0][3] = 4

	result[1][0] = 30
	result[1][2] = 5
	result[1][3] = 10

	result[2][0] = 6
	result[2][1] = 5
	result[2][3] = 20

	result[3][0] = 4
	result[3][1] = 10
	result[3][2] = 20

	return result
}

func Test_tsp(t *testing.T) {
	data := generateTestData()
	printData(data)
	tsp(data, 0, 0, 0, []int{0})
	require.Equal(t, 25, minSum)
	require.Equal(t, []int{0, 2, 1, 3, 0}, bestPath)
}
