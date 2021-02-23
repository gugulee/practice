package main

import (
	_ "embed"
	"fmt"
)

//go:embed README.md
var t string

func main() {
	fmt.Println(t)
}
