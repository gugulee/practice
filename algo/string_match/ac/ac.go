package ac

import "fmt"

const (
	defaultLen = 26 // assume that only have a-z letter in string
)

// Node is the node of trie
type Node struct {
	// store the character
	data rune

	// isEndingChar represent is the end of string
	isEndingChar bool
	child        []*Node

	// fail pointer
	fail *Node

	// when isEndingChar == true, record the length of string
	length int
}

// Trie ...
type Trie struct {
	root *Node
}

// NewNode ...
func NewNode(data rune) *Node {
	return &Node{
		data:  data,
		child: make([]*Node, defaultLen),
	}
}

// New ...
func New() *Trie {
	return &Trie{
		root: NewNode('/'),
	}
}

// Insert insert into trie with data
func (t *Trie) Insert(data string) {
	if 0 == len(data) {
		return
	}

	curr := t.root
	i := 0
	for ; i < len(data); i++ {
		idx := convertToSubscript(rune(data[i]))
		if nil == curr.child[idx] {
			curr.child[idx] = NewNode(rune(data[i]))
		}

		curr = curr.child[idx]
	}

	curr.isEndingChar = true
	curr.length = i
}

// Find find the trie with data
func (t *Trie) Find(data string) bool {
	if 0 == len(data) {
		return false
	}

	curr := t.root
	for _, d := range data {
		idx := convertToSubscript(d)

		// not found
		if nil == curr.child[idx] {
			return false
		}

		curr = curr.child[idx]
	}

	return curr.isEndingChar
}

// Print print the trie using dfs
func (t *Trie) Print() {
	var result string
	print(t.root, result)
}

func print(curr *Node, result string) {
	if curr.data != '/' {
		result += string(curr.data)
	}

	if curr.isEndingChar {
		fmt.Println(result)
	}

	for i := range curr.child {
		if nil != curr.child[i] {
			print(curr.child[i], result)
		}
	}
}

// ACmatch match multiple pattern in the main string,
// main = "abcd",
// pattern = []string{"c", "bc", "bcd", "abcd"},
// result is "abcd", "bcd", "bc" ,"c"
func (t *Trie) ACmatch(main string) {
	n := len(main)
	p := t.root
	for i := 0; i < n; i++ {
		idx := main[i] - 'a'
		for p.child[idx] == nil && p != t.root {
			p = p.fail // 失败指针发挥作用的地方
		}

		p = p.child[idx]
		if p == nil { // 如果没有匹配的，从root开始重新匹配
			p = t.root
		}
		tmp := p
		for tmp != t.root { // 打印出可以匹配的模式串
			if tmp.isEndingChar == true {
				pos := i - tmp.length + 1
				fmt.Printf("start pos: %d, length: %d, string: %s\n", pos, tmp.length, main[pos:pos+tmp.length])
			}
			tmp = tmp.fail
		}
	}
}

func (t *Trie) buildFailurePointer() {
	var queue []*Node

	// this is the end when find fail pointer
	t.root.fail = nil

	// handle t.root.child[x] solely
	for i := range t.root.child {
		if nil != t.root.child[i] {
			t.root.child[i].fail = t.root
			queue = append(queue, t.root.child[i])
		}
	}

	for 0 != len(queue) {
		p := queue[0]
		queue = queue[1:]

		// find the fail pointer of p.child[x]
		for i := range p.child {
			q := p.fail
			// pc is the node which find fail pointer
			pc := p.child[i]

			// when p is the end character( pc == nil), that is no child, continue
			if nil == pc {
				continue
			}

			// until q=t.root.fail
			for nil != q {
				// found
				if nil != q.child[i] {
					pc.fail = q.child[i]
					break
				}

				q = q.fail
			}

			// not found
			if nil == q {
				pc.fail = t.root
			}

			queue = append(queue, pc)
		}
	}
}

func convertToSubscript(c rune) int {
	return int(c - 'a')
}
