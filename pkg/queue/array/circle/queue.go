package circle

// Queue is queue of array
type Queue struct {
	data []interface{}

	// the head of the queue
	head int

	// the tail of the queue
	tail int

	// the length of the queue
	capacity int
}

// New return a new queue
func New(length int) *Queue {
	return &Queue{
		data:     make([]interface{}, length),
		head:     0,
		tail:     0,
		capacity: length,
	}
}

// IsEmpty return true if the queue is empty, else return false
func (q *Queue) IsEmpty() bool {
	return q.head == q.tail
}

// IsFull return true if the queue is full, else return false
func (q *Queue) IsFull() bool {
	return q.head == (q.tail+1)%q.capacity
}

// Enqueue ...
func (q *Queue) Enqueue(value interface{}) {
	if q.IsFull() {
		return
	}

	q.data[q.tail] = value

	// ensure tail not overflow
	q.tail = (q.tail + 1) % q.capacity
}

// Dequeue ...
func (q *Queue) Dequeue() interface{} {
	// queue is empty, return nil
	if q.IsEmpty() {
		return nil
	}

	ret := q.data[q.head]

	// ensure head not overflow
	q.head = (q.head + 1) % q.capacity
	return ret
}

// Print print the queue
func (q *Queue) Print() []interface{} {
	if q.IsEmpty() {
		return nil
	}

	var out []interface{}
	for i := q.head; i != q.tail; {
		out = append(out, q.data[i])
		i = (i + 1) % q.capacity
	}

	return out
}

// HasNext returns whether another next element exists.
func (q *Queue) HasNext() bool {
	return !q.IsEmpty()
}

// Head returns the head of the queue
func (q *Queue) Head() interface{} {
	return q.data[q.head]
}
