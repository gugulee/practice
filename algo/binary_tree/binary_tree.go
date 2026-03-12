package binarytree

import (
	"strings"

	set "github.com/deckarep/golang-set"
	"github.com/practice/pkg/tools"
)

// Node is the node of the binary tree
type Node struct {
	data        string
	left, right *Node
}

// BinaryTree ...
type BinaryTree struct {
	// the node that head.left and head.right point to is the root node
	head *Node
}

func newNode(data string) *Node {
	return &Node{data: data}
}

// New ...
func New() *BinaryTree {
	return &BinaryTree{newNode("")}
}

// PreOrder is the recursive traversal
func (bt *BinaryTree) PreOrder() string {
	if bt.head.left != bt.head.right {
		return ""
	}

	var r []string
	preOrder(bt.head.left, &r)

	return strings.Join(r, "->")
}

// PostOrder is the recursive traversal
func (bt *BinaryTree) PostOrder() string {
	if bt.head.left != bt.head.right {
		return ""
	}

	var r []string
	postOrder(bt.head.left, &r)

	return strings.Join(r, "->")
}

// InOrder is the recursive traversal
func (bt *BinaryTree) InOrder() string {
	if bt.head.left != bt.head.right {
		return ""
	}

	var r []string
	inOrder(bt.head.left, &r)

	return strings.Join(r, "->")
}

// PreOrder1 is the non-recursive traversal
func (bt *BinaryTree) PreOrder1() string {
	if bt.head.left != bt.head.right {
		return ""
	}

	var r []string
	var n *Node
	stack := []*Node{bt.head.left}
	length := len(stack)

	for 0 != length {
		n = stack[length-1]
		stack = stack[0 : length-1]

		r = append(r, n.data)

		if nil != n.right {
			stack = append(stack, n.right)
		}

		if nil != n.left {
			stack = append(stack, n.left)
		}

		length = len(stack)
	}

	return strings.Join(r, "->")
}

// PostOrder1 is the non-recursive traversal
func (bt *BinaryTree) PostOrder1() string {
	if bt.head.left != bt.head.right {
		return ""
	}

	var r []string
	var n *Node
	stack := []*Node{bt.head.left}
	visited := set.NewSet()
	length := len(stack)

	for 0 != length {
		n = stack[length-1]

		if visited.Contains(n.data) {
			r = append(r, n.data)
			stack = stack[0 : length-1]
		} else {
			if nil != n.right {
				stack = append(stack, n.right)
			}

			if nil != n.left {
				stack = append(stack, n.left)
			}

			visited.Add(n.data)
		}

		length = len(stack)
	}

	return strings.Join(r, "->")
}

// PostOrder2 is the non-recursive traversal
func (bt *BinaryTree) PostOrder2() string {
	if bt.head.left != bt.head.right {
		return ""
	}

	var r []string
	var n *Node
	stack := []*Node{bt.head.left}

	for 0 != len(stack) {
		n = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		r = append(r, n.data)

		if nil != n.left {
			stack = append(stack, n.left)
		}

		if nil != n.right {
			stack = append(stack, n.right)
		}
	}

	tools.Reverse(r)

	return strings.Join(r, "->")
}

// InOrder1 is the non-recursive traversal
func (bt *BinaryTree) InOrder1() string {
	if bt.head.left != bt.head.right {
		return ""
	}

	var r []string
	var n *Node
	stack := []*Node{bt.head.left}
	visited := set.NewSet()
	length := len(stack)

	for 0 != length {
		n = stack[length-1]
		stack = stack[0 : length-1]

		if visited.Contains(n.data) {
			r = append(r, n.data)
		} else {
			if nil != n.right {
				stack = append(stack, n.right)
			}

			stack = append(stack, n)

			if nil != n.left {
				stack = append(stack, n.left)
			}

			visited.Add(n.data)
		}

		length = len(stack)
	}

	return strings.Join(r, "->")
}

// InOrder2 is the non-recursive traversal
func (bt *BinaryTree) InOrder2() string {
	if bt.head.left != bt.head.right {
		return ""
	}

	var r []string
	stack := []*Node{}
	n := bt.head.left

	for 0 != len(stack) || nil != n {
		// push
		for nil != n {
			stack = append(stack, n)
			n = n.left
		}

		//pop
		n1 := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		r = append(r, n1.data)
		n = n1.right
	}

	return strings.Join(r, "->")
}

func preOrder(n *Node, r *[]string) {
	if nil == n {
		return
	}

	*r = append(*r, n.data)
	preOrder(n.left, r)
	preOrder(n.right, r)
}

func postOrder(n *Node, r *[]string) {
	if nil == n {
		return
	}

	postOrder(n.left, r)
	postOrder(n.right, r)
	*r = append(*r, n.data)
}

func inOrder(n *Node, r *[]string) {
	if nil == n {
		return
	}

	inOrder(n.left, r)
	*r = append(*r, n.data)
	inOrder(n.right, r)
}

// BFS ...
func (bt *BinaryTree) BFS() string {
	if bt.head.left != bt.head.right {
		return ""
	}

	var r []string
	queue := []*Node{bt.head.left}
	var n *Node

	for 0 != len(queue) {
		n = queue[0]
		queue = queue[1:]

		r = append(r, n.data)

		if nil != n.left {
			queue = append(queue, n.left)
		}

		if nil != n.right {
			queue = append(queue, n.right)
		}
	}

	return strings.Join(r, "->")
}
