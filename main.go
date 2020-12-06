package main

import "fmt"

var sweetsSize = []int{3, 5, 7, 10, 13}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{8, 8}
	copy(a[2:3+1], b)
	fmt.Println(a)
}
