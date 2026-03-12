package counting

func getMax(a []int) int {
	max := a[0]
	length := len(a)

	for i := 1; i < length; i++ {
		if max < a[i] {
			max = a[i]
		}
	}

	return max
}

func sum(a []int) {
	length := len(a)
	for i := 1; i < length; i++ {
		a[i] += a[i-1]
	}
}

func countingSort(original []int) {
	max := getMax(original)
	length := len(original)

	bucket := make([]int, max+1)

	for _, i := range original {
		bucket[i]++
	}

	sum(bucket)

	r := make([]int, length)

	for i := length - 1; i >= 0; i-- {
		element := original[i]
		position := bucket[element] - 1
		r[position] = element
		bucket[element]--
	}

	copy(original, r)
}

// Sort ...
func Sort(original []int) {
	countingSort(original)
}
