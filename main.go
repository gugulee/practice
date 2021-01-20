package main

import (
	"fmt"
)

type test struct {
	a int
}

func (t *test) add(a int) *test {
	t.a = a
	fmt.Printf("1 %p\n", t)
	return t
}

func main() {
    t := &test{}
    fmt.Printf("1 %p\n", t)
	out := t.add(10)
    fmt.Println(out)
}
