package mc

import "testing"

func Test_minimumCoins(t *testing.T) {
	items := []int{1, 3, 5}
	w := 9
	mimimumCoins(items, w, 0)
}
