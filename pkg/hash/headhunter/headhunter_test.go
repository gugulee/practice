package headhunter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	h := New(0)
	fmt.Println(h)
}

func TestUpdate(t *testing.T) {
	h := New(100)

	in := map[int]int{1: 100, 2: 100, 3: 50, 4: 10, 5: 134,
		7: 110, 8: 200, 9: 12, 10: 43, 13: 300, 16: 145, 17: 123, 18: 301}

	for k, v := range in {
		h.Update(newNode(k, v))
	}

	// only one node{1,100}
	h.Update(newNode(1, 100))
	fmt.Println(h.skip.Print())
	fmt.Println("------------------------------------------------------------")

	// delete node{1,100} and insert node{1,210}
	n := newNode(1, 210)
	h.Update(n)
	require.Nil(t, h.skip.Search(newNode(1, 100)))
	require.NotNil(t, h.skip.Search(n))
	require.Equal(t, 210, h.node[1].score)

	n = newNode(12, 10)
	h.Update(n)
	require.NotNil(t, h.skip.Search(n))
}

func TestSearch(t *testing.T) {
	h := New(100)

	in := map[int]int{1: 100, 2: 100, 3: 50, 4: 10, 5: 134,
		7: 110, 8: 200, 9: 12, 10: 43, 13: 300, 16: 145, 17: 123, 18: 301}

	for k, v := range in {
		h.Update(newNode(k, v))
	}

	require.Equal(t, 100, h.Search(1))
	require.NotNil(t, h.skip.Search(newNode(1, 100)))

	require.Equal(t, 301, h.Search(18))
	require.NotNil(t, h.skip.Search(newNode(18, 301)))

	require.Equal(t, -100, h.Search(20))
	require.Nil(t, h.skip.Search(newNode(20, 100)))
}

func TestSearchRange(t *testing.T) {
	h := New(100)

	in := map[int]int{1: 100, 2: 100, 3: 50, 4: 10, 5: 134,
		7: 110, 8: 200, 9: 12, 10: 43, 13: 300, 16: 145, 17: 123, 18: 301, 19: 10}

	for k, v := range in {
		h.Update(newNode(k, v))
	}

	fmt.Println(h.skip.Print())
	fmt.Println("------------------------------------------------------------")

	outs := h.SearchRange(10, 100)
	require.ElementsMatch(t, []int{4, 19, 9, 10, 3, 1, 2}, outs)

	outs = h.SearchRange(100, 200)
	require.ElementsMatch(t, []int{1, 2, 7, 17, 5, 16, 8}, outs)
}

func TestDelete(t *testing.T) {
	h := New(100)

	in := map[int]int{1: 100, 2: 100, 3: 50, 4: 10, 5: 134,
		7: 110, 8: 200, 9: 12, 10: 43, 13: 300, 16: 145, 17: 123, 18: 301}

	for k, v := range in {
		h.Update(newNode(k, v))
	}

	h.Delete(1)
	require.Equal(t, -100, h.Search(1))
	require.Nil(t, h.skip.Search(newNode(1, 100)))

	h.Delete(1)
	require.Equal(t, -100, h.Search(1))
}
