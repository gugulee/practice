package slngle_list

import (
	"strings"
)

type ListNode struct {
	Next  *ListNode
	Value string
}

type LinkList struct {
	head *ListNode
	len  uint32
}

//创建新node
func NewNode(value string) *ListNode {
	return &ListNode{Value: value}
}

//创建链表头
func NewLinkList() *LinkList {
	return &LinkList{NewNode(""), 0}
}

func (l *LinkList) Length() uint32 {
	return l.len
}

func (l *LinkList) Head() *ListNode {
	return l.head
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
func (l *LinkList) SearchListBynode(target *ListNode) bool {
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
		if next.Value == value {
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
		if node.Value == value {
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
	infoSlice := []string{}
	for ; next != nil; next = next.Next {
		infoSlice = append(infoSlice, next.Value)
	}
	return strings.Join(infoSlice, "->")
}
