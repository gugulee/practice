package main

import (
	"fmt"
)

func main() {
	a := []int{1, 3, 4, 5, 6, 7}
	b := make([]int, 2*len(a))
	copy(b, a)

	fmt.Println(a)
	fmt.Printf("a=%p\n", a)
	fmt.Println(b)
	fmt.Printf("b=%p\n", b)
	a = b

	fmt.Println(a)
	fmt.Printf("a=%p\n", a)
}
