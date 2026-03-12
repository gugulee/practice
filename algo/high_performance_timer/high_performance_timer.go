package highperformancetimer

import (
	"fmt"
	"time"

	"github.com/practice/pkg/heap"
)

// HighPerformanceTimer is a high performance timer
type HighPerformanceTimer struct {
	timer *heap.MinHeap
	task  map[int64]string
}

// New ...
func New(in map[int64]string) *HighPerformanceTimer {
	minHeap := heap.NewMin(len(in) + 1)
	for ts := range in {
		minHeap.Insert(int(ts))
	}

	return &HighPerformanceTimer{
		timer: minHeap,
		task:  in,
	}
}

// highPerformanceTimer is a high performance timer which was implement by heap
func (h *HighPerformanceTimer) highPerformanceTimer() {
	for !h.timer.IsEmpty() {
		now := time.Now().Unix()
		min := h.timer.Min()
		h.timer.RemoveMaxOrMin()

		gap := min - int(now)

		time.Sleep(time.Duration(gap) * time.Second)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), h.task[int64(min)])
	}
}
