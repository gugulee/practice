package cycle

import "fmt"

type ListNode struct {
	Next  *ListNode
	value interface{}
}

type LinkList struct {
	head *ListNode
	len  uint32
}

// create list node
func NewNode(value interface{}) *ListNode {
	return &ListNode{value: value}
}

// create list
func NewLinkList() *LinkList {
	list := &LinkList{NewNode(nil), 0}
	list.head.Next = list.head
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

// insert new node at list head
func (l *LinkList) InsertHead(value string) {
	head := l.head
	l.InsertAfter(head, value)
}

// insert new node at list tail
func (l *LinkList) InsertTail(value string) {
	head := l.head
	node := head
	for ; node.Next != head; node = node.Next {

	}

	l.InsertAfter(node, value)
}

// insert new node after node
func (l *LinkList) InsertAfter(node *ListNode, value string) {
	newNode := NewNode(value)

	newNode.Next = node.Next
	node.Next = newNode
	l.len++
}

// SearchListBynode determine whether the node is in the list, return true when exist, otherwise return false
func (l *LinkList) SearchListBynode(target *ListNode) bool {
	if nil == target {
		return false
	}
	node := l.head.Next
	for ; node != l.head; node = node.Next {
		if node == target {
			return true
		}
	}
	return false
}

// if value exist in list, return node, otherwise, return nil
func (l *LinkList) SearchListByValue(value string) *ListNode {
	node := l.head.Next
	for ; node != l.head; node = node.Next {
		if node.value == value {
			return node
		}
	}
	return nil
}

// delete node from list
func (l *LinkList) DeleteNode(target *ListNode) bool {
	if nil == target {
		return false
	}

	node := l.head.Next
	pre := l.head

	for ; node != l.head; node = node.Next {
		if node == target {
			break
		}
		pre = node
	}

	if node == l.head {
		return false
	}

	pre.Next = target.Next
	target = nil
	l.len--
	return true
}

// Print return the whole list
func (l *LinkList) Print() string {
	head := l.head
	next := head.Next
	info := ""

	for ; next != head; next = next.Next {
		info += fmt.Sprintf("%+v", next.value)
		if head != next.Next {
			info += "->"
		}
	}

	return info
}
