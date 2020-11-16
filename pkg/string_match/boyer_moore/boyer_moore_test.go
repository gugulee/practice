package boyermoore

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

func Test_generateBC(t *testing.T) {
	r := require.New(t)
	b := "abda"
	bc := generateBC(b)
	r.Equal(3, bc[97])
	r.Equal(1, bc[98])
	r.Equal(2, bc[100])
}

func Test_generateGS(t *testing.T) {
	r := require.New(t)
	b := "cabcab"
	prefix, suffix := generateGS(b)

	r.ElementsMatch([]int{-1, 2, 1, 0, -1, -1}, suffix)
	r.ElementsMatch([]bool{false, false, false, true, false, false}, prefix)
}

func Test_bm(t *testing.T) {
	r := require.New(t)
	data := "abcbdabdc"
	sub := "abd"
	out := bmSearch(data, sub)
	r.Equal(strings.Contains(data, sub), out)

	data = "xabcabcab"
	sub = "cabcab"
	out = bmSearch(data, sub)
	r.Equal(strings.Contains(data, sub), out)

	for i := 0; i < 1000; i++ {
		data = generateTestData(100)
		sub = "aab"
		out := bmSearch(data, sub)
		r.Equal(strings.Contains(data, sub), out)
	}
}
