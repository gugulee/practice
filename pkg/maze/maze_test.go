package maze

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func getRandomGreaterthanZero(n int) (r int) {
	for i := 0; i < 10000; i++ {
		r := rand.Intn(n)
		if 0 != r {
			return r
		}
	}

	return 0
}

func generateTestData(maxVertex, maxEdge int) (result [][]int, err error) {
	rand.Seed(time.Now().Unix())

	result = make([][]int, maxVertex)

	for sourceVertex := range result {
		edge := getRandomGreaterthanZero(maxEdge)
		if 0 == edge {
			return nil, fmt.Errorf("generate edge error")
		}

		result[sourceVertex] = make([]int, maxVertex)

		for j := 0; j < edge; j++ {
			targetVertex := rand.Intn(maxVertex)
			if targetVertex == sourceVertex {
				continue
			}

			result[sourceVertex][targetVertex] = 1
		}

	}

	return result, nil
}

func Test_printPrev(t *testing.T) {
	prev := []int{1, 3, 0, 4, 2}

	printPrev(prev, 0, 1)
}

func Test_maze(t *testing.T) {
	r := require.New(t)
	out, err := generateTestData(10, 10)
	r.NoError(err)

	fmt.Println(out)

	maze(out, 0, 8)
}

func Test_maze1(t *testing.T) {
	r := require.New(t)
	out, err := generateTestData(10, 10)
	r.NoError(err)

	fmt.Println(out)

	maze1(out, 0, 8)
}