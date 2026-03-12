package lcs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lcs(t *testing.T) {
	a := []byte("mitcmu")
	b := []byte("mtacnu")
	lcs(a, b, 0, 0, 0)
	require.Equal(t, 4, maxLcs)
}
