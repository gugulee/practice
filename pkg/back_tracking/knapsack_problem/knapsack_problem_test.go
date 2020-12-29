package kp

import (
	"fmt"
	"testing"
)

func Test_kp(t *testing.T) {
	items := []int{20, 19, 18, 7, 8}
	w := 30
	kp(items, w, 0, 0)
	fmt.Println(maxw)
}
