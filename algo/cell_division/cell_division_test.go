package celldivision

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCellDivision(t *testing.T) {
	out := cellDivision(-10)
	require.Equal(t, -1, out)

	out = cellDivision(3)
	require.Equal(t, 7, out)

	out = cellDivision(4)
	require.Equal(t, 13, out)

	out = cellDivision(5)
	require.Equal(t, 24, out)
}
