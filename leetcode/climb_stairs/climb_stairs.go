package climbstairs

import (
	"math"
	"sort"
)

// 首先，楼梯只会停在 m 个人住的楼层，和楼梯的总层数没有关系，其实可以抽象成：
// 在数组 m 中，找一个数，使得数组中每个元素和它差的绝对值之和最小。

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
//*********************************************************************************
// 反证法逻辑确实存在问题，换一个思路：
// 对于 1,2,3,5,7,9 这个数组（保证有序），对于住在 1 楼和 9 楼的人来说，他们爬楼的总数是不变的，比如停在 9 楼，总爬楼数是 8，停在 3 楼，总爬楼数是8 （3-1+9-3=8）;
// 对于 2 楼和 7 楼的人来说，他们爬楼的总数也是不变的，比如停在 7 楼，总爬楼数是 5，停在 4 楼，总爬楼数是 5 （4-2+7-4=5）。
// 对于 3 楼和 5 楼的人来说，也是同理。
// 所以，我们楼梯只需要停在中间的楼层，就可以保证总爬楼数最小。
// 对于这个例子，停在 3 楼或 5 楼都可以。

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
