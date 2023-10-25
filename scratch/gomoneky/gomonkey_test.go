package gomonkey

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/require"
)

func TestCompute(t *testing.T) {
	var c *Computer
	patches := gomonkey.ApplyMethod(reflect.TypeOf(c), "NetworkCompute", func(_ *Computer, a, b int) (int, error) {
		return 10, nil
	})

	defer patches.Reset()

	cp := &Computer{}
	sum, err := cp.Compute(1, 1)
	require.NoError(t, err)
	require.Equal(t, 10, sum)
}
