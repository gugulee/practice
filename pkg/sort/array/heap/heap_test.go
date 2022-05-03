package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
)
func TestSort(t *testing.T) {
	r := require.New(t)

	values := []int{0, 9, 3, 6, 1, 5}

	Sort(values)

	r.Equal([]int{0, 1, 3, 5, 6, 9}, values)
}
