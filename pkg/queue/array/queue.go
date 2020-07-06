package array

// Queue is queue of array
type Queue struct {
	data []interface{}

	// the head of the queue
	head int

	// the tail of the queue
	tail int
}

// New return a new queue
func New(length int) *Queue {
	return &Queue{
		data: make([]interface{}, length),
		head: 0,
		tail: 0,
	}
}

// IsEmpty return true if the queue is empty, else return false
func (q *Queue) IsEmpty() bool {
	return q.head == q.tail
}

// IsFull return true if the queue is full, else return false
func (q *Queue) IsFull() bool {
	return q.tail == len(q.data)
}

// Enqueue ...
func (q *Queue) Enqueue(value interface{}) {
	// queue is full, migrate data
	if q.IsFull() {
		// it present the queue is full and it can migrate data no longer
		if 0 == q.head {
			return
		}
		
		// migrate data
		for i := q.head; i < q.tail; i++ {
			q.data[i-q.head] = q.data[i]
		}

		// the new head and tail after migrating data
		q.tail = 0 + q.tail - q.head
		q.head = 0
	}
	q.data[q.tail] = value
	q.tail++
}

// Dequeue ...
func (q *Queue) Dequeue() interface{} {
	// queue is empty, return nil
	if q.IsEmpty() {
		return nil
	}
	ret := q.data[q.head]
	q.head++
	return ret
}

// Print print the queue
func (q *Queue) Print() []interface{} {
	if q.IsEmpty() {
		return nil
	}

	var out []interface{}
	for i := q.head; i < q.tail; i++ {
		out = append(out, q.data[i])
	}

	return out
}
