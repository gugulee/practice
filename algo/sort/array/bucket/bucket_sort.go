package bucket

import (
	"github.com/practice/pkg/sort/array/quick"
)

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

func bucketSort(original []int) {
	length := len(original)
	max := getMax(original)

	bucket := make([][]int, length)

	for i := 0; i < length; i++ {
		idx := original[i] * (length - 1) / max
		bucket[idx] = append(bucket[idx], original[i])
	}

	position := 0

	for _, b := range bucket {
		bLen := len(b)
		if bLen == 0 {
			continue
		}

		quick.Sort(b)
		copy(original[position:], b)
		position += bLen
	}
}

func bucketSort1(original []int) {
	max := getMax(original)

	bucket := make([]int, max+1)

	for _, i := range original {
		bucket[i]++
	}

	position := 0
	for i := range bucket {
		for 0 != bucket[i] {
			original[position] = i
			position++
			bucket[i]--
		}
	}
}
