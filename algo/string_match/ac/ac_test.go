package ac

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_convertToSubscript(t *testing.T) {
	for _, d := range "abcdz" {
		out := convertToSubscript(d)
		fmt.Println(out)
	}
}

func Test_buildFailurePointer(t *testing.T) {
	r := require.New(t)
	trie := New()

	words := []string{"ce", "bc", "bcd", "abce"}

	for _, w := range words {
		trie.Insert(w)
	}

	trie.buildFailurePointer()

	/* the depth is 0 */
	node := trie.root
	r.Equal('/', node.child[0].fail.data)
	r.Equal('/', node.child[1].fail.data)
	r.Equal('/', node.child[2].fail.data)

	/* the depth is 1 */
	node = trie.root.child[0]
	r.Equal('b', node.child[1].fail.data)

	node = trie.root.child[1]
	r.Equal('c', node.child[2].fail.data)

	node = trie.root.child[2]
	r.Equal('/', node.child[4].fail.data)

	/* the depth is 2 */
	node = trie.root.child[0].child[1]
	r.Equal('c', node.child[2].fail.data)

	node = trie.root.child[1].child[2]
	r.Equal('/', node.child[3].fail.data)

	/* the depth is 3 */
	node = trie.root.child[0].child[1].child[2]
	r.Equal('e', node.child[4].fail.data)
}

func Test_Insert(t *testing.T) {
	r := require.New(t)
	trie := New()

	words := []string{"c", "bc", "bcd", "abcd"}

	for _, w := range words {
		trie.Insert(w)
	}

	r.Equal('/', trie.root.data)

	r.Equal('a', trie.root.child[0].data)
	r.Equal('b', trie.root.child[1].data)
	r.Equal('c', trie.root.child[2].data)
	r.True(trie.root.child[2].isEndingChar)
	r.Equal(1, trie.root.child[2].length)

	r.Equal('b', trie.root.child[0].child[1].data)
	r.Equal('c', trie.root.child[1].child[2].data)
	r.True(trie.root.child[1].child[2].isEndingChar)
	r.Equal(2, trie.root.child[1].child[2].length)

	r.Equal('c', trie.root.child[0].child[1].child[2].data)
	r.Equal('d', trie.root.child[1].child[2].child[3].data)
	r.True(trie.root.child[1].child[2].child[3].isEndingChar)
	r.Equal(3, trie.root.child[1].child[2].child[3].length)

	r.Equal('d', trie.root.child[0].child[1].child[2].child[3].data)
	r.True(trie.root.child[0].child[1].child[2].child[3].isEndingChar)
	r.Equal(4, trie.root.child[0].child[1].child[2].child[3].length)
}

func Test_Find(t *testing.T) {
	r := require.New(t)
	trie := New()

	words := []string{"c", "bc", "bcd", "abcd"}

	for _, w := range words {
		trie.Insert(w)
	}

	r.True(trie.Find("c"))
	r.True(trie.Find("bc"))
	r.True(trie.Find("bcd"))
	r.True(trie.Find("abcd"))
	r.False(trie.Find("abce"))
	r.False(trie.Find("abc"))
	r.False(trie.Find("d"))
}

func Test_Print(t *testing.T) {
	trie := New()

	words := []string{"c", "bc", "bcd", "abcd"}

	for _, w := range words {
		trie.Insert(w)
	}

	trie.Print()
}

func Test_ACmatch(t *testing.T) {
	r := require.New(t)
	trie := New()

	words := []string{"c", "bc", "bcd", "abcd"}
	main := "abcd"

	for _, w := range words {
		trie.Insert(w)
	}

	trie.buildFailurePointer()

	trie.ACmatch(main)

	_ = r
}
