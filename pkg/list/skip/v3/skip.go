package v3

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	// "github.com/practice/pkg/constants"
	"github.com/practice/pkg/tools"
	vs "github.com/practice/pkg/utils/value_store"
)

const (
	defaultCapacity = 16
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Node is the node of the skip list
type Node struct {
	// next it the next node of node.next[x],
	// such as node.next[0] is the next node where at the height 0
	next []*Node

	// value store value
	value vs.Value

	// layer is the layer where the node dewll in the skip list.
	// when layer is 2, node.next[2], node.next[1] and node.next[0] must exist.
	// for example, when the skip list is "1->3->5->7->9" and layer is 2 and node is node.1(it represent node.value=1),
	// node.next[0]=node.3, node.next[1]=node.5 and node.next[2]=node.9
	layer int
}

// Skip is the skip list and Skip.list[height].head is the head of the list
type Skip struct {
	// head is the head of the skip list
	head *Node

	// height is the highest layer where held data
	height int

	// capacity it the capacity of the skip list and height <= capacity
	capacity int
}

// NewNode return the node
// func NewNode(layer int, insert func(vs ValueStore) error) *Node {
func NewNode(layer int, value vs.Value) *Node {
	node := &Node{
		next:  make([]*Node, layer),
		value: value,
		layer: layer,
	}

	return node
}

// New return the skip list, only the original list in the skip list
func New(capacity int, value vs.Value) *Skip {
	if 0 == capacity {
		capacity = defaultCapacity
	}

	// min is the minimum is the skip list
	return &Skip{
		capacity: defaultCapacity,
		head:     NewNode(capacity, value),
	}
}

// Insert into the skip list with the value
func (s *Skip) Insert(value vs.Value) {
	// we will insert the value in the layer where layer < newNodeLayer
	newNodeLayer := randomLayer(s.capacity)
	// avoid the height increasing too fast
	if newNodeLayer > s.height {
		s.height++
		newNodeLayer = s.height
	}

	newNode := NewNode(newNodeLayer, value)

	// updateNodes store the node before new node
	updateNodes := make([]*Node, newNodeLayer)
	for i := 0; i < newNodeLayer; i++ {
		updateNodes[i] = s.head
	}

	updateNode := s.head
	// find the update node which updateNode.value < value < updateNode.next.value
	for i := s.height - 1; 0 <= i; i-- {
		for nil != updateNode.next[i] && -1 == updateNode.next[i].value.Compare(value) {
			updateNode = updateNode.next[i]
		}

		// we only store the update node where layer < newNodeLayer
		if i < newNodeLayer {
			updateNodes[i] = updateNode
		}
	}

	for i := 0; i < newNodeLayer; i++ {
		newNode.next[i] = updateNodes[i].next[i]
		updateNodes[i].next[i] = newNode
	}
}

// Search search the node which held the value in the skip list
func (s *Skip) Search(value vs.Value) *Node {
	p := s.head

	// find p which value <= p.next[i].value
	for i := s.height - 1; 0 <= i; i-- {
		for nil != p.next[i] && -1 == p.next[i].value.Compare(value) {
			p = p.next[i]
		}

		if nil != p.next[i] && 0 == p.next[i].value.Compare(value) {
			return p.next[i]
		}
	}

	return nil
}

// SearchRange search the skip list, return value(start.value <= value <= end.value)
func (s *Skip) SearchRange(start, end vs.Value) (r []vs.Value) {
	p := s.head

	// find p which start.value <= p.next[0].value
	for i := s.height - 1; 0 <= i; i-- {
		for nil != p.next[i] && -1 == p.next[i].value.CompareValue(start) {
			p = p.next[i]
		}
	}

	// find p which p.next[0].value <= end.value
	for nil != p.next[0] && p.next[0].value.CompareValue(end) <= 0 {
		r = append(r, p.next[0].value)
		p = p.next[0]
	}

	return r
}

// SearchRank search the skip list, 
// return value(there are offsetEnd-offsetStart+1 nodes between begin+offsetStart and begin+offsetStart+offsetEnd),
// no same node in the skip list when offsetStart == 0
func (s *Skip) SearchRank(begin vs.Value, offsetStart, offsetEnd int) (r []vs.Value) {
	p := s.head

	// find p which start.value <= p.next[0].value
	for i := s.height - 1; 0 <= i; i-- {
		for nil != p.next[i] && -1 == p.next[i].value.CompareValue(begin) {
			p = p.next[i]
		}
	}

	// move forward by 'start' nodes
	for nil != p.next[0] && offsetStart >0{
		p = p.next[0]
		offsetStart--
	}

	for nil != p.next[0] && offsetEnd > 0 {
		r = append(r, p.next[0].value)
		p = p.next[0]
		offsetEnd--
	}

	return r
}

// Delete delete the node which held the value in the skip list
func (s *Skip) Delete(value vs.Value) {
	p := s.head

	// find the node which node.value < value < node.next.value
	for i := s.height - 1; 0 <= i; i-- {
		for nil != p.next[i] && -1 == p.next[i].value.Compare(value) {
			p = p.next[i]
		}

		if nil != p.next[i] && 0 == p.next[i].value.Compare(value) {
			p.next[i] = p.next[i].next[i]
		}
	}
}

// Print print the skip list(< height)
func (s *Skip) Print() string {
	r := make([]string, s.height)

	for i := s.height - 1; 0 <= i; i-- {
		r[i] = s.head.Print(i)
	}

	tools.Reverse(r)
	return strings.Join(r, "\n")
}

// Print print the layer of the skip list
func (n *Node) Print(layer int) string {
	if nil == n {
		return ""
	}
	node := n.next[layer]
	var info = []string{"head"}

	for ; node != nil; node = node.next[layer] {
		v := node.value.String()
		info = append(info, v)
	}

	return fmt.Sprintf("layer %d: %s", layer, strings.Join(info, "->"))
}

// Search search the list with the value forward,
// return the node which node.value < value < node.next.value
func (n *Node) Search(value vs.Value, layer int) *Node {
	node := n

	for nil != node && -1 == node.value.Compare(value) {
		node = node.next[layer]
	}

	if nil != node && 0 == node.value.Compare(value) {
		return node
	}
	return nil
}

func randomLayer(capacity int) int {
	layer := 1
	for i := 0; i < capacity; i++ {
		if rand.Intn(10000)&1 == 1 {
			layer++
		}
	}

	return layer
}
