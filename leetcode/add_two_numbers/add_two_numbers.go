package atn

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func InsertTail(n *ListNode, val int) {
	node := n
	for ; nil != node.Next; node = node.Next {
	}

	newNode := NewNode(val)
	newNode.Next = node.Next
	node.Next = newNode
}

func InsertAfter(n *ListNode, val int) {
	newNode := NewNode(val)
	newNode.Next = n.Next
	n.Next = newNode
}

func List(n *ListNode) []int {
	var ret []int
	node := n
	for ; nil != node; node = node.Next {
		ret = append(ret, node.Val)
	}

	return ret
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// this is a sentinel node
	l3 := NewNode(-1)
	l3Tail := l3

	tensDigit := 0
	for nil != l1 || nil != l2 {
		val := 0
		if nil != l1 {
			val = l1.Val
			l1 = l1.Next
		}
		if nil != l2 {
			val += l2.Val
			l2 = l2.Next
		}

		val += tensDigit
		oneDigit := val % 10
		tensDigit = val / 10

		InsertAfter(l3Tail, oneDigit)
		l3Tail = l3Tail.Next
	}

	if 0 != tensDigit {
		InsertAfter(l3Tail, tensDigit)
	}

	return l3.Next
}

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func NewNode(val int) *ListNode {
// 	return &ListNode{Val: val}
// }

// func (n *ListNode) Tail() *ListNode {
// 	node := n
// 	for ; nil != node.Next; node = node.Next {
// 	}
// 	return node
// }

// func (n *ListNode) InsertTail(val int) {
// 	tail := n.Tail()
// 	tail.InsertAfter(val)
// }

// func (n *ListNode) InsertAfter(val int) {
// 	newNode := NewNode(val)
// 	newNode.Next = n.Next
// 	n.Next = newNode
// }

// func (n *ListNode) List() []int {
// 	var ret []int
// 	node := n
// 	for ; nil != node; node = node.Next {
// 		ret = append(ret, node.Val)
// 	}

// 	return ret
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	// this is a sentinel node
// 	l3 := NewNode(-1)

// 	tensDigit := 0
// 	for nil != l1 || nil != l2 {
// 		val := 0
// 		if nil != l1 {
// 			val = l1.Val
// 			l1 = l1.Next
// 		}
// 		if nil != l2 {
// 			val += l2.Val
// 			l2 = l2.Next
// 		}

// 		val += tensDigit
// 		oneDigit := val % 10
// 		tensDigit = val / 10

// 		l3.InsertTail(oneDigit)
// 	}

// 	if 0 != tensDigit {
// 		l3.InsertTail(tensDigit)
// 	}

// 	return l3.Next
// }
