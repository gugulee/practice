package climbstairs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_climbStairs(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	// 假如 m := []int{1, 5, 7, 2, 3, 9}，第一个人住在 1 楼，第二个住在 5 楼，第三个人住在 7 楼等等
	tests := []struct {
		desc   string
		m      []int
		expect int
	}{
		{
			"case1",
			[]int{1, 5, 7, 2, 3, 9},
			5,
		},
		{
			"case2",
			[]int{1, 5, 3, 2, 3, 9},
			3,
		},
		{
			"case3",
			[]int{20, 3, 5, 20, 1, 7, 20},
			7,
		},
		{
			"case4",
			[]int{20, 3, 5, 20, 1, 20},
			20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			optimal := climbStairs(tt.m)
			require.Equal(tt.expect, optimal)
		})
	}
}

func Test_climbStairs1(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	// 假如 m := []int{1, 5, 7, 2, 3, 9}，第一个人住在 1 楼，第二个住在 5 楼，第三个人住在 7 楼等等
	tests := []struct {
		desc   string
		m      []int
		expect int
	}{
		{
			"case1",
			[]int{1, 5, 7, 2, 3, 9},
			3,
		},
		{
			"case2",
			[]int{1, 5, 3, 2, 3, 9},
			3,
		},
		{
			"case3",
			[]int{20, 3, 5, 20, 1, 7, 20},
			7,
		},
		{
			"case4",
			[]int{20, 3, 5, 20, 1, 20},
			5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			optimal := climbStairs1(tt.m)
			require.Equal(tt.expect, optimal)
		})
	}
}
