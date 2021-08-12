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
	lru := New(0, 7)
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
	require.Equal(t, "helloaa<-wea<-helloa<-are<-we<-world<-hello", lru.ReverstPrint())
	require.Equal(t, "hello", lru.head.next.data)

	// if lru is full, can not insert node in lru
	lru.Insert("practice")
	require.Equal(t, 7, lru.length)
	require.Equal(t, "helloaa<-wea<-helloa<-are<-we<-world<-hello", lru.ReverstPrint())
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

	require.Equal(t, "wea", lru.last.data)
	require.Equal(t, "wea<-helloa<-are<-world", lru.ReverstPrint())
}

func TestDeleteNode(t *testing.T) {
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
	lru.DeleteNode(lru.Search(word))
	require.Nil(t, lru.Search(word))

	word = "we"
	lru.DeleteNode(lru.Search(word))
	require.Nil(t, lru.Search(word))

	word = "helloaa"
	lru.DeleteNode(lru.Search(word))
	require.Nil(t, lru.Search(word))

	require.Equal(t, "wea", lru.last.data)
	require.Equal(t, "wea<-helloa<-are<-world", lru.ReverstPrint())
}

func TestLru(t *testing.T) {
	lru := New(0, 7)
	words := []string{"hello", "world", "we", "are", "helloa", "wea"}

	// insert
	for _, word := range words {
		lru.Lru(word)
		require.Equal(t, word, lru.last.data)
	}

	require.Equal(t, "wea<-helloa<-are<-we<-world<-hello", lru.ReverstPrint())

	// not full and not found
	word := "helloaa"
	lru.Lru(word)
	require.Equal(t, "helloaa<-wea<-helloa<-are<-we<-world<-hello", lru.ReverstPrint())

	// delete "hello" and insert "app" when full
	word = "app"
	lru.Lru(word)
	require.Equal(t, "app<-helloaa<-wea<-helloa<-are<-we<-world", lru.ReverstPrint())

	// delete "wea" and insert "wea" when found
	word = "wea"
	lru.Lru(word)
	require.Equal(t, "wea<-app<-helloaa<-helloa<-are<-we<-world", lru.ReverstPrint())
}
