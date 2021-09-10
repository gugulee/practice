package algo

import (
	"fmt"
	"time"

	circlequeue "github.com/practice/pkg/queue/array/circle"
)

type RollingWindow struct {
	// the max request limits at interval(e.g. 1s, 60s, 3600s)
	limit int

	// the interval
	Interval int

	// a cycle queue
	queue *circlequeue.Queue
}

func New(limit int) *RollingWindow {
	return &RollingWindow{
		limit: limit,
		queue: circlequeue.New(limit + 1),
	}
}

func (rw *RollingWindow) Limit() {
	// enqueue directly
	if rw.queue.IsEmpty() {
		rw.queue.Enqueue(time.Now().UnixNano())
	}

	// dequeue expire element
	for i := rw.queue.First(); i != q.tail; {

	}
	head := rw.queue.First()
	tail := rw.queue.Last()
	fmt.Println(head)
	fmt.Println(tail)

}
