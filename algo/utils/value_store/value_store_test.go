package valuestore

import (
	"testing"

	"github.com/practice/pkg/constants"
	"github.com/stretchr/testify/require"
)

func Test_KV_Compare(t *testing.T) {
	var v1 Value = &KV{Key: 1, Value: 10}
	var v2 Value = &KV{Key: 2, Value: 5}
	var v3 Value = &KV{Key: 3, Value: 15}
	var v4 Value = &KV{Key: 4, Value: constants.MinInt}
	var v5 Value = &dummy{}

	require.Equal(t, 1, v1.Compare(v2))

	require.Equal(t, 0, v1.Compare(v1))

	require.Equal(t, -1, v1.Compare(v3))

	require.Equal(t, -1, v4.Compare(v1))

	require.Equal(t, 1, v1.Compare(v4))

	require.Equal(t, -100, v1.Compare(v5))
}

func Test_KV_String(t *testing.T) {
	var v1 Value = &KV{Key: 1, Value: 10}
	require.Equal(t, "key: 1, value: 10", v1.String())
}

func Test_Ivalue_Compare(t *testing.T) {
	var v1 Value = &Ivalue{10}
	var v2 Value = &Ivalue{5}
	var v3 Value = &Ivalue{15}
	var v4 Value = &KV{Key: 10, Value: 100}
	var v5 Value = &Ivalue{constants.MinInt}

	require.Equal(t, 1, v1.Compare(v2))

	require.Equal(t, 0, v1.Compare(v1))

	require.Equal(t, -1, v1.Compare(v3))

	require.Equal(t, -100, v1.Compare(v4))
	require.Equal(t, -1, v5.Compare(v1))
}

func Test_Ivalue_String(t *testing.T) {
	var v1 Value = &Ivalue{10}
	require.Equal(t, "10", v1.String())
}
