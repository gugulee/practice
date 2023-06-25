package bubble

func bubbleSort(original []int) {
	length := len(original)
	for j := 0; j < length; j++ {
		flag := false
		for i := 0; i < length-j-1; i++ {
			if original[i] > original[i+1] {
				original[i], original[i+1] = original[i+1], original[i]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func bubbleSortReview(original []int) {
	n := len(original)
	for i := 0; i < n; i++ {
		swap := false
		for j := 0; j < n-i-1; j++ {
			if original[j] > original[j+1] {
				original[j], original[j+1] = original[j+1], original[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}
