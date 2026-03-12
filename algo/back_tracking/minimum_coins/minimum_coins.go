package mc

import (
	"fmt"
	"math"
)

var minCoin = math.MaxInt16
var coins = make([]int, 9)
var mem = make([][]bool, 10)

// items := []int{1, 3, 5}
func minimumCoins(items []int, w, curMinCoin int) {
	if 0 == w {
		fmt.Println(coins)
		if minCoin > curMinCoin {
			minCoin = curMinCoin
		}
		return
	}

	for i := range items {
		coins[curMinCoin] = items[i]
		if w-items[i] >= 0 {
			minimumCoins(items, w-items[i], curMinCoin+1)
		}
		coins[curMinCoin] = 0
	}
}

// items := []int{1, 3, 5}
func minimumCoins1(items []int, w, curw, curMinCoin int) {
	if curw == w {
		// fmt.Println(coins)
		if minCoin > curMinCoin {
			minCoin = curMinCoin
		}
		return
	}

	if mem[curMinCoin][curw] {
		return
	}

	mem[curMinCoin][curw] = true

	for i := range items {
		if curw+items[i] <= w {
			minimumCoins1(items, w, curw+items[i], curMinCoin+1)
		}
	}
}
