package rabinkarp

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

func Test_rkSearch(t *testing.T) {
	r := require.New(t)
	data := "abcdefghijklmnopqrstuvwxyz"
	sub := "opq"
	r.Equal(strings.Contains(data, sub), rkSearch(data, sub))

	for i := 0; i < 10000; i++ {
		data = generateTestData(100)
		sub = "aab"
		out := rkSearch(data, sub)
		r.Equal(strings.Contains(data, sub), out)
	}
}

func Test_rkSearchMatrixString(t *testing.T) {
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

	out := rkSearchMatrixString(main, pattern)
	r.True(out)

	// pattern = [][]string{
	// 	{"c", "a", "f"},
	// 	{"e", "f", "d"},
	// }

	// out = rkSearchMatrixString(main, pattern)
	// r.False(out)
}