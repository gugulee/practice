package median

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getMedianFromStatic(t *testing.T) {
	r := require.New(t)
	in := []int{10, 23, 200, 3, 13, 59, 90, 101}
	out := getMedianFromStatic(in)

	r.Equal(23, out)

	in = []int{10, 23, 200, 13, 59, 90, 101}
	out = getMedianFromStatic(in)

	r.Equal(59, out)

	in = []int{10}
	out = getMedianFromStatic(in)

	r.Equal(10, out)
}

func Test_DynamicMedia(t *testing.T) {
	r := require.New(t)
	in := []int{10, 23, 200, 3, 13, 59, 90}
	dm := New(in)

	media := dm.Median()
	r.Equal(23, media)
	r.Equal([]int{23, 13, 10, 3}, dm.maxHeap.Slice())
	r.Equal([]int{59, 90, 200}, dm.minHeap.Slice())

	dm.Insert(15)
	media = dm.Median()
	r.Equal(15, media)
	r.Equal([]int{15, 13, 10, 3}, dm.maxHeap.Slice())
	r.Equal([]int{23, 59, 200, 90}, dm.minHeap.Slice())

	dm.Insert(40)
	media = dm.Median()
	r.Equal(23, media)
	r.Equal([]int{23, 15, 10, 3, 13}, dm.maxHeap.Slice())
	r.Equal([]int{40, 59, 200, 90}, dm.minHeap.Slice())

	dm.Insert(30)
	media = dm.Median()
	r.Equal(23, media)
	r.Equal([]int{23, 15, 10, 3, 13}, dm.maxHeap.Slice())
	r.Equal([]int{30, 40, 200, 90, 59}, dm.minHeap.Slice())

	dm.Insert(1)
	media = dm.Median()
	r.Equal(23, media)
	r.Equal([]int{23, 15, 10, 3, 13, 1}, dm.maxHeap.Slice())
	r.Equal([]int{30, 40, 200, 90, 59}, dm.minHeap.Slice())

	dm.Insert(20)
	media = dm.Median()
	r.Equal(20, media)
	r.Equal([]int{20, 15, 10, 3, 13, 1}, dm.maxHeap.Slice())
	r.Equal([]int{23, 40, 30, 90, 59, 200}, dm.minHeap.Slice())
}
