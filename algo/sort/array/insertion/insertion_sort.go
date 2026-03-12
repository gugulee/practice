package insertion

func insertionSort(original []int) {
	length := len(original)
	if length < 2 {
		return
	}

	for i := 1; i < length; i++ {
		value := original[i]
		j := i - 1

		for ; j >= 0; j-- {
			if original[j] > value {
				original[j+1] = original[j]
			} else {
				break
			}
		}
		original[j+1] = value
	}
}

func insertionSortReview(original []int) {
	n := len(original)

	for i := 1; i < n; i++ {
		j := i - 1
		cur := original[i]

		for ; j >= 0; j-- {
			if original[j] > cur {
				original[j+1] = original[j]
			} else {
				break
			}
		}

		original[j+1] = cur
	}
}
