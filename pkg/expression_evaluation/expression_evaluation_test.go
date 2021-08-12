package expression_evaluation

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOperatorPriority(t *testing.T) {
	out := OperatorPriority("+", "-")
	require.Equal(t, false, out)

	out = OperatorPriority("+", "*")
	require.Equal(t, true, out)

	out = OperatorPriority("/", "*")
	require.Equal(t, false, out)

	out = OperatorPriority("nil", "-")
	require.Equal(t, true, out)
}

func TestExpressionEvaluation(t *testing.T) {
	out, err := ExpressionEvaluation("3+5*8-6")
	require.NoError(t, err)
	require.Equal(t, 37, out)

	out, err = ExpressionEvaluation("3+5*8*6")
	require.NoError(t, err)
	require.Equal(t, 243, out)

	out, err = ExpressionEvaluation("3+5*8*6-10/2")
	require.NoError(t, err)
	require.Equal(t, 238, out)
}
