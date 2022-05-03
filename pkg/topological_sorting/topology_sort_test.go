package kahn

import (
	"testing"
)

func generateTestData(n int) (result [][]bool) {
	result = make([][]bool, n)
	for i := 0; i < n; i++ {
		result[i] = make([]bool, n)
	}

	result[1][0] = true
	result[1][3] = true

	result[2][1] = true

	result[4][1] = true

	return result
}

func Test_topoSortByKahn(t *testing.T) {
	data := generateTestData(5)
	topoSortByKahn(data)
}

func Test_topoSortByDFS(t *testing.T) {
	data := generateTestData(5)
	topoSortByDFS(data)
}
