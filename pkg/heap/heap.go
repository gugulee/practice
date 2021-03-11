package heap

// Heap ...
type Heap interface {
	Insert(value int)
	RemoveMaxOrMin()
}

// MaxHeap ...
type MaxHeap struct {
	// we don't store data at data[0]
	data     []int
	capacity int
	// used is the max index which store data
	used int
}

// NewMax return a empty max heap
func NewMax(capacity int) *MaxHeap {
	// set default capacity
	if 0 == capacity {
		capacity = 16
	}

	return &MaxHeap{
		data:     make([]int, capacity),
		capacity: capacity,
		used:     0,
	}
}

// Insert insert into the heap with the data
func (h *MaxHeap) Insert(value int) {
	// return if the heap is full
	if h.capacity == h.used+1 {
		return
	}

	h.used++
	h.data[h.used] = value

	for i := h.used; i>>1 > 0 && h.data[i] > h.data[i>>1]; i >>= 1 {
		h.data[i], h.data[i>>1] = h.data[i>>1], h.data[i]
	}

}

// RemoveMaxOrMin delete the maximum node
func (h *MaxHeap) RemoveMaxOrMin() {
	// return if the heap is empty
	if h.used <= 0 {
		return
	}

	// move the last data to h.data[1] which is the first data
	h.data[1] = h.data[h.used]
	h.data[h.used] = 0
	h.used--

	MaxHeapify(h.data, h.used, 1)
}

// Max return the maximum
func (h *MaxHeap) Max() int {
	if h.IsEmpty() {
		return -1
	}
	return h.data[1]
}

// IsEmpty return true if the heap is empty
func (h *MaxHeap) IsEmpty() bool {
	return h.used <= 0
}

// IsFull return true if the heap is full
func (h *MaxHeap) IsFull() bool {
	return h.capacity <= h.used+1
}

// Slice return the slice
func (h *MaxHeap) Slice() (result []int) {
	result = make([]int, h.used)

	for i := 1; i <= h.used; i++ {
		result[i-1] = h.data[i]
	}
	return result
}

// Used return the used
func (h *MaxHeap) Used() int {
	return h.used
}

// MinHeap ...
type MinHeap struct {
	// we don't store data at data[0]
	data     []int
	capacity int
	// maxIdx is the max index which store data
	used int
}

// NewMin return a empty min heap
func NewMin(capacity int) *MinHeap {
	// set default capacity
	if 0 == capacity {
		capacity = 16
	}

	return &MinHeap{
		data:     make([]int, capacity),
		capacity: capacity,
		used:     0,
	}
}

// Insert insert into the heap with the data
func (h *MinHeap) Insert(value int) {
	// return if the heap is full
	if h.IsFull() {
		return
	}

	h.used++
	h.data[h.used] = value

	for i := h.used; i>>1 > 0 && h.data[i] < h.data[i>>1]; i >>= 1 {
		h.data[i], h.data[i>>1] = h.data[i>>1], h.data[i]
	}
}

// RemoveMaxOrMin delete the maximum node
func (h *MinHeap) RemoveMaxOrMin() {
	// return if the heap is empty
	if h.used <= 0 {
		return
	}

	// move the last data to h.data[1] which is the first data
	h.data[1] = h.data[h.used]
	h.data[h.used] = 0
	h.used--

	MinHeapify(h.data, h.used, 1)
}

// Contains return true if the heap contains value
func (h *MinHeap) Contains(value int) bool {
	for i := 1; i <= h.used; i++ {
		if value == h.data[i] {
			return true
		}
	}

	return false
}

// Min return the minimum
func (h *MinHeap) Min() int {
	if h.IsEmpty() {
		return -1
	}
	return h.data[1]
}

// Slice return the slice
func (h *MinHeap) Slice() (result []int) {
	result = make([]int, h.used)

	for i := 1; i <= h.used; i++ {
		result[i-1] = h.data[i]
	}
	return result
}

// IsEmpty return true if the heap is empty
func (h *MinHeap) IsEmpty() bool {
	return h.used <= 0
}

// IsFull return true if the heap is full
func (h *MinHeap) IsFull() bool {
	return h.capacity <= h.used+1
}

// Used return the used
func (h *MinHeap) Used() int {
	return h.used
}

// BuildMaxHeap build max heap
func BuildMaxHeap(a []int) {
	maxIdx := len(a) - 1
	for i := maxIdx / 2; 1 <= i; i-- {
		MaxHeapify(a, maxIdx, i)
	}
}

// BuildMinHeap build min heap
func BuildMinHeap(a []int) {
	maxIdx := len(a) - 1
	for i := maxIdx / 2; 1 <= i; i-- {
		MinHeapify(a, maxIdx, i)
	}
}

// MaxHeapify heapify a which the top is maximum
func MaxHeapify(a []int, used, i int) {
	for {
		max := i // record the max position
		if i*2 <= used && a[i] < a[i*2] {
			max = i * 2
		}

		if i*2+1 <= used && a[max] < a[i*2+1] {
			max = i*2 + 1
		}

		if max == i {
			break
		}

		a[i], a[max] = a[max], a[i]

		i = max
	}
}

// MinHeapify heapify a which the top is minimum
func MinHeapify(a []int, used, i int) {
	for {
		min := i // record the min position
		if i*2 <= used && a[i] > a[i*2] {
			min = i * 2
		}

		if i*2+1 <= used && a[min] > a[i*2+1] {
			min = i*2 + 1
		}

		if min == i {
			break
		}

		a[i], a[min] = a[min], a[i]

		i = min
	}
}
