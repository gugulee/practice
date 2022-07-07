package selection

func selectionSort(original []int) {
	length := len(original)
	for i := 0; i < length; i++ {
		minIdx := i
		for j := i + 1; j < length; j++ {
			if original[minIdx] > original[j] {
				minIdx = j
			}
		}
		original[i], original[minIdx] = original[minIdx], original[i]
	}
}
