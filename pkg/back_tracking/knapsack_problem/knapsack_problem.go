package kp

import (
	"fmt"
	"math"
)

var maxw = math.MinInt16

// kp return the max weight
// items := []int{20, 19, 18, 7, 8}
func kp(items []int, w, cw, i int) {
	fmt.Println("maxw=", maxw, "cw=", cw, "i=", i)
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
