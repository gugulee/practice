package matrixmultiply

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func generateTestData(n int) (result [][]int) {

	result = make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = rand.Intn(10)
		}
	}

	return result
}

func Test_matrixMultiply(t *testing.T) {
	a := [][]int{
		{1, 2},
		{3, 4},
	}

	b := [][]int{
		{4, 3},
		{2, 1},
	}

	out := matrixMultiply(a, b)
	require.Equal(t, [][]int{
		{8, 5},
		{20, 13},
	}, out)

	a = generateTestData(4)
	b = generateTestData(4)
	out = matrixMultiply(a, b)
	print(a, b)
	fmt.Println(out)
}

func Test_divideAndConquer(t *testing.T) {
	// a := [][]int{
	// 	{1, 2},
	// 	{3, 4},
	// }

	// b := [][]int{
	// 	{4, 3},
	// 	{2, 1},
	// }

	a := [][]int{
		{2},
	}

	b := [][]int{
		{3},
	}

	print(a, b)
	out := divideAndConquer(a, b)
	require.Equal(t, [][]int{
		{8, 5},
		{20, 13},
	}, out)
}
