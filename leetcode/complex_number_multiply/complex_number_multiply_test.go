package cnm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_complexNumberMultiply(t *testing.T) {
	num1 := "1+1i"
	num2 := "1+1i"
	out := complexNumberMultiply(num1, num2)
	require.Equal(t, "0+2i", out)
}
