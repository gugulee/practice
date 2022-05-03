package longestpalindrome

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestPalindromet(t *testing.T) {
	s := "babad"
	out := longestPalindrome(s)
	require.Equal(t, "bab", out)
}
