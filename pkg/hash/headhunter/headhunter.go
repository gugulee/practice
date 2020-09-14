package headhunter

import (
	"fmt"

	"github.com/practice/pkg/constants"
	skipv3 "github.com/practice/pkg/list/skip/v3"
	valuestore "github.com/practice/pkg/utils/value_store"
)

const (
	defaultCapacity = 16
)

var _ valuestore.Value = &Node{}

// Node ...
type Node struct {
	id    int
	score int
}

// Compare compare id and score,
// return -100 if *Node does not implement ValueStore
func (n *Node) Compare(value valuestore.Value) int {
	n1, ok := value.(*Node)

	if !ok {
		return -100
	}

	if n.score > n1.score {
		return 1
	} else if n.score < n1.score {
		return -1
	}

	// compare id when score are equal
	if n.id > n1.id {
		return 1
	} else if n.id < n1.id {
		return -1
	}

	// id and score are equal
	return 0
}

// CompareValue compare score only,
// return -100 if *Node does not implement ValueStore
func (n *Node) CompareValue(value valuestore.Value) int {
	n1, ok := value.(*Node)

	if !ok {
		return -100
	}

	if n.score > n1.score {
		return 1
	} else if n.score < n1.score {
		return -1
	}

	return 0
}

func (n *Node) String() string {
	return fmt.Sprintf("%d/%d", n.id, n.score)
}

// Hash ...
type Hash struct {
	node []*Node

	// capacity it the capacity of the hash table
	capacity int

	// skip is the skip list
	skip *skipv3.Skip
}

func newNode(id, score int) *Node {
	return &Node{
		id:    id,
		score: score,
	}
}

// New ...
func New(capacity int) *Hash {
	if 0 == capacity {
		capacity = defaultCapacity
	}

	return &Hash{
		node:     make([]*Node, capacity),
		capacity: capacity,
		skip:     skipv3.New(0, newNode(constants.MinInt, constants.MinInt)),
	}
}

// Update insert and update value
func (h *Hash) Update(value valuestore.Value) {
	newValue, ok := value.(*Node)
	if !ok {
		return
	}

	// insert value in the hash table if the value does not exist in it
	if nil == h.node[newValue.id] {
		h.node[newValue.id] = newNode(newValue.id, newValue.score)
	} else {
		// update the value if the value exist
		oldValue := h.node[newValue.id]

		// no need update when value is equal
		if oldValue.score == newValue.score {
			return
		}

		// delete value from the skip list
		h.skip.Delete(oldValue)

		h.node[newValue.id].score = newValue.score
	}

	// insert new value in the skip list
	h.skip.Insert(value)
}

// Search search the hash table with the value and return the score
func (h *Hash) Search(id int) int {
	if nil == h.node[id] {
		return -100
	}

	return h.node[id].score
}

// SearchRange search the hash table, return a list of id in a score range(start <= score <= end)
func (h *Hash) SearchRange(start, end int) (ids []int) {
	// we only care score, so id=0
	r := h.skip.SearchRange(newNode(0, start), newNode(0, end))
	ids = make([]int, len(r))

	for i := range r {
		n := r[i].(*Node)
		ids[i] = n.id
	}

	return ids
}

// Delete delete the node that held data
func (h *Hash) Delete(id int) {
	if nil == h.node[id] {
		return
	}

	// delete value from the skip list
	h.skip.Delete(h.node[id])
	h.node[id] = nil
}
