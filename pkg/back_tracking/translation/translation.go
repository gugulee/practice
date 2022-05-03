package translation

import (
	"github.com/practice/pkg/heap"
)

var minHeap *heap.MinHeap

// if data[][]=[
// [8 6 3],
// [6 5 4 2],
// [9 6]]
// find the top 3 in data,
// e.g.,
// top1=data[0][0]+data[1][0]+data[2][0]=23
// top2=data[0][0]+data[1][1]+data[2][0]=22
// top3=data[0][1]+data[1][0]+data[2][0]=21
func translation(data [][]int, topk, cur, sum int) {
	length := len(data)
	if cur == length {
		if minHeap.Used() >= topk {
			if minHeap.Contains(sum) {
				return
			}

			if sum > minHeap.Min() {
				minHeap.RemoveMaxOrMin()
			}
		}
		minHeap.Insert(sum)
		return
	}

	for i := range data[cur] {
		translation(data, topk, cur+1, data[cur][i]+sum)
	}
}
