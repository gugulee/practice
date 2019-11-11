package main

import "fmt"

func main() {
	a := []int{7, 6, 2, 4, 1, 9, 3, 8, 0}
	for i := 0; i < len(a); {
		if a[i] < 10 {
			fmt.Println(a)
			a = a[1:]
		}

	}
}
