package kp

import (
	"math"
)

var maxw = math.MinInt16

// kp return the max weight
// items := []int{20, 19, 18, 7, 8}
func kp(items []int, w, cw, i int) {
	if cw == w || i == len(items) {
		if cw > maxw {
			maxw = cw
		}
		return
	}

	kp(items, w, cw, i+1)
	if cw+items[i] <= w {
		kp(items, w, cw+items[i], i+1)
	}
}
