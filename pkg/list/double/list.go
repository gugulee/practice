package double

import (
	"fmt"
)

type ListNode struct {
	Next     *ListNode
	Previous *ListNode
	value    interface{}
}

type LinkList struct {
	head *ListNode
	len  uint32
}

//创建新node
func NewNode(value interface{}) *ListNode {
	return &ListNode{value: value}
}

//创建链表头
func NewLinkList() *LinkList {
	return &LinkList{NewNode(nil), 0}
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

//在链表尾部插入新节点
func (l *LinkList) InsertTail(value string) {
	node := l.head
	for nil != node.Next {
		node = node.Next
	}
	l.InsertAfter(node, value)
}

//在链表头后面插入新节点
func (l *LinkList) InsertHead(value string) {
	node := l.head
	l.InsertAfter(node, value)
}

//在node后面插入新节点
func (l *LinkList) InsertAfter(node *ListNode, value string) {
	newNode := NewNode(value)

	next := node.Next
	newNode.Previous = node
	newNode.Next = next

	if nil != next {
		next.Previous = newNode
	}
	node.Next = newNode

	l.len++
}

//判断node是否在链表中，存在返回true，否则返回false
func (l *LinkList) SearchListByNode(target *ListNode) bool {
	if nil == target {
		return false
	}
	next := l.head.Next
	for ; next != nil; next = next.Next {
		if next == target {
			return true
		}
	}
	return false
}

//根据value查找是否存在链表中，返回该node，否则返回nil
func (l *LinkList) SearchListByValue(value string) *ListNode {
	next := l.head.Next
	for ; next != nil; next = next.Next {
		if next.value == value {
			return next
		}
	}
	return nil
}

// delete node
func (l *LinkList) DeleteNode(target *ListNode) bool {
	if nil == target {
		return false
	}
	node := l.head.Next

	for ; node != nil; node = node.Next {
		if node == target {
			break
		}
	}

	if node == nil {
		return false
	}

	prev := target.Previous
	next := target.Next

	prev.Next = next
	if nil != next {
		next.Previous = prev
	}

	target = nil
	l.len--
	return true
}

// delete node by value
func (l *LinkList) DeleteNodeByValue(value string) bool {
	if 0 == len(value) {
		return false
	}
	node := l.head.Next

	for ; node != nil; node = node.Next {
		if node.value == value {
			break
		}
	}

	if node == nil {
		return false
	}

	prev := node.Previous
	next := node.Next

	prev.Next = next
	if nil != next {
		next.Previous = prev
	}

	node = nil
	l.len--
	return true
}

func (l *LinkList) Print() string {
	next := l.head.Next
	info := ""
	for ; next != nil; next = next.Next {
		info += fmt.Sprintf("%+v", next.value)
		if nil != next.Next {
			info += "->"
		}
	}
	return info
}

func (l *LinkList) ReversePrint() string {
	node := l.head
	info := ""

	// till the latest node
	for ; nil != node.Next; node = node.Next {

	}

	previous := node

	for ; previous != l.head; previous = previous.Previous {
		info += fmt.Sprintf("%+v", previous.value)
		if previous.Previous != l.head {
			info += "->"
		}
	}
	return info
}
