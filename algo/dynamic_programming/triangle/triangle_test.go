package triangle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func generateTestData() (result [][]int) {
	result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = -1
		}
	}

	result[0][0] = 5

	result[1][0] = 7
	result[1][1] = 8

	result[2][0] = 2
	result[2][1] = 3
	result[2][2] = 4

	result[3][0] = 4
	result[3][1] = 9
	result[3][2] = 6
	result[3][3] = 1

	result[4][0] = 2
	result[4][1] = 7
	result[4][2] = 9
	result[4][3] = 4
	result[4][4] = 5

	return result
}

func Test_triangle(t *testing.T) {
	data := generateTestData()
	min := triangle(data)
	require.Equal(t, 20, min)
}

func Test_triangle1(t *testing.T) {
	data := generateTestData()
	min := triangle1(data)
	require.Equal(t, 20, min)
}
