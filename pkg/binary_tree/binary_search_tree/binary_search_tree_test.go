package binarysearchtree

import (
	// "sort"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsert(t *testing.T) {
	requires := require.New(t)
	bst := New()
	data := []int{33, 17, 50, 13, 18, 34, 58, 16, 25, 51, 66, 19, 27, 55}

	for _, d := range data {
		bst.Insert(d)
	}

	out := bst.BFS()

	requires.Equal([]int{33, 17, 50, 13, 18, 34, 58, 16, 25, 51, 66, 19, 27, 55}, out)
}

func TestSearch(t *testing.T) {
	requires := require.New(t)
	bst := New()
	data := []int{33, 16, 50, 13, 18, 34, 58, 15, 17, 25, 51, 66, 19, 27, 55}

	for _, d := range data {
		bst.Insert(d)
	}

	// verification
	d := 33
	out := bst.Search(d)
	requires.Equal(d, out.data)

	d = 13
	out = bst.Search(d)
	requires.Equal(d, out.data)

	d = 55
	out = bst.Search(d)
	requires.Equal(d, out.data)

	d = 1
	out = bst.Search(d)
	requires.Nil(out)
}

func TestFindMin(t *testing.T) {
	requires := require.New(t)
	t.Run("normal", func(t *testing.T) {
		bst := New()
		data := []int{33, 16, 50, 13, 18, 34, 58, 15, 17, 25, 51, 66, 19, 27, 55}

		for _, d := range data {
			bst.Insert(d)
		}

		// verification
		out := bst.FindMin()
		requires.Equal(13, out.data)
	})

	t.Run("abnormal", func(t *testing.T) {
		bst := New()

		// verification
		out := bst.FindMin()
		requires.Nil(out)
	})
}

func TestFindMax(t *testing.T) {
	requires := require.New(t)
	t.Run("normal", func(t *testing.T) {
		bst := New()
		data := []int{33, 16, 50, 13, 18, 34, 58, 15, 17, 25, 51, 66, 19, 27, 55}

		for _, d := range data {
			bst.Insert(d)
		}

		// verification
		out := bst.FindMax()
		requires.Equal(66, out.data)
	})

	t.Run("abnormal", func(t *testing.T) {
		bst := New()

		// verification
		out := bst.FindMax()
		requires.Nil(out)
	})
}

func TestDelete(t *testing.T) {
	requires := require.New(t)
	bst := New()
	data := []int{33, 16, 50, 13, 18, 34, 58, 15, 17, 25, 51, 66, 19, 27, 55}

	for _, d := range data {
		bst.Insert(d)
	}

	bst.Delete(13)
	requires.Equal([]int{33, 16, 50, 15, 18, 34, 58, 17, 25, 51, 66, 19, 27, 55}, bst.BFS())

	bst.Delete(18)
	requires.Equal([]int{33, 16, 50, 15, 19, 34, 58, 17, 25, 51, 66, 27, 55}, bst.BFS())

	bst.Delete(55)
	requires.Equal([]int{33, 16, 50, 15, 19, 34, 58, 17, 25, 51, 66, 27}, bst.BFS())
}

func TestInOrder(t *testing.T) {
	requires := require.New(t)
	bst := New()
	data := []int{33, 16, 50, 13, 18, 34, 58, 15, 17, 25, 51, 66, 19, 27, 55}

	for _, d := range data {
		bst.Insert(d)
	}

	sort.Ints(data)
	requires.Equal(data, bst.InOrder())
}

func TestHeight(t *testing.T) {
	requires := require.New(t)
	bst := New()
	data := []int{33, 16, 50, 13, 18, 34, 58, 15, 17, 25, 51, 66, 19, 27, 55}

	for _, d := range data {
		bst.Insert(d)
	}

	requires.Equal(4, bst.Height())
}
