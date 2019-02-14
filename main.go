package main

import (
	"fmt"
)

func main() {
	var a int8 = -2
	fmt.Printf("%d\n", uint8(a))
	fmt.Printf("%d\n", ^uint8(-a)+1)
}
