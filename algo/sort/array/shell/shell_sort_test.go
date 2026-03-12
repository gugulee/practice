package shell

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShellSort(t *testing.T) {
	in := []int{99, 5, 69, 33, 56, 13, 22, 55, 77, 48, 12, 88, 2, 69, 99}
	shellSort(in)
	require.Equal(t, []int{2, 5, 12, 13, 22, 33, 48, 55, 56, 69, 69, 77, 88, 99, 99}, in)
}

func TestShellSort1(t *testing.T) {
	in := []int{99, 5, 69, 33, 56, 13, 22, 55, 77, 48, 12, 88, 2, 69, 99}
	shellSort1(in)
	require.Equal(t, []int{2, 5, 12, 13, 22, 33, 48, 55, 56, 69, 69, 77, 88, 99, 99}, in)
}
