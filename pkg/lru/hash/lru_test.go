package hash

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func TestInsertTail(t *testing.T) {
	e := newEntry("")
	e.InsertTail("hello")
	e.InsertTail("world")
	e.InsertTail("we")
	e.InsertTail("are")

	fmt.Println(e.Print())
	require.Equal(t, "hello->world->we->are", e.Print())
}

func TestEntryDelete(t *testing.T) {
	e := newEntry("")
	e.InsertTail("hello")
	e.InsertTail("world")
	e.InsertTail("we")
	e.InsertTail("are")

	fmt.Println(e.Print())

	word := "hello"
	node := e.Delete(word)
	require.Equal(t, word, node.data)

	word = "are"
	node = e.Delete(word)
	require.Equal(t, word, node.data)

	word = "you"
	node = e.Delete(word)
	require.Nil(t, node)
}

func TestInsert(t *testing.T) {
	lru := New(0, 0)
	words := []string{"hello", "world", "we", "are", "helloa", "wea", "helloaa"}

	// insert
	for _, word := range words {
		lru.Insert(word)
		require.Equal(t, word, lru.last.data)
	}

	require.Equal(t, 7, lru.length)

	// validate hash collision list
	word := "hello"
	require.Equal(t, word, lru.entry[hash(word, lru.capacity)].hNext.data)
	require.Equal(t, "helloa", lru.entry[hash(word, lru.capacity)].hNext.hNext.data)
	require.Equal(t, "helloaa", lru.entry[hash(word, lru.capacity)].hNext.hNext.hNext.data)

	word = "world"
	require.Equal(t, word, lru.entry[hash(word, lru.capacity)].hNext.data)

	word = "we"
	require.Equal(t, word, lru.entry[hash(word, lru.capacity)].hNext.data)
	require.Equal(t, "wea", lru.entry[hash(word, lru.capacity)].hNext.hNext.data)

	word = "are"
	require.Equal(t, word, lru.entry[hash(word, lru.capacity)].hNext.data)

	// validate the double list
	require.Equal(t, "helloaa->wea->helloa->are->we->world->hello", lru.ReverstPrint())
}

func TestSearch(t *testing.T) {
	lru := New(0, 0)
	words := []string{"hello", "world", "we", "are", "helloa", "wea", "helloaa"}

	// insert
	for _, word := range words {
		lru.Insert(word)
	}

	fmt.Println(lru.ReverstPrint())
	fmt.Println("-----------------------------------")

	word := "hello"
	require.Equal(t, word, lru.Search(word).data)

	word = "helloa"
	require.Equal(t, word, lru.Search(word).data)

	word = "helloaa"
	require.Equal(t, word, lru.Search(word).data)

	word = "wea"
	require.Equal(t, word, lru.Search(word).data)

	word = "weaa"
	require.Nil(t, lru.Search(word))
}

func TestDelete(t *testing.T) {
	lru := New(0, 0)
	words := []string{"hello", "world", "we", "are", "helloa", "wea", "helloaa"}

	// insert
	for _, word := range words {
		lru.Insert(word)
		require.Equal(t, word, lru.last.data)
	}

	fmt.Println(lru.ReverstPrint())
	fmt.Println("-----------------------------------")

	word := "hello"
	lru.Delete(word)
	require.Nil(t, lru.Search(word))

	word = "we"
	lru.Delete(word)
	require.Nil(t, lru.Search(word))

	word = "helloaa"
	lru.Delete(word)
	require.Nil(t, lru.Search(word))

	require.Equal(t, lru.Search("wea"), lru.last)
	require.Equal(t, "wea->helloa->are->world", lru.ReverstPrint())
}

func TestLruCache(t *testing.T) {

}
