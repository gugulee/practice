package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := true
	var b int32
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
}
