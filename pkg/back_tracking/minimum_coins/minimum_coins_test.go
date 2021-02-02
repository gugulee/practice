package mc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumCoins(t *testing.T) {
	items := []int{1, 3, 5}
	w := 9
	minimumCoins(items, w, 0)
	require.Equal(t, 3, minCoin)
}

func Test_minimumCoins1(t *testing.T) {
	items := []int{1, 3, 5}
	w := 9

	for i := range mem {
		mem[i] = make([]bool, w+1)
	}

	minimumCoins1(items, w, 0, 0)
	require.Equal(t, 3, minCoin)
}
