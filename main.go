package main

import (
	"fmt"
	"sort"
)

type ySort [][]float64

func (s ySort) Len() int {
	return len(s)
}

func (s ySort) Less(i, j int) bool {
	return s[i][1] < s[j][1]
}

func (s ySort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	in := [][]float64{
		{7, 7},
		{10, 2},
		{30, 10},
		{1, 3},
		{4, 10},
		{5, 6},
	}

	r := 3
	q := 5

	right := make([][]float64, q-r+1)

	copy(right, in[r:q+1])
	
	fmt.Println(right)

	sort.Sort(ySort(right))

	fmt.Println(right)
}
