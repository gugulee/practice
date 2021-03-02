package dijkstra

import (
	"fmt"
	set "github.com/deckarep/golang-set"
	"math"
)

func minInSlice(s []int, visit set.Set) int {
	min := math.MaxInt64
	index := -1
	for i := 0; i < len(s); i++ {
		if visit.Contains(i) {
			continue
		}

		if s[i] < min {
			min = s[i]
			index = i
		}
	}
	return index
}

func dijkstra(node [][]int, s int) {
	nodeNum := len(node)
	mw := make([]int, nodeNum)
	for i := 0; i < nodeNum; i++ {
		if i == s {
			mw[s] = 0
			continue
		}
		mw[i] = math.MaxInt32
	}

	visit := set.NewSet()

	for {
		min := minInSlice(mw, visit)
		if -1 == min {
			break
		}
		visit.Add(min)

		for j := 0; j < nodeNum; j++ {
			temp := mw[min] + node[min][j]
			if temp < mw[j] {
				mw[j] = temp
			}
		}
	}

	for i := 0; i < nodeNum; i++ {
		fmt.Printf("mw[%d] = %d\n", i, mw[i])
	}
}
