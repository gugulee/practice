package cycle

import "fmt"

type ListNode struct {
	Next     *ListNode
	Previous *ListNode
	value    interface{}
}

type LinkList struct {
	head *ListNode
	len  uint32
}

func NewNode(value interface{}) *ListNode {
	return &ListNode{value: value}
}

func NewLinkList() *LinkList {
	list := &LinkList{NewNode(nil), 0}
	list.head.Next = list.head
	list.head.Previous = list.head
	return list
}

func (l *LinkList) Length() uint32 {
	return l.len
}

func (l *LinkList) Head() *ListNode {
	return l.head
}

func (n *ListNode) Value() interface{} {
	return n.value
}

// insert new node at the end of the list
func (l *LinkList) InsertTail(value string) {
	node := l.head
	l.InsertAfter(node.Previous, value)
}

// insert new node at the head of the list
func (l *LinkList) InsertHead(value string) {
	node := l.head
	l.InsertAfter(node, value)
}

// insert new node after node
func (l *LinkList) InsertAfter(node *ListNode, value string) {
	newNode := NewNode(value)

	next := node.Next

	newNode.Previous = node
	newNode.Next = next

	next.Previous = newNode
	node.Next = newNode

	l.len++
}

func (l *LinkList) Print() string {
	next := l.head.Next
	info := ""
	for ; next != l.head; next = next.Next {
		info += fmt.Sprintf("%+v", next.value)
		if l.head != next.Next {
			info += "->"
		}
	}
	return info
}

func (l *LinkList) ReversePrint() string {
	prev := l.head.Previous
	info := ""

	for ; prev != l.head; prev = prev.Previous {
		info += fmt.Sprintf("%+v", prev.value)
		if prev.Previous != l.head {
			info += "->"
		}
	}
	return info
}

// SearchListByNode return true if node is in the list, else return false
func (l *LinkList) SearchListByNode(target *ListNode) bool {
	if nil == target {
		return false
	}
	next := l.head.Next
	for ; next != l.head; next = next.Next {
		if next == target {
			return true
		}
	}
	return false
}

// SearchListByValue return the node which contain the value, else return nil
func (l *LinkList) SearchListByValue(value string) *ListNode {
	next := l.head.Next
	for ; next != l.head; next = next.Next {
		if next.value == value {
			return next
		}
	}
	return nil
}

// delete node by value
func (l *LinkList) DeleteNodeByValue(value string) bool {
	if 0 == len(value) {
		return false
	}
	node := l.head.Next

	for ; node != l.head; node = node.Next {
		if node.value == value {
			break
		}
	}

	if node == l.head {
		return false
	}

	prev := node.Previous
	next := node.Next

	prev.Next = next
	next.Previous = prev

	node = nil
	l.len--
	return true
}
