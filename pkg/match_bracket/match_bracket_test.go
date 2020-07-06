package match_bracket

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExpressionEvaluation(t *testing.T) {
	out := MatchBracket("3+5)*8-6")
	require.Equal(t, false, out)


	out = MatchBracket("(3+5)*8-6")
	require.Equal(t, true, out)

	out = MatchBracket("[(3+5)*8]-6")
	require.Equal(t, true, out)

	out = MatchBracket("{[(3+5)*8]-6}+(18+2*9")
	require.Equal(t, false, out)

	out = MatchBracket("{[(3+5)*8]-6}+(18+2*9)")
	require.Equal(t, true, out)

	out = MatchBracket("{[({3+5)*8]-6}+(18+2*9)")
	require.Equal(t, false, out)
}
