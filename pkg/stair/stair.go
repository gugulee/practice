package stair

import "fmt"

// oneStep represent the number of stairs when make a step
var oneStep = []int{1, 2}

func stair(n int, result []int) {
	if 0 == n {
		fmt.Println(result)
		return
	}
	if 1 == n {
		fmt.Println(append(result, 1))
		return
	}

	for _, s := range oneStep {
		var newResult []int
		newResult = append(newResult, result...) // copy result
		newResult = append(newResult, s)
		stair(n-s, newResult)
	}
}

func countStair(n int) int {
	if 1 == n {
		return 1
	}
	if 2 == n {
		return 2
	}

	return countStair(n-1) + countStair(n-2)
}
