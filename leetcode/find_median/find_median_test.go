package findmedian

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	out := findMedianSortedArrays(nums1, nums2)
	require.Equal(t, 2.5, out)

	nums1 = []int{1, 3}
	nums2 = []int{2, 4, 5}
	out = findMedianSortedArrays(nums1, nums2)
	require.Equal(t, 3.0, out)

	nums1 = []int{1, 2}
	nums2 = []int{-1, 3}
	out = findMedianSortedArrays(nums1, nums2)
	require.Equal(t, 1.5, out)

	nums1 = []int{1, 3, 4, 9}
	nums2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	out = findMedianSortedArrays(nums1, nums2)
	require.Equal(t, 4.0, out)
}
