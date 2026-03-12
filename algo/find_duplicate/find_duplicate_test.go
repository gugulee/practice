package findduplicate

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindScopeDuplicate(t *testing.T) {
	in := []int{9, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := 9

	out := FindScopeDuplicate(in, n)
	require.Equal(t, 9, out)
}

func TestFindDuplicate(t *testing.T) {
	var origin []int
	for i := 0; i <= 10000; i++ {
		origin = append(origin, i)
	}

	in := append(origin, 8394)

	out, err := FindDuplicate(in)
	require.NoError(t, err)
	require.Equal(t, 8394, out)
}
