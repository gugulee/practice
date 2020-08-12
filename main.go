package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 4}
	b := &a
	fmt.Printf("%p\n", b)
	fmt.Println(b[0])
}
