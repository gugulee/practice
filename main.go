package main

import (
	"fmt"
)

func main() {
	var a uint64 = 0x112233445566
	var b uint32 = uint32(a)
	fmt.Printf("0x%x\n", a)
	fmt.Printf("0x%x\n", b)
}
