package main

import (
	"fmt"
)

// key is the app id
type Limiter map[string]limiterAlgo

// key is the path
type limiterAlgo map[string]string

func main() {
	a := make(limiterAlgo)
	a["/user"] = "hello"
	b := make(Limiter)
	b["app-1"] = a

	fmt.Println(b)
}
