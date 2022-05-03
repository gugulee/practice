package binarysearchtree

import (
	"github.com/practice/pkg/constants"
)

// Node is the node of the binary tree
type Node struct {
	data        int
	left, right *Node
}

// BinarySearchTree ...
type BinarySearchTree struct {
	// the node that head.right point to is the root node
	head *Node
}

func newNode(data int) *Node {
	return &Node{data: data}
}

// New ...
func New() *BinarySearchTree {
	return &BinarySearchTree{newNode(constants.MinInt)}
}

// Insert ...
func (bst *BinarySearchTree) Insert(data int) {
	node := bst.head

	for nil != node {
		if data > node.data {
			if nil == node.right {
				node.right = newNode(data)
				return
			}
			node = node.right
		} else if data < node.data {
			if nil == node.left {
				node.left = newNode(data)
				return
			}
			node = node.left
		}
	}
}

// Search ...
func (bst *BinarySearchTree) Search(data int) *Node {
	node := bst.head
	for nil != node {
		if data > node.data {
			node = node.right
		} else if data < node.data {
			node = node.left
		} else {
			return node
		}
	}

	return nil
}

// Delete ...
func (bst *BinarySearchTree) Delete(data int) {
	// prev is the father node, and p is the deleted node
	var prev *Node
	p := bst.head

	for nil != p && data != p.data {
		prev = p
		if data > p.data {
			p = p.right
		} else {
			p = p.left
		}
	}

	// not found
	if nil == p {
		return
	}

	/* the deleted node has left child node and right child node */
	if nil != p.left && nil != p.right {
		// find the minimum node in the right sub-tree of p
		minPrev := p
		minP := p.right
		for nil != minP.left {
			minPrev = minP
			minP = minP.left
		}

		// replace p's data with minP's data
		p.data = minP.data

		// delete minP.
		// minP.left must be empty, minP is not the minimum node otherwise
		p = minP
		prev = minPrev
	}

	/* the deleted node is the leaf node or has one child node(left child node or right child node) */
	// tmp is the p's child node
	var tmp *Node
	if nil != p.left {
		tmp = p.left
	} else {
		tmp = p.right
	}

	if prev.left == p {
		prev.left = tmp
	} else {
		prev.right = tmp
	}
}

// BFS ...
func (bst *BinarySearchTree) BFS() (r []int) {
	queue := []*Node{bst.head.right}
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

	return r
}

// FindMin find the minimum node
func (bst *BinarySearchTree) FindMin() *Node {
	node := bst.head.right
	if nil == node {
		return nil
	}

	for nil != node.left {
		node = node.left
	}

	return node
}

// FindMax find the maximum node
func (bst *BinarySearchTree) FindMax() *Node {
	node := bst.head.right
	if nil == node {
		return nil
	}

	for nil != node.right {
		node = node.right
	}

	return node
}

// InOrder return the positive sequence
func (bst *BinarySearchTree) InOrder() (r []int) {
	node := bst.head.right
	if nil == node {
		return nil
	}

	stack := []*Node{}

	for 0 != len(stack) || nil != node {
		for nil != node {
			stack = append(stack, node)
			node = node.left
		}

		// pop
		node1 := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		r = append(r, node1.data)
		node = node1.right
	}

	return r
}

// Height return the height of tree
func (bst *BinarySearchTree) Height() int {
	node := bst.head.right
	if nil == node {
		return -1
	}

	return height(node)
}

func height(node *Node) int {
	if nil == node {
		return -1
	}

	h1 := height(node.left) + 1
	h2 := height(node.right) + 1

	if h1 > h2 {
		return h1
	}

	return h2
}
