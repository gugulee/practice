package single

import (
	"fmt"
)

type ListNode struct {
	Next  *ListNode
	value interface{}
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
	for ; nil != node.Next; node = node.Next {

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

	newNode.Next = node.Next
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

//删除node
func (l *LinkList) DeleteNode(target *ListNode) bool {
	if nil == target {
		return false
	}
	node := l.head.Next
	pre := l.head

	for ; nil != node; node = node.Next {
		if node == target {
			break
		}
		pre = node
	}

	if nil == node {
		return false
	}

	pre.Next = target.Next
	target = nil
	l.len--
	return true
}

//根据value删除node
func (l *LinkList) DeleteNodeByValue(value string) bool {
	node := l.head.Next
	pre := l.head
	for ; nil != node; node = node.Next {
		if node.value == value {
			pre.Next = node.Next
			node = nil
			l.len--
			return true
		}
		pre = node
	}

	return false
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

//reverse linked list
func (l *LinkList) Reverse() {
	if nil == l.head {
		return
	}

	node := l.head.Next
	// one node at least except for head
	if nil == node {
		return
	}

	var pre *ListNode = nil

	for nil != node {
		tmp := node.Next
		node.Next = pre
		pre = node
		node = tmp
	}

	l.head.Next = pre
}
