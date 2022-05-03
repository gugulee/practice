package bucket

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBucketSort(t *testing.T) {
	in := []int{1, 6, 3, 5, 8, 6, 4}
	bucketSort(in)
	require.Equal(t, []int{1, 3, 4, 5, 6, 6, 8}, in)
}


func TestBucketSort1(t *testing.T) {
	in := []int{1, 6, 3, 5, 8, 6, 4}
	bucketSort1(in)
	require.Equal(t, []int{1, 3, 4, 5, 6, 6, 8}, in)
}