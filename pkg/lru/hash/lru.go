package hash

import (
	"math"
	"strings"
)

const (
	defaultCapacity = 16
	defaultLength   = 16
)

// Entry store the entry of hash table
type Entry struct {
	// next and prev belongs to the double list,
	// and hNext belongs to the list of hash table,
	// the double list is the lru cache,
	// and the list of hash table is the collision list
	next, prev, hNext *Entry
	data              string
}

// LRU ...
type LRU struct {
	// the capacity of hash table
	capacity int

	// the max length of double list, it's also the capacity of the lru
	maxLength int

	// the length of list
	length int

	// last is the last node of the double list, we will insert new node after the last node,
	// head is the head node of the double list, the head node does not store data,
	// that is head.next is the first node store data
	head, last *Entry

	entry []*Entry
}

func hash(s string, capacity int) (r int) {
	for i, c := range s {
		r += int((c - 'a')) * int(math.Pow(26, float64(i)))
	}
	return (r ^ (r >> 16)) & (capacity - 1)
}

func newEntry(data string) *Entry {
	return &Entry{
		data: data,
	}
}

// New ...
func New(capacity, maxLength int) *LRU {
	if 0 == capacity {
		capacity = defaultCapacity
	}

	if 0 == maxLength {
		maxLength = defaultLength
	}

	entry := make([]*Entry, capacity)

	for i := 0; i < capacity; i++ {
		entry[i] = &Entry{}
	}

	lru := &LRU{
		capacity:  capacity,
		maxLength: maxLength,
		length:    0,
		head:      &Entry{},
		entry:     entry,
	}

	// the last node is the head initially
	lru.last = lru.head
	return lru
}

// InsertTail insert the new node in the tail of the list,
// last is the last node of list
func (e *Entry) InsertTail(data string) *Entry {
	node := e
	for ; nil != node.hNext; node = node.hNext {
	}
	return node.InsertAfter(data)
}

// InsertAfter insert new node after node
func (e *Entry) InsertAfter(data string) *Entry {
	newNode := newEntry(data)

	// insert new node in the collision list
	newNode.hNext = e.hNext
	e.hNext = newNode

	return newNode
}

// Search search the collision list with data
func (e *Entry) Search(data string) *Entry {
	node := e
	for ; nil != node; node = node.hNext {
		if data == node.data {
			return node
		}
	}
	return nil
}

// Delete delete the node which held the data
func (e *Entry) Delete(data string) *Entry {
	prev := e
	node := e.hNext

	for ; nil != node; node = node.hNext {
		if data == node.data {
			break
		}
		prev = node
	}

	if nil == node {
		return nil
	}

	prev.hNext = node.hNext

	return node
}

// Print print the collision list
func (e *Entry) Print() string {
	var r []string
	node := e
	for ; nil != node; node = node.hNext {
		if 0 != len(node.data) {
			r = append(r, node.data)
		}
	}

	return strings.Join(r, "->")
}

// Capacity return the capacity of hash table in the LRU
func (lru *LRU) Capacity() int {
	return lru.capacity
}

// IsFull return true if lru is full
func (lru *LRU) IsFull() bool {
	return lru.maxLength <= lru.length
}

// Insert insert data into hash table
func (lru *LRU) Insert(data string) {
	// return if lru is full
	if lru.IsFull() {
		return
	}

	hashValue := hash(data, lru.capacity)

	// insert node in the tail of lru.entry[hashValue],
	// and insert new node in the tail of the double list
	newNode := lru.entry[hashValue].InsertTail(data)

	/* insert new node after the last node of the double list*/
	node := lru.last
	node.next = newNode
	newNode.prev = node

	// update the last node
	lru.last = newNode

	// length add 1
	lru.length++
}

// Search search the lru with data
func (lru *LRU) Search(data string) *Entry {
	hashValue := hash(data, lru.capacity)
	return lru.entry[hashValue].Search(data)
}

// Delete delete the node which held the data in lru
func (lru *LRU) Delete(data string) {
	hashValue := hash(data, lru.capacity)

	// delete node from hash collision list
	node := lru.entry[hashValue].Delete(data)

	if nil == node {
		return
	}

	// delete node from double list
	prev := node.prev
	next := node.next

	prev.next = next
	if nil != next {
		next.prev = prev
	}

	lru.length--

	// if the delete node is the last node, update the last node
	if node == lru.last {
		lru.last = prev
	}
}

// DeleteNode delete the node in the double list
func (lru *LRU) DeleteNode(node *Entry) {
	if nil == node {
		return
	}

	hashValue := hash(node.data, lru.capacity)

	// delete node from hash collision list
	lru.entry[hashValue].Delete(node.data)

	// delete node from double list
	prev := node.prev
	next := node.next

	prev.next = next
	if nil != next {
		next.prev = prev
	}

	lru.length--

	// if the delete node is the last node, update the last node
	if node == lru.last {
		lru.last = prev
	}
}

// ReverstPrint ...
func (lru *LRU) ReverstPrint() string {
	var r []string
	node := lru.last

	for ; nil != node; node = node.prev {
		if 0 != len(node.data) {
			r = append(r, node.data)
		}
	}

	return strings.Join(r, "<-")
}

// Lru is the lru cache which was implemented be hash table
func (lru *LRU) Lru(data string) {
	switch lru.Search(data) {
	case nil: // no found
		// delete the node held the oldest data(the first node in the double list) when lru is full
		if lru.IsFull() {
			lru.DeleteNode(lru.head.next)
		}

	default: // found
		// delete node
		lru.Delete(data)
	}

	// insert node after the last node
	lru.Insert(data)
}
