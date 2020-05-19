package single

import (
	"fmt"
)

type ListNode struct {
	next  *ListNode
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
	if nil == n {
		return nil
	}
	return n.value
}

func (n *ListNode) PNext() **ListNode {
	if nil == n {
		return nil
	}
	return &n.next
}

func (n *ListNode) Next() *ListNode {
	if nil == n {
		return nil
	}
	return n.next
}

//在链表尾部插入新节点
func (l *LinkList) InsertTail(value interface{}) {
	node := l.head
	for ; nil != node.next; node = node.next {

	}
	l.InsertAfter(node, value)
}

//在链表头后面插入新节点
func (l *LinkList) InsertHead(value interface{}) {
	node := l.head
	l.InsertAfter(node, value)
}

//在node后面插入新节点
func (l *LinkList) InsertAfter(node *ListNode, value interface{}) {
	newNode := NewNode(value)

	newNode.next = node.next
	node.next = newNode
	l.len++
}

//判断node是否在链表中，存在返回true，否则返回false
func (l *LinkList) SearchListByNode(target *ListNode) bool {
	if nil == target {
		return false
	}
	next := l.head.next
	for ; next != nil; next = next.next {
		if next == target {
			return true
		}
	}
	return false
}

//根据value查找是否存在链表中，返回该node，否则返回nil
func (l *LinkList) SearchListByValue(value interface{}) *ListNode {
	next := l.head.next
	for ; next != nil; next = next.next {
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
	node := l.head.next
	pre := l.head

	for ; nil != node; node = node.next {
		if node == target {
			break
		}
		pre = node
	}

	if nil == node {
		return false
	}

	pre.next = target.next
	l.len--
	return true
}

//根据value删除node
func (l *LinkList) DeleteNodeByValue(value string) bool {
	node := l.head.next
	pre := l.head
	for ; nil != node; node = node.next {
		if node.value == value {
			pre.next = node.next
			node = nil
			l.len--
			return true
		}
		pre = node
	}

	return false
}

func (l *LinkList) Print() string {
	next := l.head.next
	info := ""
	for ; next != nil; next = next.next {
		info += fmt.Sprintf("%+v", next.value)
		if nil != next.next {
			info += "->"
		}
	}
	return info
}

//reverse linked list
func (l *LinkList) Reverse() {
	// two node at least only need reverse, except for head
	// nil == l.head.next represent no node, except for head
	// nil == l.head.next.next represent one node, except for head
	if nil == l.head.next || nil == l.head.next.next {
		return
	}

	node := l.head.next

	var pre *ListNode = nil

	for nil != node {
		tmp := node.next
		node.next = pre
		pre = node
		node = tmp
	}

	l.head.next = pre
}

// HasCycle1 detect if list has cycle
func (l *LinkList) HasCycle1() bool {
	// one node at least only need check, except for head
	// nil == l.head.next represent no node, except for head
	if nil == l.head.next {
		return false
	}

	node := l.head

	// iterate l.len + 1 (including head), if the next node of the last node is not nil, the list has cycle
	for i := 0; i < (int(l.len) + 1); i++ {
		node = node.next
	}

	if nil != node {
		return true
	}

	return false
}

// HasCycle1 detect if list has cycle
func (l *LinkList) HasCycle2() bool {
	// nil == l.head.next represent no node except for head, it has no cycle
	if nil == l.head.next {
		return false
	}

	slow := l.head
	fast := l.head

	for nil != fast && nil != fast.next {
		slow = slow.next
		fast = fast.next.next

		if fast == slow {
			return true
		}
	}
	return false
}

// MergeSortedList merge list 'l' and 'l1' into a new list, return new list
func (l *LinkList) MergeSortedList(l1 *LinkList) *LinkList {
	newList := NewLinkList()
	if nil == l.head && nil == l1.head {
		return newList
	}

	node := l.head.next
	node1 := l1.head.next
	newNode := newList.head
	newList.len = l.len + l1.len

	for {
		if nil == node {
			newNode.next = node1
			break
		}

		if nil == node1 {
			newNode.next = node
			break
		}

		if node.value.(string) > node1.value.(string) {
			newNode.next = node1
			node1 = node1.next
		} else {
			newNode.next = node
			node = node.next
		}
		newNode = newNode.next
	}
	return newList
}

// DeleteBottomN delete the N node from the bottom
func (l *LinkList) DeleteBottomN(n int) error {
	// if n > l.len, can not delete
	if l.len < uint32(n) {
		return fmt.Errorf("N can not be greater than the length of the list")
	}

	if nil == l.head.next {
		return nil
	}

	node := l.head

	// find the the previous of the N node
	iterateCount := int(l.len) - n
	for i := 0; i < iterateCount; i++ {
		node = node.next
	}

	node.next = node.next.next
	l.len--
	return nil
}

// DeleteBottomN1 delete the N node from the bottom
func (l *LinkList) DeleteBottomN1(n int) error {
	// if n > l.len, can not delete
	if l.len < uint32(n) {
		return fmt.Errorf("N can not be greater than the length of the list")
	}

	if nil == l.head.next {
		return nil
	}

	fast := l.head
	slow := l.head
	for i := 0; i < n; i++ {
		fast = fast.next
	}

	for nil != fast.next {
		slow = slow.next
		fast = fast.next
	}

	slow.next = slow.next.next
	l.len--
	return nil
}

// FindMiddleNode find zhe middle node
func (l *LinkList) FindMiddleNode() *ListNode {
	if nil == l.head || nil == l.head.next {
		return nil
	}

	slow := l.head
	fast := l.head

	for nil != fast && nil != fast.next {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

// Tail return the last node
func (l *LinkList) Tail() *ListNode {
	if nil == l.head.next {
		return nil
	}

	node := l.head
	for ; nil != node.next; node = node.next {

	}
	return node
}
