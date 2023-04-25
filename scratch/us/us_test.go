package us

import (
	"fmt"
	"testing"
)

func Test_test(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	var funcs []func()

	for _, num := range numbers {
		funcs = append(funcs, func() {
			fmt.Println(num * num)
		})
	}

	for _, f := range funcs {
		f()
	}
}
