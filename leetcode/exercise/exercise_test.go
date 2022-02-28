package exercise

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_exercise(t *testing.T) {
	in := []int{-2,1,-3,4,-1,2,1,-5,4}
	out := maxSubArray(in)
	require.Equal(t, 6, out)

}
