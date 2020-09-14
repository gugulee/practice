package main

import (
	"fmt"
)

// Node is the node of the binary tree
type Node struct {
	data        string
	left, right *Node
}

func main() {
	a := &Node{data: "a"}
	stack := []*Node{a}
	fmt.Println(stack)

	stack = append(stack, a.left)
	fmt.Println(stack)
}
