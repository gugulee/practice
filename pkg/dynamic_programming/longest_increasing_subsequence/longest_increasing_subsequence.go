package lis

import (
	"fmt"

	"github.com/practice/pkg/tools"
)

/*
a := []int{4, 10, 4, 3, 8, 9}
-------------------------------------------------------
	4	10	4	3	8	9
	1	2	1	1	2	3
*/

func longestIncreasingSubsequence(a []int) int {
	lena := len(a)
	dp := make([]int, lena)
	// base case：dp 数组全都初始化为 1

	for i := 0; i < lena; i++ {
		dp[i] = 1
	}

	for i := 0; i < lena; i++ {
		for j := 0; j < i; j++ {
			if a[i] > a[j] {
				dp[i] = tools.Max(dp[i], dp[j]+1)
			}
		}
	}
	fmt.Println(dp)
	res := 0
	for i := 0; i < len(dp); i++ {
		res = tools.Max(res, dp[i])
	}
	return res
}
