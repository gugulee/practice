package lis

import (
	"fmt"
	"math"
)

var maxLis = math.MinInt16
var maxResult []int

// a := []int{2, 9, 3, 6}
func longestIncreasingSubsequence(a, result []int, i int) {
	if i == len(a) {
		if isIncreasing(result) && maxLis < len(result) {
			maxLis = len(result)
			maxResult = make([]int, maxLis)
			copy(maxResult, result)
		}
		return
	}

	longestIncreasingSubsequence(a, result, i+1)

	longestIncreasingSubsequence(a, append(result, a[i]), i+1)
}

func isIncreasing(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

// a := []int{2, 9, 3, 6}
func longestIncreasingSubsequence1(a, result []int, i int) {
	lena := len(a)
	if i == lena {
		if maxLis < len(result) {
			maxLis = len(result)
			maxResult = make([]int, maxLis)
			copy(maxResult, result)
		}
		return
	}

	longestIncreasingSubsequence1(a, result, i+1)

	if len(result) == 0 || a[i] > result[len(result)-1] {
		longestIncreasingSubsequence1(a, append(result, a[i]), i+1)
	}
}

// a := []int{2, 9, 3, 6}
func longestIncreasingSubsequence2(a, result []int, prev, i int) {
	lena := len(a)
	if i == lena {
		fmt.Println(result)
		if maxLis < len(result) {
			maxLis = len(result)
			maxResult = make([]int, maxLis)
			copy(maxResult, result)
		}
		return
	}

	longestIncreasingSubsequence2(a, result, prev, i+1)

	if prev == -1 || a[i] > a[prev] {
		longestIncreasingSubsequence2(a, append(result, a[i]), i, i+1)
	}
}

// a := []int{2, 9, 3, 6}
func longestIncreasingSubsequence3(a []int, prev, lis, i int) {
	lena := len(a)
	if i == lena {
		if maxLis < lis {
			maxLis = lis
		}
		return
	}

	longestIncreasingSubsequence3(a, prev, lis, i+1)

	if prev == -1 || a[i] > a[prev] {
		longestIncreasingSubsequence3(a, i, lis+1, i+1)
	}
}
