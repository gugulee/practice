package findkthlargest

// a[:i] is the processed area where is less then pivot,
// a[j:] is the unprocessed area
func partition(a []int, start, end int) int {
	// we select the end as the pivot
	pivot := a[end]

	i := start

	for j := start; j < end; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[end] = a[end], a[i]

	return i
}

// findKthLargest find the Kth largest in original
func findKthLargest(original []int, start, end, k int) int {
	p := partition(original, start, end)
	var result int
	switch {
	// original[p] is the element that we want to find
	case p+1 == k:
		result = original[p]
	// continue to find in original[p+1,end]
	case p+1 < k:
		result = findKthLargest(original, p+1, end, k)
	// continue to find in original[start,p-1]
	case p+1 > k:
		result = findKthLargest(original, start, p-1, k)
	}

	return result
}
