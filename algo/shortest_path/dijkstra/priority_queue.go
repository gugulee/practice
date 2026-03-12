package dijkstra

type vertex struct {
	id       int
	distance int
}

// PriorityQueue is a mimimum heap
type PriorityQueue struct {
	// we don't store data at data[0]
	data  []*vertex
	count int
	used  int
}

// NewPriorityQueue ...
func NewPriorityQueue(count int) *PriorityQueue {
	return &PriorityQueue{
		data:  make([]*vertex, count+1),
		count: count,
		used:  0,
	}
}

// IsFull return true if the heap is full
func (p *PriorityQueue) IsFull() bool {
	return p.count <= p.used
}

// IsEmpty return true if the heap is empty
func (p *PriorityQueue) IsEmpty() bool {
	return p.used <= 0
}

func (p *PriorityQueue) pull() *vertex {
	// return if the heap is empty
	if p.IsEmpty() {
		return nil
	}

	// move the last data to h.data[1] which is the first data
	tmp := p.data[1]
	p.data[1] = p.data[p.used]
	p.data[p.used] = nil
	p.used--

	HeapifyUpToDown(p.data, p.used, 1)
	return tmp
}

func (p *PriorityQueue) add(id, distance int) {
	// return if the heap is full
	if p.IsFull() {
		return
	}

	p.used++
	p.data[p.used] = &vertex{id, distance}

	HeapifyDownToUp(p.data, p.used)
	// for i := p.used; i>>1 > 0 && p.data[i].distance < p.data[i>>1].distance; i >>= 1 {
	// 	p.data[i], p.data[i>>1] = p.data[i>>1], p.data[i]
	// }
}

func (p *PriorityQueue) update(id, distance int) {
	i := 1
	for ; i <= p.used; i++ {
		if p.data[i].id == id {
			p.data[i].distance = distance
			break
		}
	}

	// id does not exist in queue
	if i > p.used {
		return
	}

	HeapifyDownToUp(p.data, i)
	HeapifyUpToDown(p.data, p.used, i)
}

// slice return the slice
func (p *PriorityQueue) slice() (result []vertex) {
	result = make([]vertex, p.used)

	for i := 1; i <= p.used; i++ {
		result[i-1] = *p.data[i]
	}
	return result
}

// HeapifyUpToDown ...
func HeapifyUpToDown(a []*vertex, used, i int) {
	for {
		min := i // record the min position
		if i*2 <= used && a[i].distance > a[i*2].distance {
			min = i * 2
		}

		if i*2+1 <= used && a[min].distance > a[i*2+1].distance {
			min = i*2 + 1
		}

		if min == i {
			break
		}

		a[i], a[min] = a[min], a[i]

		i = min
	}
}

// HeapifyDownToUp ...
func HeapifyDownToUp(a []*vertex, used int) {
	for i := used; i>>1 > 0 && a[i].distance < a[i>>1].distance; i >>= 1 {
		a[i], a[i>>1] = a[i>>1], a[i]
	}
}
