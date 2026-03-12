package md

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func generateTestData(n int) (result [][]int) {
	result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	result[0][0] = 1
	result[0][1] = 3
	result[0][2] = 5
	result[0][3] = 9

	result[1][0] = 2
	result[1][1] = 1
	result[1][2] = 3
	result[1][3] = 4

	result[2][0] = 5
	result[2][1] = 2
	result[2][2] = 6
	result[2][3] = 7

	result[3][0] = 6
	result[3][1] = 8
	result[3][2] = 4
	result[3][3] = 3

	return result
}

func Test_minimumDistance(t *testing.T) {
	data := generateTestData(4)
	minimumDistance(data, 0, 0, 0)
	require.Equal(t, 19, minDist)
}
