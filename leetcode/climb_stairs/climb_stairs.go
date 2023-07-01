package climbstairs

import (
	"math"
	"sort"
)

// 首先，楼梯只会停在 m 个人住的楼层，和楼梯的总层数没有关系，其实可以抽象成：
// 在数组 m 种，找一个数，使得数组中每个元素和它差的绝对值之和最小。

// brute force
func climbStairs(m []int) int {
	minSum, optimal := math.MaxInt64, -1

	// iterate all possible floors, find the minimum sum of absolute values
	for i := 0; i < len(m); i++ {
		// put the candidate floor to the first position
		m[0], m[i] = m[i], m[0]
		sum := 0

		for j := 1; j < len(m); j++ {
			sum += int(math.Abs(float64(m[0] - m[j])))
		}

		if sum < minSum {
			minSum = sum
			optimal = m[0]
		}

		// restore the order
		m[i], m[0] = m[0], m[i]
	}

	return optimal
}

// 简单证明：
// 将数组 m 分为两部分：部分 1 包含小于或等于中位数 M 的元素，部分 2 包含大于中位数 M 的元素。
// 假设，任选一个数字 X，并且 X 不是中位数，那么 X 必然是部分 1 或者部分 2 中的一个。
// 利用反证法，假设能取得一个数子 X，使得 X 和其它数字差的绝对值小于 M 和其它数字差的绝对值之和。
// 情况1，假设 X 在部分 1 中：
// 部分 1 中的任意数字与 X 的绝对值之和为 S1，部分 2 中的任意数字与 M 的绝对值之和为 S2, 那么必有 S1 <= S2；
// 部分 2 中的任意数字与 X 的绝对值之和为 S3，部分 2 中的任意数字与 M 的绝对值之和为 S4, 那么必有 S3 > S4；
// 所以没法保证，S1 与 S3 之和小于 S2 与 S4 之和。假设不成立，所以中位数 M 和其它数字差的绝对值之最小。
// 情况2，假设 X 在部分 2 中, 证明方法也是类似的。
// 所以得出结论：中位数 M 和其它数字差的绝对值之最小。

// sort and find the median
func climbStairs1(m []int) int {
	sort.Ints(m)
	n := len(m)

	// if n is even, midian is the middle two numbers, return the smaller one
	if n%2 == 0 {
		return m[(n-1)/2]
	}

	// if n is odd, return the median
	return m[n/2]
}
