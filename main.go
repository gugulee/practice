package main

import "fmt"

var sweetsSize = []int{3, 5, 7, 10, 13}

func main() {
	j := 0
	for j, s := range sweetsSize {
		_ = s
		j = j
	}
	fmt.Println(j)
}
