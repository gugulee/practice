package regex

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_regex(t *testing.T) {

	text := "bc"
	p := "b?c"
	rmatch(0, 0, p, text)
	require.True(t, match)
	match = false

	text = "bac"
	p = "b?c"
	rmatch(0, 0, p, text)
	require.True(t, match)
	match = false

	text = "bcaac"
	p = "b?c"
	rmatch(0, 0, p, text)
	require.False(t, match)
	match = false

	text = "baaac"
	p = "b*c"
	rmatch(0, 0, p, text)
	require.True(t, match)
	match = false

	text = "baaacad"
	p = "b*c?d"
	rmatch(0, 0, p, text)
	require.True(t, match)
	match = false

	text = "bc"
	p = "b*c"
	rmatch(0, 0, p, text)
	require.True(t, match)
	match = false
}
