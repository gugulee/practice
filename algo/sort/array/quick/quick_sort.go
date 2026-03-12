package quick

// Sort ...
func Sort(original []int) {
	quickSort(original, 0, len(original)-1)
}

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

	p := partition(original, start, end)
	quickSort(original, start, p-1)
	quickSort(original, p+1, end)
}

func quickSortReview(original []int, start, end int) {
	if end <= start {
		return
	}

	p := partitionReview(original, start, end)
	quickSortReview(original, start, p-1)
	quickSortReview(original, p+1, end)
}

func partitionReview(a []int, start, end int) int {
	piovt := a[end]

	i := start
	for j := start; j < end; j++ {
		if a[j] < piovt {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}

	a[i], a[end] = a[end], a[i]

	return i
}
