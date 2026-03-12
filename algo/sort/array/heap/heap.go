package heap

import (
	"github.com/practice/pkg/heap"
)

// Sort is the heap sort
func Sort(a []int) {
	heap.BuildMaxHeap(a)
	maxIdx := len(a) - 1

	for i := maxIdx; i > 1; {
		a[1], a[i] = a[i], a[1]
		i--
		heap.MaxHeapify(a, i, 1)
	}
}
