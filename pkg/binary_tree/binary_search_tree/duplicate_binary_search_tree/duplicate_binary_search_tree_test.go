package duplicatebinarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsert(t *testing.T) {
	requires := require.New(t)
	t.Run("1", func(t *testing.T) {
		bst := New()
		data := []int{13, 8, 18, 6, 10, 16, 20}

		for _, d := range data {
			bst.Insert(d)
		}

		out := bst.BFS()
		requires.Equal([]int{13, 8, 18, 6, 10, 16, 20}, out)

		bst.Insert(18)
		out = bst.BFS()
		requires.Equal([]int{13, 8, 18, 6, 10, 16, 20, 18}, out)
	})

	t.Run("2", func(t *testing.T) {
		bst := New()
		data := []int{13, 8, 18, 6, 10, 16}

		for _, d := range data {
			bst.Insert(d)
		}

		out := bst.BFS()
		requires.Equal([]int{13, 8, 18, 6, 10, 16}, out)

		bst.Insert(18)
		out = bst.BFS()
		requires.Equal([]int{13, 8, 18, 6, 10, 16, 18}, out)
	})

	t.Run("3", func(t *testing.T) {
		bst := New()
		data := []int{13, 8, 18, 6, 10, 16, 30, 20, 40}

		for _, d := range data {
			bst.Insert(d)
		}

		out := bst.BFS()
		requires.Equal([]int{13, 8, 18, 6, 10, 16, 30, 20, 40}, out)

		bst.Insert(18)
		out = bst.BFS()
		requires.Equal([]int{13, 8, 18, 6, 10, 16, 30, 20, 40, 18}, out)
	})
}

func TestSearch(t *testing.T) {
	requires := require.New(t)
	bst := New()
	data := []int{13, 8, 18, 6, 10, 16, 20, 18}

	for _, d := range data {
		bst.Insert(d)
	}

	// verification
	d := 13
	out := bst.Search(d)
	requires.Equal(1, len(out))

	d = 20
	out = bst.Search(d)
	requires.Equal(1, len(out))

	d = 1
	out = bst.Search(d)
	requires.Equal(0, len(out))
	requires.Nil(out)

	d = 18
	out = bst.Search(d)
	requires.Equal(2, len(out))
}

func TestDelete(t *testing.T) {
	requires := require.New(t)
	bst := New()
	data := []int{13, 8, 18, 6, 10, 16, 20, 18, 19, 18}

	for _, d := range data {
		bst.Insert(d)
	}

	bst.Delete(6)
	requires.Equal([]int{13, 8, 18, 10, 16, 20, 18, 19, 18}, bst.BFS())

	bst.Delete(18)
	requires.Equal([]int{13, 8, 19, 10, 16, 20}, bst.BFS())

	bst.Delete(13)
	requires.Equal([]int{16, 8, 19, 10, 20}, bst.BFS())
}
