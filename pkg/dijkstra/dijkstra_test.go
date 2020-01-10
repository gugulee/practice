package dijkstra

import (
	set "github.com/deckarep/golang-set"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func generateRelation(userMax int) (node [][]int) {
	node = make([][]int, userMax)
	for i := range node {
		node[i] = make([]int, userMax)
		for j := range node[i] {
			node[i][j] = math.MaxInt32
		}
	}

	node[0][1] = 5
	node[0][2] = 3
	node[0][3] = 2
	node[0][4] = 4

	node[1][5] = 3

	node[2][1] = 2
	node[2][6] = 1

	node[3][6] = 4
	node[3][8] = 8

	node[4][3] = 1
	node[4][8] = 6

	node[5][7] = 1

	node[6][5] = 1
	node[6][8] = 2

	node[8][7] = 4

	return
}

func TestMinInSlice(t *testing.T) {
	in := []int{215, 3, 32, 54, 32, 13531, 314314, 6566, 34343, 777, 234}
	out := minInSlice(in, set.NewSet())
	require.Equal(t, 1, out)

	in = []int{215, 3123, 32, 54, 32, 13531, 314314, 6566, 34343, 777, 234}
	out = minInSlice(in, set.NewSet())
	require.Equal(t, 2, out)

	in = []int{215, 3123, 3, 54, 13531, 314314, 6566, 34343, 777, 234}
	out = minInSlice(in, set.NewSet(2))
	require.Equal(t, 3, out)

	in = []int{215, 3123, 3, 54, 13531, 314314, 6566, 34343, 777, 234}
	out = minInSlice(in, set.NewSet(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	require.Equal(t, -1, out)
}

// 0 5 3 2 4 5 4 6 6
func TestDijkstra(t *testing.T) {
	node := generateRelation(9)
	//for i := range node {
	//	for j := range node[i] {
	//		if math.MaxInt32 != node[i][j] {
	//			fmt.Printf("node[%d,%d] = %d,", i, j, node[i][j])
	//		}
	//	}
	//	fmt.Println()
	//}

	Dijkstra(node, 0)
}
