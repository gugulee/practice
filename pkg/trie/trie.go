package trie

import "fmt"

type Trie struct {
	root *Node
}

type Node struct {
	Child       []*Node
	Char        string
	Explanation string
	Prefix      string
}

func reverse(a []*Node) {
	if 0 == len(a) {
		return
	}

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func NewTrie() *Trie {
	trie := &Trie{
		root: NewNode("", ""),
	}
	return trie
}

func NewNode(c string, explanation string) *Node {
	return &Node{
		Char:        c,
		Explanation: explanation,
	}
}

func (t *Trie) Insert(word string, explanation string) {
	if 0 == len(word) {
		return
	}

	insert(t.root, word, word, explanation)
}

func insert(node *Node, word, origin, explanation string) {
	// in the last char of word, insert explanation
	if 0 == len(word) {
		node.Explanation = explanation
		node.Prefix = origin
		return
	}

	c := string(word[0])
	var tempNode *Node

	// find same node
	found := false
	for i := range node.Child {
		if nil != node.Child[i] && c == node.Child[i].Char {
			found = true
			tempNode = node.Child[i]
		}
	}

	// if not found, insert into the first child that is null
	if !found {
		newNode := NewNode(c, "")
		node.Child = append(node.Child, newNode)
		tempNode = newNode
	}

	insert(tempNode, word[1:], origin, explanation)
}

func (t *Trie) searchByStack() {
	var stack []*Node
	// push root node
	stack = append(stack, t.root)

	for len(stack) > 0 {
		// pop node
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// if it is a real word, print
		if 0 != len(current.Explanation) {
			fmt.Printf("%s %s\n", current.Prefix, current.Explanation)
		}

		// if it is leaf node, continue
		if 0 == len(current.Child) {
			continue
		}

		// reverse child
		reverse(current.Child)

		for _, child := range current.Child {
			stack = append(stack, child)
		}
	}
}

func (t *Trie) Print() {
	node := t.root
	printTrie(node)
}

func printTrie(node *Node) {
	if 0 != len(node.Explanation) {
		fmt.Printf("%s %s\n", node.Prefix, node.Explanation)
	}

	for i := range node.Child {
		if nil != node.Child[i] {
			printTrie(node.Child[i])
		}
	}
}

func (t *Trie) Find(target string) bool {
	if 0 == len(target) {
		return false
	}

	current := t.root
	for _, c := range target {
		ch := string(c)
		found := false
		for _, child := range current.Child {
			if nil != child && ch == child.Char {
				current = child
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if 0 != len(current.Explanation) {
		fmt.Printf("%s %s\n", target, current.Explanation)
		return true
	}

	return false
}

//func insert(node *Node, word string, explanation string) {
//	// in the last char of word, insert explanation
//	if 0 == len(word) {
//		node.Explanation = explanation
//		return
//	}
//
//	c := string(word[0])
//	var tempNode *Node
//
//	// find same node
//	found := 0
//	for i := range node.Child {
//		if nil == node.Child[i] {
//			continue
//		} else if c == node.Child[i].Char {
//			found = 1
//			tempNode = node.Child[i]
//		}
//	}
//
//	// if not found, insert into the first child that is null
//	if 0 == found {
//		var i int
//		for i = range node.Child {
//			if nil == node.Child[i] {
//				break
//			}
//		}
//		newNode := NewNode(c, "")
//		node.Child[i] = newNode
//		tempNode = newNode
//	}
//
//	insert(tempNode, word[1:], explanation)
//}
//
//func (t *Trie) Print() {
//	node := t.root
//	printTrie(node, "")
//}
//
//func printTrie(node *Node, word string) {
//	if nil == node.Child[0] {
//		fmt.Println(word)
//		return
//	}
//
//	for i := range node.Child {
//		if nil != node.Child[i] {
//			newWord := word
//			newWord += node.Child[i].Char
//			printTrie(node.Child[i], newWord)
//		}
//	}
//}
//
//func (t *Trie) Find(target string) bool {
//	if 0 == len(target) {
//		return false
//	}
//
//	node := t.root
//	for i := range node.Child {
//		if found := find(node.Child[i], target, ""); found {
//			return found
//		}
//	}
//	return false
//}
//
//func find(node *Node, target, explanation string) bool {
//	// target is empty and is a word exactly(explain is not empty), return true
//	if 0 == len(target) {
//		if 0 == len(explanation) {
//			return false
//		}
//		fmt.Println(explanation)
//		return true
//	}
//
//	// node is nil and target is not empty, return false
//	if nil == node {
//		return false
//	}
//
//	// if the char in node not match, return false
//	c := string(target[0])
//	if c != node.Char {
//		return false
//	}
//
//	for i := range node.Child {
//		if found := find(node.Child[i], target[1:], node.Explanation); found {
//			return found
//		}
//	}
//
//	return false
//}
