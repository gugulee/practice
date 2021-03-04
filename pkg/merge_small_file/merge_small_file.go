package mergesmallfile

import (
	"github.com/practice/pkg/heap"
	"github.com/practice/pkg/tools"
)

// mergeSmallFile merge small file into large file
func mergeSmallFile(data [][]int) (result []int) {
	length := len(data)
	resultLength := 0
	count := 0
	for _, d := range data {
		resultLength += len(d)
	}
	result = make([]int, resultLength)

	tmp := make([]int, length)

	for i := 0; i < length; i++ {
		tmp[i] = data[i][0]
	}

	for {
		// get the minimum
		minIdx := tools.GetMinIdx(tmp)

		// all the data have handled
		if minIdx < 0 {
			break
		}

		// populate the minimum into result
		result[count] = tmp[minIdx]
		count++

		data[minIdx] = data[minIdx][1:]

		// fetch one data from minIdx if data[minIdx] is not empty
		if 0 != len(data[minIdx]) {
			tmp[minIdx] = data[minIdx][0]
		} else {
			tmp[minIdx] = -1
		}
	}
	return result
}

// mergeSmallFile1 merge small file into large file
func mergeSmallFile1(data [][]int) (result []int) {
	length := len(data)
	resultLength := 0
	emptyCount := 0
	count := 0
	// relation store the relation between data and array
	// key is the data and value is the index
	relation := make(map[int]int)

	for _, d := range data {
		resultLength += len(d)
	}
	result = make([]int, resultLength)

	// minHeap[0] does not store data, so the length of heap is length+1
	minHeap := heap.NewMin(length + 1)
	for i := 0; i < length; i++ {
		minHeap.Insert(data[i][0])
		relation[data[i][0]] = i
	}

	for emptyCount < length {
		// get the mimimum
		min := minHeap.Min()

		// get the minimum index
		minIdx, ok := relation[min]
		if !ok {
			return nil
		}
		delete(relation, min)

		// populate the minimum into result
		result[count] = min
		count++

		// delete the minimum
		minHeap.RemoveMaxOrMin()

		data[minIdx] = data[minIdx][1:]

		// fetch one data from minIdx if data[minIdx] is not empty
		if 0 != len(data[minIdx]) {
			minHeap.Insert(data[minIdx][0])
			relation[data[minIdx][0]] = minIdx
		} else {
			emptyCount++
		}
	}

	return result
}
