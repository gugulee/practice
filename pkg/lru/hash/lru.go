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

	// the max length of list, it's also the capacity of the lru
	maxLength int

	// the length of list
	length int

	// the last node of the list, we will insert new node after the last node
	last *Entry

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

	return &LRU{
		capacity:  capacity,
		maxLength: maxLength,
		length:    0,
		entry:     entry,
	}
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

// Insert insert data into hash table
func (lru *LRU) Insert(data string) {
	hashValue := hash(data, lru.capacity)

	// insert node in the tail of lru.entry[hashValue],
	// and insert new node in the tail of the double list
	newNode := lru.entry[hashValue].InsertTail(data)
	lru.length++

	/* insert new node after the last node of the double list*/
	// no more action if the last node of the double list is empty,
	// otherwise, insert new node after the last node
	if nil == lru.last {
		// update the last node
		lru.last = newNode
		return
	}

	nextNode := lru.last.next
	newNode.next = nextNode
	lru.last.next = newNode

	if nil != nextNode {
		newNode.prev = newNode
	}
	newNode.prev = lru.last

	// update the last node
	lru.last = newNode
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

	if nil != prev {
		prev.next = next
	}

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

	return strings.Join(r, "->")
}

// Lru is the lru cache which was implemented be hash table
func (lru *LRU) Lru(data string) {

}
