package tools

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetMin(t *testing.T) {
	req := require.New(t)
	in := []int{}
	out := GetMinIdx(in)
	req.Equal(-1, out)

	in = []int{5, 3, 19, 20, 100, 1}
	out = GetMinIdx(in)
	req.Equal(5, out)

	in = []int{5, 3, 1, 20, 100, 10}
	out = GetMinIdx(in)
	req.Equal(2, out)

	in = []int{5, 3, -1, 20, 1, 10}
	out = GetMinIdx(in)
	req.Equal(4, out)

	in = []int{-1, -1, 4, -1, 1}
	out = GetMinIdx(in)
	req.Equal(4, out)

	in = []int{-1, -1, -1}
	out = GetMinIdx(in)
	req.Equal(-1, out)
}

func TestMapDeepCopy(t *testing.T) {
	a := map[int]int{100: 10, 10: 1, 1: 2, 2: 3, 4: 3}
	b := make(map[int]int)
	MapDeepCopy(b, a)
	fmt.Println(b)
	require.NotEqual(t, fmt.Sprintf("%p", a), fmt.Sprintf("%p", b))
	require.Exactly(t, a, b)
}

func TestMax(t *testing.T) {
	out := Max()
	require.Equal(t, -1, out)

	a := 3
	b := 4
	c := 6

	out = Max(a, b)
	require.Equal(t, 4, out)

	out = Max(a, b, c)
	require.Equal(t, 6, out)
}

func Test_Remove(t *testing.T) {
	a := []int{3, 5, 7, 10, 13}

	out := Remove(a, 0)
	fmt.Println(out)
}

func Test_Min(t *testing.T) {
	out := Min()
	require.Equal(t, -1, out)

	a := 3
	b := 4
	c := 2

	out = Min(a, b)
	require.Equal(t, 3, out)

	out = Min(a, b, c)
	require.Equal(t, 2, out)
}
