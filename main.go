package main

import (
	"fmt"
)

func main() {
	a := 4 * 0
	fmt.Println(0x3039 >> a & 0xf)
	a = 4 * 1
	fmt.Println(0x3039 >> a & 0xf)
	a = 4 * 2
	fmt.Println(0x3239 >> a & 0xf)
}
