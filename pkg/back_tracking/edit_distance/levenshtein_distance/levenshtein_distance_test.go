package ld

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_levenshteinDistance(t *testing.T) {
	a := []byte("cu")
	b := []byte("mu")
	// a := []byte("mitcmu")
	// b := []byte("mtacnu")
	levenshteinDistance(a, b, 0, 0, 0)
	require.Equal(t, 1, minimum)
}
