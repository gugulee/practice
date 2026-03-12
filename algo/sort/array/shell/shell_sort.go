package shell

func shellSort(original []int) {
	length := len(original)
	if length < 2 {
		return
	}

	for gap := length >> 1; gap > 0; gap >>= 1 {
		for i := gap; i < length; i++ {
			value := original[i]
			j := i - gap

			for ; j >= 0; j -= gap {
				if original[j] > value {
					original[j+gap] = original[j]
				} else {
					break
				}
			}
			original[j+gap] = value
		}

	}
}

func shellSort1(original []int) {
	length := len(original)
	if length < 2 {
		return
	}

	for gap := length >> 1; gap > 0; gap >>= 1 {
		for i := gap; i < length; i++ {
			for j := i; j >= gap && original[j] < original[j-gap]; j -= gap {
				original[j], original[j-gap] = original[j-gap], original[j]
			}
		}
	}
}
