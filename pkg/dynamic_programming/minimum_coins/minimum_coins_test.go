package mc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumCoins(t *testing.T) {
	coins := []int{1, 3, 5}
	amount := 9
	out := minimumCoins(coins, amount)
	require.Equal(t, 3, out)
}

func Test_minimumCoins1(t *testing.T) {
	coins := []int{1, 3, 5}
	amount := 9
	out := minimumCoins1(coins, amount)
	require.Equal(t, 3, out)
}

func Test_minimumCoins2(t *testing.T) {
	coins := []int{1, 3, 5}
	amount := 9
	mem := make([]int, amount+1)
	out := minimumCoins2(coins, mem, amount)
	require.Equal(t, 3, out)
}
