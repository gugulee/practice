package main

import (
	"fmt"
)

func main() {
	in := []int{
		7, 7,
	}

	right := make([]int, len(in))


	in[0] = 100
	right[0] = 200

	fmt.Println(right)
	fmt.Println(in)
	fmt.Println(6/3*3)
}
