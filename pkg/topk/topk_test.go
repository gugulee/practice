package topk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_topKOfStaticData(t *testing.T) {
	r := require.New(t)
	in := []int{10, 23, 200, 3, 13, 59, 90, 101}
	out := topKOfStaticData(in, 3)

	r.ElementsMatch([]int{200, 101, 90}, out)

	in = []int{}
	out = topKOfStaticData(in, 3)

	r.Nil(out)

	in = []int{2, 10}
	out = topKOfStaticData(in, 3)

	r.Nil(out)
}
