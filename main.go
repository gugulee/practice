package main

import "fmt"

func main() {
	var routes []int

	fmt.Printf("%p\n", routes)

	routes = make([]int, 0, 10)

	fmt.Printf("%p\n", routes)
}
