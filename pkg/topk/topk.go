package topk

import "github.com/practice/pkg/heap"

// topkInStatic get the top k in data
func topKOfStaticData(data []int, k int) []int {
	length := len(data)

	if 0 == len(data) || length < k {
		return nil
	}

	minHeap := heap.NewMin(k + 1)

	for i := 0; i < k; i++ {
		minHeap.Insert(data[i])
	}

	for i := k; i < length; i++ {
		if minHeap.Min() < data[i] {
			minHeap.RemoveMaxOrMin()
			minHeap.Insert(data[i])
		}
	}

	return minHeap.Slice()
}
