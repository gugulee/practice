package findsamestring

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	stringLength = 5
)

func Init() {
	rand.Seed(time.Now().Unix())
}

var letter = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

// generateString generate string randomly
func generateString(n int) (r []string) {
	r = make([]string, n)
	for i := 0; i < n; i++ {
		var url string
		for j := 0; j < stringLength; j++ {
			random := rand.Intn(26)
			url += letter[random]
		}
		r[i] = url
	}

	return
}

func TestFindSameString(t *testing.T) {
	stringOne := generateString(10000)
	stringTwo := generateString(10000)
	var sameString []string

	for _, s1 := range stringOne {
		for _, s2 := range stringTwo {
			if s1 == s2 {
				sameString = append(sameString, s1)
			}
		}
	}

	out := findSameString(stringOne, stringTwo)
	require.ElementsMatch(t, sameString, out)
}
