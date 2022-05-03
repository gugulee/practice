package v2

import (
	"math"

	skipstring "github.com/practice/pkg/list/skip/v2/string"
)

const (
	// defaultCapacity is the length of the hash table
	defaultCapacity = 16

	// loadFactor represent the vacancy rate of the hash table, the larger the load factor, the less free space is
	// loadFactorThreshold is the threshold which need expand hash table
	loadFactorThreshold = 0.7
)

// DictHt is the hash table which store the words
type DictHt struct {
	// the list of hash table
	node []*skipstring.Skip

	// the size of the hash table
	capacity int

	// the used size of the hash table
	used int
}

// Dict ...
type Dict struct {
	// ht[1] will be used when rehash
	ht [2]*DictHt

	// the index of rehash
	rehashIndex int

	// rehash represent that rehash process is in progress
	rehash bool
}

// hash convert character to int,
// such as 'a' -> 0, 'b' -> 1 and 'z' -> 25
func hash(s string, capacity int) (r int) {
	for i, c := range s {
		r += int((c - 'a')) * int(math.Pow(26, float64(i)))
	}
	return (r ^ (r >> 16)) & (capacity - 1)
}

// NewDictHt ...
func NewDictHt(capacity int) *DictHt {
	switch {
	case 0 == capacity:
		capacity = defaultCapacity
	case 0 != capacity%16:
		capacity = ((capacity + 16) >> 4) * 16
	}

	node := make([]*skipstring.Skip, capacity)

	for i := range node {
		node[i] = skipstring.New(0)
	}

	return &DictHt{
		node:     node,
		capacity: capacity,
	}
}

// New ...
func New(capacity int) *Dict {
	return &Dict{
		ht:          [2]*DictHt{NewDictHt(capacity)},
		rehashIndex: -1,
		rehash:      false,
	}
}

// NodeIsEmpty return true if dh.node[x] is empty
func (dh *DictHt) NodeIsEmpty(n int) bool {
	return dh.node[n].IsEmpty()
}

// IsEmpty return true if hash table is empty
func (dh *DictHt) IsEmpty() bool {
	return 0 == dh.used
}

// Insert the word into the word hash table
func (dh *DictHt) Insert(word string) {
	hashValue := hash(word, dh.capacity)

	// if the skip list is empty, dh.used++
	if dh.node[hashValue].IsEmpty() {
		dh.used++
	}

	dh.node[hashValue].Insert(word)
}

// Search the word into the word hash table
func (dh *DictHt) Search(word string) bool {
	hashValue := hash(word, dh.capacity)

	return nil != dh.node[hashValue].Search(word)
}

// NeedExpand return true if loadFactor > loadFactorThreshold
func (d *Dict) NeedExpand() bool {
	return float64(d.ht[0].used)/float64(d.ht[0].capacity) > loadFactorThreshold
}

// MigrateData migrate data from ht[0] to ht[1]
func (d *Dict) MigrateData() {
	var migrateList *skipstring.Skip

	// find one list in ht[0] which not empty
	for i := d.rehashIndex + 1; i < d.ht[0].capacity; i++ {
		migrateList = d.ht[0].node[i]
		if !migrateList.IsEmpty() {
			d.rehashIndex = i
			break
		}
	}

	for _, word := range migrateList.AllElement() {
		d.ht[1].Insert(word)
	}
	d.ht[0].used--

	if d.ht[0].IsEmpty() {
		d.ht[0] = d.ht[1]
		d.ht[1] = nil
		d.rehashIndex = -1
		d.rehash = false
	}
}

// Insert the word into the word hash table
func (d *Dict) Insert(word string) {
	// !d.rehash represent no expanded process is in progress
	if !d.rehash && d.NeedExpand() {
		// the expanded size is twice
		d.ht[1] = NewDictHt(2 * d.ht[0].capacity)
		d.rehash = true
	}

	if d.rehash {
		d.ht[1].Insert(word)
		// migrate one data from ht[0] into ht[1]
		d.MigrateData()
	} else {
		d.ht[0].Insert(word)
	}
}

// Search search the the word hash table with the word
func (d *Dict) Search(word string) bool {
	// if rehash process is in progress, we find word in ht[1] first
	if d.rehash && d.ht[1].Search(word) {
		return true
	}

	return d.ht[0].Search(word)
}
