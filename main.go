package main

import (
	"fmt"
)

func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

func main() {
	fmt.Println(maxDepth(32))
}
