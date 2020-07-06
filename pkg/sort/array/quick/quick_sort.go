package quick

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

func quickSort(original []int, start, end int) {
	if end <= start {
		return
	}

	pivot := partition(original, start, end)
	quickSort(original, start, pivot-1)
	quickSort(original, pivot+1, end)
}
