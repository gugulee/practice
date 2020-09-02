package wordspellcheck

const (
	capacity = 26
)

// WordTable store the words
type WordTable struct {
	root *Node
}

// Node store the character and the starting address of the next node
type Node struct {
	// next is the starting address of the next node
	next [capacity]*Node

	// we only store value in the last node
	value string
}

// hash convert character to int,
// such as 'a' -> 0, 'b' -> 1 and 'z' -> 25
func hash(character byte) int {
	return int(character - 'a')
}

func newNode(v string) *Node {
	return &Node{
		next:  [capacity]*Node{},
		value: v,
	}
}

// New return the WordHash
func New() *WordTable {
	return &WordTable{root: newNode("")}
}

// Insert the word into the word hash table
func (w *WordTable) Insert(word string) {
	var v string
	length := len(word)
	node := w.root
	for i := 0; i < length; i++ {
		hashValue := hash(word[i])
		v = ""

		// only insert value in the last node
		if i == length-1 {
			v = word
		}

		if nil == node.next[hashValue] {
			node.next[hashValue] = newNode(v)
		} else {
			node.next[hashValue].value = v
		}

		node = node.next[hashValue]
	}
}

// Search search the the word hash table with the word
func (w *WordTable) Search(word string) bool {
	length := len(word)
	node := w.root
	for i := 0; i < length; i++ {
		hashValue := hash(word[i])
		if nil == node.next[hashValue] {
			return false
		}

		if i == length-1 {
			if node.next[hashValue].value == word {
				return true
			}
		}

		node = node.next[hashValue]
	}

	return false
}
