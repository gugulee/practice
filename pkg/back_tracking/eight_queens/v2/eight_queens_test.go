package v2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_printQueues(t *testing.T) {
	in := []int{0, 4, 7, 5, 2, 6, 1, 3}
	printQueues(in)
}

func Test_isOK(t *testing.T) {
	in := []int{0, 4, 7, 5, 2, 6, 1, 3}

	out := isOK(in, 0, 0)
	require.Equal(t, true, out)

	out = isOK(in, 1, 4)
	require.Equal(t, true, out)

	out = isOK(in, 2, 7)
	require.Equal(t, true, out)

	out = isOK(in, 3, 5)
	require.Equal(t, true, out)

	out = isOK(in, 4, 2)
	require.Equal(t, true, out)

	out = isOK(in, 5, 6)
	require.Equal(t, true, out)

	out = isOK(in, 6, 1)
	require.Equal(t, true, out)

	out = isOK(in, 7, 3)
	require.Equal(t, true, out)

	out = isOK(in, 2, 5)
	require.Equal(t, false, out)

	out = isOK(in, 2, 4)
	require.Equal(t, false, out)
}

func Test_eightQueens(t *testing.T) {
	in := make([]int, 8)
	eightQueens(in, 0)
	fmt.Println(count)
}
