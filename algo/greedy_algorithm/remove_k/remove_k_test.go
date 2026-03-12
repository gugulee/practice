package removek

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeK(t *testing.T) {
	out := removeK("1234567890", 9)
	require.Equal(t, "0", out)

	out = removeK("1432219", 3)
	require.Equal(t, "1219", out)

	out = removeK("1234", 2)
	require.Equal(t, "12", out)

	out = removeK("10200", 1)
	require.Equal(t, "200", out)

	out = removeK("10", 2)
	require.Equal(t, "0", out)
}

func Test_removeKdigits(t *testing.T) {
	out := removeKdigits("1234567890", 9)
	require.Equal(t, "0", out)

	out = removeKdigits("1432219", 3)
	require.Equal(t, "1219", out)

	out = removeKdigits("1234", 2)
	require.Equal(t, "12", out)

	out = removeKdigits("10200", 1)
	require.Equal(t, "200", out)

	out = removeKdigits("10", 2)
	require.Equal(t, "0", out)
}
