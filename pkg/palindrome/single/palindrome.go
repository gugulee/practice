package single

import (
	"strings"

	"github.com/practice/pkg/list/single"
	plalindrome "github.com/practice/pkg/palindrome/string"
)

/*
将单链表中的数据拷贝到slice中,然后进行判断
时间复杂度：O(n)
*/
func IsPalindrome1(l *single.LinkList) bool {
	if nil == l {
		return false
	}

	var slice []string
	node := l.Head().Next()

	if nil == node {
		return false
	}

	for ; node != nil; {
		slice = append(slice, node.Value().(string))
		node = node.Next()
	}

	return plalindrome.IsPlalindrome1(strings.Join(slice, ""))
}

/*
用slice来存放链表的前半部分数据
*/
func IsPalindrome2(l *single.LinkList) bool {
	if nil == l {
		return false
	}

	node := l.Head().Next()
	length := l.Length()
	s := make([]string, 0, length/2)
	if nil == node {
		return false
	}

	for i := uint32(1); i <= length/2; i++ {
		s = append(s, node.Value().(string))
		node = node.Next()
	}

	//if len is odd number,ignore the middle number
	if length%2 != 0 {
		node = node.Next()
	}

	for i := int32(length/2 - 1); i >= 0; i-- {
		if s[i] != node.Value() {
			return false
		}
		node = node.Next()
	}

	return true
}

/*
找到链表中间节点，将前半部分转置，再从中间向左右遍历对比
时间复杂度：O(N)
*/
func IsPalindrome3(l *single.LinkList) bool {
	if nil == l {
		return false
	}
	isPalindrome := true
	node := l.Head().PNext()
	length := l.Length()
	var pre *single.ListNode = nil

	if nil == *node {
		return false
	}

	// reversr the first half part
	for i := uint32(1); i <= length/2; i++ {
		tmp := *((*node).PNext())
		pNext := (*node).PNext()
		*pNext = pre
		pre = *node
		*node = tmp
	}

	// if the node of list is odd, skip the middle node
	left, right := pre, *node
	if length%2 != 0 {
		right = *((*right).PNext())
	}

	for nil != left && nil != right {
		if left.Value() != (*right).Value() {
			isPalindrome = false
			break
		}
		left = *((*left).PNext())
		right = *((*right).PNext())
	}

	for nil != pre {
		tmp := *(pre.PNext())
		pNext := pre.PNext()
		*pNext = *node
		*node = pre
		pre = tmp
	}

	return isPalindrome
}
