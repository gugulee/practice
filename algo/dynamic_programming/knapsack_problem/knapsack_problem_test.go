package kp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kp(t *testing.T) {
	items := []int{2, 2, 4, 6, 3}
	w := 9
	out := kp(items, w)
	require.Equal(t, 9, out)
}

func Test_kp1(t *testing.T) {
	items := []int{2, 2, 4, 6, 3}
	w := 9
	out := kp1(items, w)
	require.Equal(t, 9, out)
}

func Test_kp2(t *testing.T) {
	items := []int{2, 2, 4, 6, 3}
	value := []int{3, 4, 8, 9, 6}
	w := 9
	maxW, maxV := kp2(items, value, w)
	require.Equal(t, 9, maxW)
	require.Equal(t, 18, maxV)
}

func Test_kp3(t *testing.T) {
	items := []int{2, 2, 4, 6, 3}
	value := []int{3, 4, 8, 9, 6}
	w := 9
	maxW, maxV := kp3(items, value, w)
	require.Equal(t, 9, maxW)
	require.Equal(t, 18, maxV)
}
