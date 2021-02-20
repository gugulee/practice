package main

import "fmt"

func main() {
	a, b := 3, 5

	if 3 == a || b == 4 {
		if 3 == a && 4 == b {
			fmt.Println(a, b)
		}
		fmt.Println("hhh")
	}
}
