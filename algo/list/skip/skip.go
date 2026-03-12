package skip

import (
	"fmt"
	"strings"

	"github.com/practice/pkg/tools"
)

const (
	min             = -10000
	defaultCapacity = 4
)

// Node is the node of the skip list
type Node struct {
	next     *Node
	previous *Node
	up       *Node
	down     *Node
	value    int
}

// List is the list of the skip list
type List struct {
	listHead *Node
	length   int
}

// Skip is the skip list and Skip.list[height].head is the head of the list
type Skip struct {
	// list[0] it the original list, list[1] is the first index and so on
	list []*List

	// height is the height of the skip list, height is 0 when only exists the original list
	height int

	// capacity it the capacity of the skip list, need scale when capacity <= height
	capacity int
}

// newNode return the node
func newNode(value int) *Node {
	return &Node{value: value}
}

// newList return the node
func newList(value int) *List {
	return &List{
		listHead: newNode(value),
		length:   0,
	}
}

// New return the skip list, only the original list in the skip list
func New(capacity int) *Skip {
	if 0 == capacity {
		capacity = defaultCapacity
	}

	list := make([]*List, capacity)

	for i := 0; i < capacity; i++ {
		list[i] = newList(min)
	}

	// connect list[i+1] and list[i]
	for i := capacity - 1; i > 0; i-- {
		high := list[i].listHead
		low := list[i-1].listHead
		high.down = low
		low.up = high
	}

	// min is the minimum is the skip list
	return &Skip{
		list:     list,
		height:   0,
		capacity: capacity,
	}
}

// scale scale the skip list when capacity <= height
func (s *Skip) scale() {
	newCapacity := s.capacity * 2

	list := make([]*List, newCapacity)

	copy(list, s.list)

	for i := s.capacity; i < newCapacity; i++ {
		list[i] = newList(min)
	}

	// connect list[i+1] and list[i]
	for i := newCapacity - 1; s.capacity <= i; i-- {
		high := list[i].listHead
		low := list[i-1].listHead
		high.down = low
		low.up = high
	}

	s.list = list
	s.capacity = newCapacity
}

// Head return the head of the skip list
func (s *Skip) Head() *Node {
	return s.list[s.height].listHead
}

// Insert into the skip list with the value
func (s *Skip) Insert(value int) {
	// turnNodes store the node in s.list[x] which point to the lower list
	turnNodes := make([]*Node, s.height+1)
	turnNode := s.Head()

	// find the target node which target.value <= value && value < target.next.value
	for i := s.height; 0 <= i; i-- {
		turnNode = turnNode.SearchForward(value)
		turnNodes[i] = turnNode
		turnNode = turnNode.down
	}

	// turnNodes[0] is the target node, we will insert the new node after the target node
	var high, low *Node
	for i := 0; i <= s.height; i++ {
		turnNodes[i].InsertAfter(value, &s.list[i].length)

		if i > 0 {
			// connect list[i+1] and list[i]
			high = turnNodes[i].next
			low = turnNodes[i-1].next
			high.down = low
			low.up = high
		}

		// if the list[i].length is odd, no more actions
		if s.list[i].length&1 == 1 {
			break
		}

		// if the list[i].length == 2, create the higher list and add height
		if s.list[i].length == 2 {
			s.height++

			// need scale when capacity <= height
			if s.capacity <= s.height {
				s.scale()
			}

			turnNodes = append(turnNodes, s.list[i+1].listHead)
		}
	}
}

// Delete delete the node which held the value
func (s *Skip) Delete(value int) {
	turnNode := s.Head()
	var target *Node

	// find the node which node.value <= value && value < node.next.value
	for i := s.height; 0 <= i; i-- {
		turnNode = turnNode.SearchForward(value)
		target = turnNode
		turnNode = turnNode.down
	}

	// return when turnNode.value != value
	if target.value != value {
		return
	}

	for i := 0; nil != target; i++ {
		target.Delete(&s.list[i].length)
		target = target.up
	}
}

// Search search the skip list with the value
func (s *Skip) Search(value int) *Node {
	turnNode := s.Head()
	var target *Node

	// find the node which node.value <= value && value < node.next.value
	for i := s.height; 0 <= i; i-- {
		turnNode = turnNode.SearchForward(value)
		target = turnNode
		turnNode = turnNode.down
	}

	// return the turnNode when turnNode.value = value
	if target.value == value {
		return target
	}

	return nil
}

// Print print the skip list(<=height)
func (s *Skip) Print(height int) string {
	r := make([]string, height+1)

	for i := height; 0 <= i; i-- {
		r[i] = s.list[i].Print()
	}

	tools.Reverse(r)
	return strings.Join(r, "\n")
}

// Print print the list
func (l *List) Print() string {
	node := l.listHead
	info := make([]string, l.length+1)

	for i := 0; node != nil; node = node.next {
		v := fmt.Sprintf("%d", node.value)
		if min == node.value {
			v = "head"
		}
		info[i] = v
		i++
	}

	return strings.Join(info, "<->")
}

// SearchForward search the list with the value backward,
// return the node which node.value < value && value < node.next.valu
func (n *Node) SearchForward(value int) *Node {
	node := n
	var result *Node

	for nil != node && node.value <= value {
		result = node
		node = node.next
	}

	return result
}

// InsertAfter insert the new node after n
func (n *Node) InsertAfter(value int, length *int) {
	newNode := newNode(value)

	next := n.next
	newNode.previous = n
	newNode.next = next

	if nil != next {
		next.previous = newNode
	}
	n.next = newNode

	*length++
}

// Delete delete the node which held the value
func (n *Node) Delete(length *int) {
	next := n.next
	prev := n.previous

	prev.next = next
	if nil != next {
		next.previous = prev
	}

	*length--
}
