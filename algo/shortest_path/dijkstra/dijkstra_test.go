package dijkstra

import (
	"testing"
)

func generateRelation(userMax int) (node [][]int) {
	node = make([][]int, userMax)
	for i := range node {
		node[i] = make([]int, userMax)
	}

	node[0][1] = 10
	node[0][4] = 15

	node[1][2] = 15
	node[1][3] = 2

	node[2][5] = 5

	node[3][2] = 1
	node[3][5] = 12

	node[4][5] = 10

	return
}

func TestDijkstra(t *testing.T) {
	node := generateRelation(6)
	dijkstra(node, 0, 5)
}
