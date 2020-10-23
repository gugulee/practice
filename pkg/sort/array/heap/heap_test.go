package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MaxHeap_Insert(t *testing.T) {
	r := require.New(t)
	h := NewMax(16)

	values := []int{33, 17, 21, 16, 13, 15, 9, 5, 6, 7, 8, 1, 2}
	for _, v := range values {
		h.Insert(v)
	}

	r.Equal(13, h.used)
	r.Equal([]int{0, 33, 17, 21, 16, 13, 15, 9, 5, 6, 7, 8, 1, 2, 0, 0}, h.data)

	h.Insert(22)
	r.Equal(14, h.used)
	r.Equal([]int{0, 33, 17, 22, 16, 13, 15, 21, 5, 6, 7, 8, 1, 2, 9, 0}, h.data)

	h.Insert(40)
	r.Equal(15, h.used)
	r.Equal([]int{0, 40, 17, 33, 16, 13, 15, 22, 5, 6, 7, 8, 1, 2, 9, 21}, h.data)

	h.Insert(50)
	r.Equal(15, h.used)
	r.Equal([]int{0, 40, 17, 33, 16, 13, 15, 22, 5, 6, 7, 8, 1, 2, 9, 21}, h.data)
}

func Test_MaxHeap_RemoveMaxOrMin(t *testing.T) {
	r := require.New(t)
	h := NewMax(16)

	values := []int{33, 27, 21, 16, 13, 15, 19, 5, 6, 7, 8, 1, 2, 12}
	for _, v := range values {
		h.Insert(v)
	}

	r.Equal(14, h.used)
	r.Equal([]int{0, 33, 27, 21, 16, 13, 15, 19, 5, 6, 7, 8, 1, 2, 12, 0}, h.data)

	h.RemoveMaxOrMin()

	r.Equal(13, h.used)
	r.Equal([]int{0, 27, 16, 21, 12, 13, 15, 19, 5, 6, 7, 8, 1, 2, 0, 0}, h.data)
}

func Test_MinHeap_Insert(t *testing.T) {
	r := require.New(t)
	h := NewMin(16)

	values := []int{7, 5, 8, 9, 10, 4, 6, 3}
	for _, v := range values {
		h.Insert(v)
	}

	r.Equal(8, h.used)
	r.Equal([]int{0, 3, 4, 5, 7, 10, 8, 6, 9, 0, 0, 0, 0, 0, 0, 0}, h.data)
}

func Test_MinHeap_RemoveMaxOrMin(t *testing.T) {
	r := require.New(t)
	h := NewMin(16)

	values := []int{7, 5, 8, 9, 10, 4, 6, 3}
	for _, v := range values {
		h.Insert(v)
	}

	r.Equal(8, h.used)
	r.Equal([]int{0, 3, 4, 5, 7, 10, 8, 6, 9, 0, 0, 0, 0, 0, 0, 0}, h.data)

	h.RemoveMaxOrMin()

	r.Equal(7, h.used)
	r.Equal([]int{0, 4, 7, 5, 9, 10, 8, 6, 0, 0, 0, 0, 0, 0, 0, 0}, h.data)
}

func Test_Slice(t *testing.T) {
	r := require.New(t)
	h := NewMin(16)

	values := []int{7, 5, 8, 9, 10, 4, 6, 3}
	for _, v := range values {
		h.Insert(v)
	}

	r.Equal([]int{3, 4, 5, 7, 10, 8, 6, 9}, h.Slice())
}

func Test_BuildMaxHeap(t *testing.T) {
	r := require.New(t)

	values := []int{0, 7, 5, 19, 8, 4, 1, 20, 13, 16}

	BuildMaxHeap(values)

	r.Equal([]int{0, 20, 16, 19, 13, 4, 1, 7, 5, 8}, values)
}

func Test_BuildMinHeap(t *testing.T) {
	r := require.New(t)

	values := []int{0, 7, 5, 8, 9, 10, 4, 6, 3}

	BuildMinHeap(values)

	r.Equal([]int{0, 3, 5, 4, 7, 10, 8, 6, 9}, values)
}

func TestSort(t *testing.T) {
	r := require.New(t)

	values := []int{0, 9, 3, 6, 1, 5}

	Sort(values)

	r.Equal([]int{0, 1, 3, 5, 6, 9}, values)
}
