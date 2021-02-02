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

func Test_kp1(t *testing.T) {
	items := []int{2, 2, 4, 6, 3}
	w := 9

	mem = make([][]bool, len(items))
	for i := range mem {
		mem[i] = make([]bool, w+1)
	}

	kp1(items, w, 0, 0)
	fmt.Println(maxw)
}
