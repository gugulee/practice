package tsp

import (
	"fmt"
	"math"
)

var minSum = math.MaxInt16
var bestPath = []int{}

func tsp(data [][]int, cur, start, sum int, path []int) {
	if len(path) == len(data) {
		sum += data[cur][start]
		path = append(path, start)
		if minSum > sum {
			minSum = sum
			bestPath = path
		}
		return
	}

	for i := range data[cur] {
		// except for the cur node and the visited node
		if cur == i || contains(path, i) {
			continue
		}

		newPath := make([]int, len(path))
		copy(newPath, path)
		newPath = append(newPath, i)
		if data[cur][i]+sum < minSum {
			tsp(data, i, start, data[cur][i]+sum, newPath)
		}
	}
}

func printData(in [][]int) {
	for row := 0; row < len(in); row++ {
		for column := 0; column < len(in[row]); column++ {
			fmt.Printf("%d ", in[row][column])
		}
		fmt.Println()
	}
}

func contains(a []int, t int) bool {
	for i := range a {
		if t == a[i] {
			return true
		}
	}

	return false
}
