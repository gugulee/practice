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
