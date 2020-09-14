package tools

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapDeepCopy(t *testing.T) {
	a := map[int]int{100: 10, 10: 1, 1: 2, 2: 3, 4: 3}
	b := make(map[int]int)
	MapDeepCopy(b, a)
	fmt.Println(b)
	require.NotEqual(t, fmt.Sprintf("%p", a), fmt.Sprintf("%p", b))
	require.Exactly(t, a, b)
}
