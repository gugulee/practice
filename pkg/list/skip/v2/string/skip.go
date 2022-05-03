package string

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/practice/pkg/tools"
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
	next  []*Node
	value string

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
func NewNode(layer int, value string) *Node {
	return &Node{
		next:  make([]*Node, layer),
		value: value,
		layer: layer,
	}
}

// New return the skip list, only the original list in the skip list
func New(capacity int) *Skip {
	if 0 == capacity {
		capacity = defaultCapacity
	}

	// min is the minimum is the skip list
	return &Skip{
		head:     NewNode(capacity, ""),
		capacity: capacity,
	}
}

// Head return the head of the skip list
func (s *Skip) Head() *Node {
	return s.head
}

// IsEmpty return true if the skip is empty
func (s *Skip) IsEmpty() bool {
	return nil == s.head.next[0]
}

// Insert into the skip list with the value
func (s *Skip) Insert(value string) {
	// we will insert the value in the layer where layer < newNodeLayer
	newNodeLayer := randomLayer()
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
		for nil != updateNode.next[i] && updateNode.next[i].value < value {
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
func (s *Skip) Search(value string) *Node {
	p := s.head

	// find p which value <= p.value
	for i := s.height - 1; 0 <= i; i-- {
		for nil != p.next[i] && p.next[i].value <= value {
			p = p.next[i]
		}

		if nil != p && value == p.value {
			return p
		}
	}

	return nil
}

// Delete delete the node which held the value in the skip list
func (s *Skip) Delete(value string) {
	p := s.head

	// find the node which node.value < value < node.next.value
	for i := s.height - 1; 0 <= i; i-- {
		for nil != p.next[i] && p.next[i].value < value {
			p = p.next[i]
		}

		if nil != p.next[i] && value == p.next[i].value {
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

// AllElement return all element of the skip list
func (s *Skip) AllElement() (r []string) {
	if s.IsEmpty() {
		return nil
	}

	node := s.head
	for ; node != nil; node = node.next[0] {
		if 0 != len(node.value) {
			r = append(r, node.value)
		}
	}
	return r
}

// Print print the layer of the skip list
func (n *Node) Print(layer int) string {
	node := n
	var info []string

	for ; node != nil; node = node.next[layer] {
		v := node.value
		if 0 == len(node.value) {
			v = "head"
		}

		info = append(info, v)
	}

	return fmt.Sprintf("layer %d: %s", layer, strings.Join(info, "->"))
}

// Search search the list with the value forward,
// return the node which node.value < value < node.next.value
func (n *Node) Search(value string, layer int) *Node {
	node := n

	for nil != node && node.value < value {
		node = node.next[layer]
	}

	if nil != node && value == node.value {
		return node
	}
	return nil
}

func randomLayer() int {
	layer := 1
	for i := 0; i < defaultCapacity; i++ {
		if rand.Intn(10000)&1 == 1 {
			layer++
		}
	}

	return layer
}
