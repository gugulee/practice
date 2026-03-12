package bruteforce

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var letter map[int]string = map[int]string{
	0: "a", 1: "b", 2: "c", 3: "d", 4: "e", 5: "f", 6: "g", 7: "h", 8: "i", 9: "j",
	10: "k", 11: "l", 12: "m", 13: "n", 14: "o", 15: "p", 16: "q", 17: "r", 18: "s", 19: "t",
	20: "u", 21: "v", 22: "w", 23: "x", 24: "y", 25: "z",
}

func generateTestData(n int) (result string) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < n; i++ {
		t := rand.Intn(5)
		result += letter[t]
	}

	return
}
func Test_bfSearch(t *testing.T) {
	r := require.New(t)
	data := "aaaaaab"
	sub := "aab"
	r.Equal(strings.Contains(data, sub), bfSearch(data, sub))

	for i := 0; i < 1000; i++ {
		data = generateTestData(100)
		sub = "aab"
		out := bfSearch(data, sub)
		r.Equal(strings.Contains(data, sub), out)
	}
}

func Test_bfSearch1(t *testing.T) {
	r := require.New(t)
	data := "aaaaaab"
	sub := "aab"
	r.Equal(strings.Contains(data, sub), bfSearch1(data, sub))

	for i := 0; i < 1000; i++ {
		data = generateTestData(100)
		sub = "aab"
		out := bfSearch1(data, sub)
		r.Equal(strings.Contains(data, sub), out)
	}
}

func Test_bfSearchMatrixString(t *testing.T) {
	r := require.New(t)
	main := [][]string{
		{"d", "a", "b", "c", "d"},
		{"e", "f", "a", "d", "g"},
		{"c", "c", "a", "f", "h"},
		{"d", "e", "f", "c", "i"},
	}

	pattern := [][]string{
		{"c", "a", "f"},
		{"e", "f", "c"},
	}

	out := bfSearchMatrixString(main, pattern)
	r.True(out)

	pattern = [][]string{
		{"c", "a", "f"},
		{"e", "f", "d"},
	}

	out = bfSearchMatrixString(main, pattern)
	r.False(out)
}

func Test_truncate(t *testing.T) {
	r := require.New(t)
	main := [][]string{
		{"d", "a", "b", "c", "d"},
		{"e", "f", "a", "d", "g"},
		{"c", "c", "a", "f", "h"},
		{"d", "e", "f", "c", "i"},
	}

	out := truncate(main, 0, 1, 0, 2)
	r.Equal([][]string{
		{"d", "a", "b"},
		{"e", "f", "a"},
	}, out)

	out = truncate(main, 1, 1, 0, 2)
	r.Equal([][]string{
		{"e", "f", "a"},
	}, out)

	out = truncate(main, 2, 1, 0, 2)
	r.Nil(out)

	out = truncate(main, 1, 2, 3, 2)
	r.Nil(out)

	out = truncate(main, 1, 4, 0, 2)
	r.Nil(out)

	out = truncate(main, 1, 3, 0, 5)
	r.Nil(out)
}

func Test_isSliceEqual(t *testing.T) {
	r := require.New(t)

	a := []string{"d", "a", "b", "c", "d"}
	b := []string{"d", "a", "b", "c", "d"}
	r.True(isSliceEqual(a, b))

	a = []string{"d", "a", "b", "c", "d"}
	b = []string{"d", "c", "b", "c", "d"}
	r.False(isSliceEqual(a, b))

	a = []string{}
	b = []string{}
	r.False(isSliceEqual(a, b))

	a = []string{}
	b = []string{"d", "c", "b", "c", "d"}
	r.False(isSliceEqual(a, b))

	a = []string{"d", "a", "b", "c", "d"}
	b = []string{"d", "c", "b", "c"}
	r.False(isSliceEqual(a, b))
}

func Test_isEqual(t *testing.T) {
	r := require.New(t)

	a := [][]string{
		{"c", "a", "f"},
		{"e", "f", "c"},
	}
	b := [][]string{
		{"c", "a", "f"},
		{"e", "f", "c"},
	}
	r.True(isEqual(a, b))

	a = [][]string{
		{"c", "d", "f"},
		{"e", "f", "c"},
	}
	b = [][]string{
		{"c", "a", "f"},
		{"e", "f", "c"},
	}
	r.False(isEqual(a, b))

	a = [][]string{}
	b = [][]string{}
	r.False(isEqual(a, b))

	a = [][]string{}
	b = [][]string{
		{"c", "a", "f"},
		{"e", "f", "c"},
	}
	r.False(isEqual(a, b))

	a = [][]string{
		{"c", "d", "f"},
		{"e", "f", "c"},
	}
	b = [][]string{
		{"c", "a", "f"},
	}
	r.False(isEqual(a, b))

	a = [][]string{
		{"c", "d", "f"},
		{"e", "f", "c"},
	}
	b = [][]string{
		{"c", "a", "f"},
		{"e", "f"},
	}
	r.False(isEqual(a, b))
}
