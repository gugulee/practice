package single

import (
	"strings"

	"github.com/practice/pkg/list/single"
	plalindrome "github.com/practice/pkg/plalindrome/string"
)

/*
将单链表中的数据拷贝到slice中,然后进行判断
时间复杂度：O(n)
 */
func IsPalindrome1(l *single.LinkList) bool {
	if nil == l {
		return false
	}

	slice := []string{}
	node := l.Head().Next

	if nil == node {
		return false
	}

	for ; node != nil; node = node.Next {
		slice = append(slice, node.Value)
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

	node := l.Head().Next
	len := l.Length()
	s := make([]string, 0, len/2)
	if nil == node {
		return false
	}

	for i := uint32(1); i <= len/2; i++ {
		s = append(s, node.Value)
		node = node.Next
	}

	//if len is odd number,ignore the middle number
	if len%2 != 0 {
		node = node.Next
	}

	for i := int32(len/2 - 1); i >= 0; i-- {
		if s[i] != node.Value {
			return false
		}
		node = node.Next
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
	node := l.Head().Next
	len := l.Length()
	var pre *single.ListNode = nil

	if nil == node {
		return false
	}

	for i := uint32(1); i <= len/2; i++ {
		tmp := node.Next
		node.Next = pre
		pre = node
		node = tmp
	}

	left, right := pre, node
	if len%2 != 0 {
		right = right.Next
	}

	for nil != left && nil != right {
		if left.Value != right.Value {
			isPalindrome = false
			break
		}
		left = left.Next
		right = right.Next
	}

	for nil != pre {
		tmp := pre.Next
		pre.Next = node
		node = pre
		pre = tmp
	}

	return isPalindrome
}
