package lcs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lcs(t *testing.T) {
	a := []byte("mitcmu")
	b := []byte("mtacnu")
	out := lcs(a, b)
	require.Equal(t, 4, out)
}
