package kp

import (
	"fmt"
	"testing"
)

func Test_kp(t *testing.T) {
	items := []int{20, 19, 18, 7, 8}
	// items := []int{5, 4, 3, 2, 1}
	w := 30
	kp(items, w, 0, 0)
	fmt.Println(maxw)
}
