package algo

import (
	"time"

	circlequeue "github.com/practice/pkg/queue/array/circle"
	"github.com/practice/rate_limit/config"
)

var _ Interface = &RollingWindow{}

type RollingWindow struct {
	// the max request limits at interval(e.g. 1s, 60s, 3600s)
	limit int

	// the interval
	interval int

	// a cycle queue
	queue *circlequeue.Queue
}

func NewRollingWindow(limit *config.Limit) *RollingWindow {
	return &RollingWindow{
		limit:    limit.Limit,
		queue:    circlequeue.New(limit.Limit + 1),
		interval: limit.Interval,
	}
}

func (rw *RollingWindow) Limit() bool {
	// milisecond
	now := time.Now().UnixNano() / 1e6

	// delete the expired entries from the queue,
	// the unit of the entries of the queue is milisecond and the unit of interval is second
	for rw.queue.HasNext() {
		cur := rw.queue.Head().(int64)

		if cur+int64(rw.interval*1e3) > now {
			break
		}
		rw.queue.Dequeue()
	}

	// no more entries insert into queue if the queue is full
	if rw.queue.IsFull() {
		return false
	}

	// insert into queue if it's not full
	rw.queue.Enqueue(now)
	return true
}
