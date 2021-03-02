package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	fmt.Printf("1%p \n", a)
	a = append(a, 231)
	fmt.Printf("1%p \n", a)
	fmt.Println(a)
}
