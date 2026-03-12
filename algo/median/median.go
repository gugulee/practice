package median

import (
	"fmt"
	"sort"

	"github.com/practice/pkg/heap"
)

// getMedianFromStatic get the median from the static data
func getMedianFromStatic(data []int) int {
	length := len(data)
	medianIdx := (length + 1) / 2
	maxHeap := heap.NewMax(medianIdx + 1)

	for i := 0; i < medianIdx; i++ {
		maxHeap.Insert(data[i])
	}

	for i := medianIdx; i < length; i++ {
		if data[i] < maxHeap.Max() {
			maxHeap.RemoveMaxOrMin()
			maxHeap.Insert(data[i])
		}
	}

	return maxHeap.Max()
}

// DynamicMedia ...
type DynamicMedia struct {
	// store the first half of the data,
	// The difference between maxHeap and minHeap is 0 or 1, and the count of the maxHeap > minHeap
	maxHeap *heap.MaxHeap

	// store the second half of the data
	minHeap *heap.MinHeap
}

// New ...
func New(initialData []int) *DynamicMedia {
	// sort the data
	sort.Ints(initialData)

	// initialize heap
	max := heap.NewMax(0)
	min := heap.NewMin(0)

	length := len(initialData)
	medianIdx := (length + 1) / 2

	// insert the first half of the data in the maxHeap
	for i := 0; i < medianIdx; i++ {
		max.Insert(initialData[i])
	}

	// insert the second half of the data in the minHeap
	for i := medianIdx; i < length; i++ {
		min.Insert(initialData[i])
	}

	return &DynamicMedia{
		maxHeap: max,
		minHeap: min,
	}
}

// Insert insert the data into heap
func (dm *DynamicMedia) Insert(data int) error {
	/* insert data */
	if data <= dm.maxHeap.Max() {
		dm.maxHeap.Insert(data)
	} else {
		dm.minHeap.Insert(data)
	}

	/* migrate data if necessary */
	switch dm.maxHeap.Used() - dm.minHeap.Used() {
	case -1: // minHeap's count > maxheap's count, migrate the data of the minHeap to the maxHeap
		// delete the min from the minHeap
		min := dm.minHeap.Min()
		dm.minHeap.RemoveMaxOrMin()

		// insert the min of the minHeap into Maxheap
		dm.maxHeap.Insert(min)
	case 2: // maxHeap's count > minHeap's count, migrate the data of the maxHeap to the minHeap
		// delete the max from the maxHeap
		max := dm.maxHeap.Max()
		dm.maxHeap.RemoveMaxOrMin()

		// insert the max of the maxHeap into minHeap
		dm.minHeap.Insert(max)
	case 0, 1: // no more action
	default:
		return fmt.Errorf("the count of the maxHeap and minHeap is incorrect")
	}

	return nil
}

// Median get the median from the dynamic data
func (dm *DynamicMedia) Median() int {
	return dm.maxHeap.Max()
}
