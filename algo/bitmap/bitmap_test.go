package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	out := New(63)
	require.Equal(t, 1, len(out.data))

	out = New(0)
	require.Equal(t, 1, len(out.data))

	out = New(64)
	require.Equal(t, 2, len(out.data))

	out = New(100)
	require.Equal(t, 2, len(out.data))

	out = New(128)
	require.Equal(t, 3, len(out.data))
}

func Test_Set(t *testing.T) {
	b := New(100)
	b.Set(5)
	require.Equal(t, uint64(32), b.data[0])
	b.Clean(5)

	b.Set(62)
	require.Equal(t, uint64(0x4000000000000000), b.data[0])
	b.Clean(62)

	b.Set(63)
	require.Equal(t, uint64(0x8000000000000000), b.data[0])

	b.Set(100)
	require.Equal(t, uint64(0x1000000000), b.data[1])
}

func Test_Get(t *testing.T) {
	b := New(100)
	for i := 0; i < 66; i++ {
		b.Set(i)
	}

	require.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), b.data[0])
	require.Equal(t, uint64(3), b.data[1])

	for i := 0; i < 66; i++ {
		require.Equal(t, true, b.Get(i))
	}
	require.Equal(t, false, b.Get(67))
}
func Test_Clean(t *testing.T) {
	b := New(100)
	b.data[0] = 2147483711
	b.Clean(5)
	require.Equal(t, uint64(2147483679), b.data[0])

	b.Clean(4)
	require.Equal(t, uint64(2147483663), b.data[0])

	b.Clean(3)
	require.Equal(t, uint64(2147483655), b.data[0])

	b.Clean(2)
	require.Equal(t, uint64(2147483651), b.data[0])

	b.Clean(1)
	require.Equal(t, uint64(2147483649), b.data[0])

	b.Clean(0)
	require.Equal(t, uint64(2147483648), b.data[0])

	b.Clean(31)
	require.Equal(t, uint64(0), b.data[0])
}
