package trie

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

func TestInset(t *testing.T) {
	fi, err := os.Open(`word`)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	trie := NewTrie()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if 0 == len(string(a)) {
			continue
		}
		d := strings.Index(string(a), " ")
		trie.Insert(string(a)[0:d], string(a)[d+1:])
	}
	//trie.Print()
	out := trie.Find("Antarctic")
	assert.Equal(t, true, out)
}

func TestSearchByStackt(*testing.T) {
	fi, err := os.Open(`word`)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	trie := NewTrie()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if 0 == len(string(a)) {
			continue
		}
		d := strings.Index(string(a), " ")
		trie.Insert(string(a)[0:d], string(a)[d+1:])
	}
	trie.searchByStack()
}
