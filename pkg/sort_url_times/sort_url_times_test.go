package sorturltimes

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	urlLength = 8
	urlPrefix = "www"
	urlAffix  = "com"
)

var letter = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

// generateURL generate URL
func generateURL(n int) (r map[string]int) {
	rand.Seed(time.Now().Unix())
	r = make(map[string]int, n)
	for i := 0; i < n; i++ {
		times := rand.Intn(100)
		var url string
		for j := 0; j < urlLength; j++ {
			random := rand.Intn(26)
			url += letter[random]
		}
		url = urlPrefix + "." + url + "." + urlAffix
		r[url] = times
	}

	return
}

func TestSortURLTimes(t *testing.T) {
	urls := generateURL(10000)
	out := sortURLTimes(urls)
	var times []int
	for _, u := range out {
		t := urls[u]
		times = append(times, t)
	}

	for i := 0; i < len(times)-1; i++ {
		require.Equal(t, true, times[i] <= times[i+1])
	}
}
