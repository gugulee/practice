package main

import "fmt"

var sweetsSize = []int{3, 5, 7, 10, 13}

func main() {
	a := "51234"
	i := 0
	for i < len(a) {
		if i == len(a)-1 {
			break
		}

		if a[i] > a[i+1] {
			a = a[:i] + a[i+1:]
		} else {
			i++
		}
	}

	fmt.Println(a)
}
