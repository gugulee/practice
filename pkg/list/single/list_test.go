package single

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"
)

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func TestInsetHead(t *testing.T) {
	tests := []struct {
		sep  string
		want string
	}{
		{sep: "hello"},
	}
	for _, test := range tests {
		var want []string
		for _, s := range test.sep {
			want = append(want, string(s))
		}
		reverse(want)
		test.want = strings.Join(want, "->")

		l := NewLinkList()
		for _, str := range test.sep {
			l.InsertHead(string(str))
		}
		out := l.Print()
		if test.want != out {
			t.Errorf("err: LIST: %q,out=%q,want %q", test.sep, out, test.want)
		}
		t.Logf("info: LIST: %q,length=%d", test.sep, l.len)
	}
}

func TestInsetTail(t *testing.T) {
	tests := []struct {
		sep  string
		want string
	}{
		{sep: "hello"},
	}
	for _, test := range tests {
		var want []string
		for _, s := range test.sep {
			want = append(want, string(s))
		}
		test.want = strings.Join(want, "->")

		l := NewLinkList()
		for _, str := range test.sep {
			l.InsertTail(string(str))
		}
		out := l.Print()
		if test.want != out {
			t.Errorf("err: LIST: %q,out=%q,want %q", test.sep, out, test.want)
		}
		t.Logf("info: LIST: %q,length=%d", test.sep, l.len)
	}
}

func TestSearchListBynode(t *testing.T) {
	l := NewLinkList()
	strs := "hello"
	for _, str := range strs {
		l.InsertTail(string(str))
	}
	tests := []struct {
		sep string
	}{
		{sep: "h"},
		{sep: "o"},
		{sep: "a"},
		{sep: "c"},
	}

	for _, test := range tests {
		if ! strings.Contains(strs, test.sep) {
			continue
		}
		node := l.SearchListByValue(test.sep)
		require.NotNilf(t, node, "LIST: %q, %q should be found", strs, test.sep)
		out := l.SearchListByNode(node)
		require.Truef(t, out, "LIST: %q, %q should be found", strs, test.sep)
	}
}

func TestSearchListByValue(t *testing.T) {
	l := NewLinkList()
	node := l.SearchListByValue("h")
	require.Nil(t, node)

	strs := "hello"
	for _, str := range strs {
		l.InsertTail(string(str))
	}

	tests := []struct {
		sep string
	}{
		{sep: "h"},
		{sep: "o"},
		{sep: "a"},
		{sep: "c"},
	}

	for _, test := range tests {
		node := l.SearchListByValue(test.sep)
		if ! strings.Contains(strs, test.sep) {
			require.Nil(t, node)
			continue
		}
		require.NotNilf(t, node, "LIST: %q, %q should be found", strs, test.sep)
	}
}

func TestDeleteNodeByValue(t *testing.T) {
	l := NewLinkList()
	strs := "hello"
	for _, str := range strs {
		l.InsertTail(string(str))
	}

	tests := []struct {
		sep string
	}{
		{sep: "o"},
		{sep: "l"},
		{sep: "l"},
		{sep: "a"},
		{sep: "e"},
		{sep: "h"},
	}
	for _, test := range tests {
		out := l.DeleteNodeByValue(test.sep)
		if strings.Contains(strs, test.sep) {
			require.Equal(t, true, out)
		}else {
			require.Equal(t, false, out)
		}
	}
}

func TestPrint(t *testing.T) {
	tests := []struct {
		sep  string
		want string
	}{
		{sep: "hello"},
		{sep: "world,lee"},
	}

	for _, test := range tests {
		var want []string
		for _, s := range test.sep {
			want = append(want, string(s))
		}
		test.want = strings.Join(want, "->")

		l := NewLinkList()
		for _, str := range test.sep {
			l.InsertTail(string(str))
		}
		out := l.Print()
		if test.want != out {
			t.Errorf("LIST: %q,out=%q,want %q", test.sep, out, test.want)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		sep  string
		want string
	}{
		{sep: ""},
		{sep: "h"},
		{sep: "ab"},
		{sep: "abc"},
		{sep: "abcdef"},
	}

	for _, test := range tests {
		var want []string
		for _, s := range test.sep {
			want = append(want, string(s))
		}
		reverse(want)
		test.want = strings.Join(want, "->")

		l := NewLinkList()
		for _, str := range test.sep {
			l.InsertTail(string(str))
		}
		l.Reverse()
		require.Equal(t, test.want, l.Print())
	}
}

func TestHasCycle1(t *testing.T) {
	l := NewLinkList()
	for _, str := range "abcdef" {
		l.InsertTail(string(str))
	}

	node := l.head
	var tmp *ListNode
	var i int
	for ; nil != node.next; node = node.next {
		if 3 == i {
			tmp = node
		}
		i++
	}

	// the next of the last node is head, the list has cycle
	node.next = l.head

	out := l.HasCycle1()
	require.Equal(t, true, out)

	// the next of the last node is own, the list has cycle
	node.next = node

	out = l.HasCycle1()
	require.Equal(t, true, out)

	// the next of the last node is a random node between the first and the last, the list has cycle
	node.next = tmp

	out = l.HasCycle1()
	require.Equal(t, true, out)

	// dismantle the cycle
	node.next = nil

	out = l.HasCycle1()
	require.Equal(t, false, out)
}

var _ = Describe("TestHasCycle2", func() {
	It("only head node", func() {
		l := NewLinkList()

		out := l.HasCycle2()
		Expect(out).Should(Equal(false))
	})

	It("only one node except for head", func() {
		l := NewLinkList()
		for _, str := range "a" {
			l.InsertTail(string(str))
		}

		node := l.head.next

		// the next of the last node is head, the list has cycle
		node.next = l.head

		out := l.HasCycle2()
		Expect(out).Should(Equal(true))

		// dismantle the cycle
		node.next = nil

		out = l.HasCycle2()
		Expect(out).Should(Equal(false))
	})

	It("two node at least except for head", func() {
		l := NewLinkList()
		for _, str := range "abcdef" {
			l.InsertTail(string(str))
		}
		node := l.head
		var tmp *ListNode
		var i int
		for ; nil != node.next; node = node.next {
			if 3 == i {
				tmp = node
			}
			i++
		}

		// the next of the last node is head, the list has cycle
		node.next = l.head

		out := l.HasCycle2()
		Expect(out).Should(Equal(true))

		// the next of the last node is own, the list has cycle
		node.next = node

		out = l.HasCycle2()
		Expect(out).Should(Equal(true))

		// the next of the last node is a random node between the first and the last, the list has cycle
		node.next = tmp

		out = l.HasCycle2()
		Expect(out).Should(Equal(true))

		// dismantle the cycle
		node.next = nil

		out = l.HasCycle1()
		Expect(out).Should(Equal(false))
	})
})

var _ = Describe("TestMergeSortedList", func() {
	It("l1 and l2 has no node, the new list has no node", func() {
		l1 := NewLinkList()
		l2 := NewLinkList()
		l3 := l1.MergeSortedList(l2)
		Expect(len(l3.Print())).Should(Equal(0))
		Expect(l3.len).Should(Equal(uint32(0)))
	})

	It("l1 has no node and l2 has node, the new list has l2`s node", func() {
		l1 := NewLinkList()
		l2 := NewLinkList()
		for _, str := range "a" {
			l2.InsertTail(string(str))
		}

		l3 := l1.MergeSortedList(l2)
		Expect(l3.Print()).Should(Equal("a"))
		Expect(l3.len).Should(Equal(uint32(1)))
	})

	It("l2 has no node and l1 has node, the new list has l1`s node", func() {
		l1 := NewLinkList()
		l2 := NewLinkList()
		for _, str := range "c" {
			l1.InsertTail(string(str))
		}

		l3 := l1.MergeSortedList(l2)
		Expect(l3.Print()).Should(Equal("c"))
		Expect(l3.len).Should(Equal(uint32(1)))
	})

	It("l1 and l2 has node, the new list has l1+l2`s node", func() {
		l1 := NewLinkList()
		l2 := NewLinkList()
		for _, str := range "1357" {
			l1.InsertTail(string(str))
		}

		for _, str := range "2468" {
			l2.InsertTail(string(str))
		}

		l3 := l1.MergeSortedList(l2)
		Expect(l3.Print()).Should(Equal("1->2->3->4->5->6->7->8"))
		Expect(l3.len).Should(Equal(uint32(8)))

		l1 = NewLinkList()
		l2 = NewLinkList()
		for _, str := range "13" {
			l1.InsertTail(string(str))
		}

		for _, str := range "2468" {
			l2.InsertTail(string(str))
		}

		l3 = l1.MergeSortedList(l2)
		Expect(l3.Print()).Should(Equal("1->2->3->4->6->8"))
		Expect(l3.len).Should(Equal(uint32(6)))
	})
})

var _ = Describe("TestDeleteBottomN", func() {
	It("only head node", func() {
		l := NewLinkList()

		err := l.DeleteBottomN(2)
		Expect(err).Should(MatchError("N can not be greater than the length of the list"))
	})

	It("only one node except for head", func() {
		l := NewLinkList()

		for _, str := range "a" {
			l.InsertTail(string(str))
		}

		err := l.DeleteBottomN(1)
		Expect(err).NotTo(HaveOccurred())
		Expect(l.Print()).Should(Equal(""))
	})

	It("two node at least except for head", func() {
		l := NewLinkList()

		for _, str := range "abcde" {
			l.InsertTail(string(str))
		}

		// delete the first from the bottom
		err := l.DeleteBottomN(1)
		Expect(err).NotTo(HaveOccurred())
		Expect(l.Print()).Should(Equal("a->b->c->d"))

		// delete the first node
		err = l.DeleteBottomN(4)
		Expect(err).NotTo(HaveOccurred())
		Expect(l.Print()).Should(Equal("b->c->d"))
	})
})

var _ = Describe("TestDeleteBottomN1", func() {
	It("only head node", func() {
		l := NewLinkList()

		err := l.DeleteBottomN1(2)
		Expect(err).Should(MatchError("N can not be greater than the length of the list"))
	})

	It("only one node except for head", func() {
		l := NewLinkList()

		for _, str := range "a" {
			l.InsertTail(string(str))
		}

		err := l.DeleteBottomN1(1)
		Expect(err).NotTo(HaveOccurred())
		Expect(l.Print()).Should(Equal(""))
	})

	It("two node at least except for head", func() {
		l := NewLinkList()

		for _, str := range "abcde" {
			l.InsertTail(string(str))
		}

		// delete the first from the bottom
		err := l.DeleteBottomN1(1)
		Expect(err).NotTo(HaveOccurred())
		Expect(l.Print()).Should(Equal("a->b->c->d"))

		// delete the first node
		err = l.DeleteBottomN1(4)
		Expect(err).NotTo(HaveOccurred())
		Expect(l.Print()).Should(Equal("b->c->d"))
	})
})

var _ = Describe("FindMiddleNode", func() {
	It("only head node", func() {
		l := NewLinkList()
		var nil *ListNode

		out := l.FindMiddleNode()
		Expect(out).Should(Equal(nil))
	})

	It("only one node except for head", func() {
		l := NewLinkList()
		for _, str := range "a" {
			l.InsertTail(string(str))
		}

		out := l.FindMiddleNode()
		Expect(out.value).Should(Equal("a"))

	})

	It("two node at least except for head", func() {
		l := NewLinkList()
		for _, str := range "abcdef" {
			l.InsertTail(string(str))
		}

		out := l.FindMiddleNode()
		Expect(out.value).Should(Equal("c"))

		l = NewLinkList()
		for _, str := range "abcdefg" {
			l.InsertTail(string(str))
		}

		out = l.FindMiddleNode()
		Expect(out.value).Should(Equal("d"))
	})
})
