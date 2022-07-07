package quick

import "testing"

func TestQuickSort(t *testing.T) {
	in := []int{8, 10, 2, 3, 6, 1, 5}
	quickSort(in, 0, len(in)-1)
}
