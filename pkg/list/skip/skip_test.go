package skip

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListPrint(t *testing.T) {
	l := newList(min)
	node := l.listHead

	node.next = newNode(10)
	l.length++
	node = node.next

	node.next = newNode(20)
	l.length++
	node = node.next

	node.next = newNode(30)
	l.length++
	node = node.next

	fmt.Println(l.Print())
}

func TestSearchForward(t *testing.T) {
	l := newList(min)
	node := l.listHead

	require.Equal(t, min, l.listHead.SearchForward(100).value)

	node.next = newNode(2)
	l.length++
	node = node.next

	node.next = newNode(4)
	l.length++
	node = node.next

	node.next = newNode(7)
	l.length++
	node = node.next

	fmt.Println(l.Print())

	require.Equal(t, min, l.listHead.SearchForward(1).value)
	require.Equal(t, 2, l.listHead.SearchForward(3).value)
	require.Equal(t, 7, l.listHead.SearchForward(9).value)
}

func TestInsertAfter(t *testing.T) {
	l := newList(min)
	node := l.listHead

	node.InsertAfter(1, &l.length)
	require.Equal(t, 1, node.next.value)
	require.Equal(t, min, node.next.previous.value)

	node = node.next
	node.InsertAfter(3, &l.length)
	require.Equal(t, 3, node.next.value)
	require.Equal(t, 1, node.next.previous.value)

	node = l.listHead.next
	node.InsertAfter(2, &l.length)
	require.Equal(t, 2, node.next.value)
	require.Equal(t, 1, node.next.previous.value)
}
func TestNew(t *testing.T) {
	s := New(4)

	for i := s.capacity - 1; i > 0; i-- {
		require.Equal(t, s.list[i-1].listHead, s.list[i].listHead.down)
		require.Equal(t, s.list[i].listHead, s.list[i-1].listHead.up)
	}
}

func TestScale(t *testing.T) {
	s := New(0)
	s.scale()

	require.Equal(t, 8, s.capacity)

	for i := s.capacity - 1; i > 0; i-- {
		require.Equal(t, s.list[i-1].listHead, s.list[i].listHead.down)
		require.Equal(t, s.list[i].listHead, s.list[i-1].listHead.up)
	}
}

func TestInsert(t *testing.T) {
	t.Run("nomal insert", func(t *testing.T) {
		s := New(0)

		s.Insert(1)
		s.Insert(3)
		s.Insert(4)
		s.Insert(5)
		s.Insert(7)
		s.Insert(8)
		s.Insert(9)
		s.Insert(10)
		s.Insert(13)
		s.Insert(16)
		s.Insert(17)
		s.Insert(18)
		s.Insert(20)
		s.Insert(22)
		s.Insert(24)
		s.Insert(26)
		s.Insert(27)
		s.Insert(0)
		s.Insert(28)
		s.Insert(2)
		s.Insert(6)
		s.Insert(30)
		s.Insert(21)

		fmt.Println()

		a := s.list[0].listHead.SearchForward(3)
		b := s.list[1].listHead.SearchForward(3)
		require.Equal(t, a, b.down)
		require.Equal(t, b, a.up)

		a = s.list[0].listHead.SearchForward(5)
		b = s.list[1].listHead.SearchForward(5)
		c := s.list[2].listHead.SearchForward(5)
		require.Equal(t, a, b.down)
		require.Equal(t, b, a.up)
		require.Equal(t, b, c.down)
		require.Equal(t, c, b.up)

		a = s.list[0].listHead.SearchForward(10)
		b = s.list[1].listHead.SearchForward(10)
		c = s.list[2].listHead.SearchForward(10)
		d := s.list[3].listHead.SearchForward(10)
		require.Equal(t, a, b.down)
		require.Equal(t, b, a.up)
		require.Equal(t, b, c.down)
		require.Equal(t, c, b.up)
		require.Equal(t, c, d.down)
		require.Equal(t, d, c.up)

		a = s.list[0].listHead.SearchForward(18)
		b = s.list[1].listHead.SearchForward(18)
		c = s.list[2].listHead.SearchForward(18)
		require.Equal(t, a, b.down)
		require.Equal(t, b, a.up)
		require.Equal(t, b, c.down)
		require.Equal(t, c, b.up)

		require.Equal(t, "head<->26"+"\n"+
			"head<->10<->26"+"\n"+
			"head<->2<->5<->10<->18<->26"+"\n"+
			"head<->0<->2<->3<->5<->8<->10<->16<->18<->22<->26<->30"+"\n"+
			"head<->0<->1<->2<->3<->4<->5<->6<->7<->8<->9<->10<->13<->16<->17<->18<->20<->21<->22<->24<->26<->27<->28<->30", s.Print(s.height))
	})

	t.Run("nomal insert", func(t *testing.T) {
		s := New(0)

		s.Insert(100)

		s.Insert(1)
		s.Insert(101)

		s.Insert(2)
		s.Insert(102)

		s.Insert(3)
		s.Insert(103)

		s.Insert(4)
		s.Insert(104)

		s.Insert(5)
		s.Insert(105)

		s.Insert(6)
		s.Insert(106)

		s.Insert(7)
		s.Insert(107)

		s.Insert(8)
		s.Insert(108)

		s.Insert(9)
		s.Insert(109)

		s.Insert(10)
		s.Insert(110)

		s.Insert(11)
		s.Insert(111)

		fmt.Println(s.Print(s.height))
	})
}

func TestSearch(t *testing.T) {
	s := New(0)

	s.Insert(1)
	s.Insert(3)
	s.Insert(4)
	s.Insert(5)
	s.Insert(7)
	s.Insert(8)
	s.Insert(9)
	s.Insert(10)
	s.Insert(13)
	s.Insert(16)
	s.Insert(17)
	s.Insert(18)

	fmt.Println(s.Print(s.height))
	fmt.Println("------------------------------------------------")

	expect := s.list[0].listHead.SearchForward(1)
	actual := s.Search(1)
	require.Equal(t, expect, actual)

	expect = s.list[0].listHead.SearchForward(3)
	actual = s.Search(3)
	require.Equal(t, expect, actual)

	expect = s.list[0].listHead.SearchForward(5)
	actual = s.Search(5)
	require.Equal(t, expect, actual)

	expect = s.list[0].listHead.SearchForward(9)
	actual = s.Search(9)
	require.Equal(t, expect, actual)

	expect = s.list[0].listHead.SearchForward(16)
	actual = s.Search(16)
	require.Equal(t, expect, actual)

	expect = s.list[0].listHead.SearchForward(18)
	actual = s.Search(18)
	require.Equal(t, expect, actual)

	actual = s.Search(0)
	require.Equal(t, (*Node)(nil), actual)

	actual = s.Search(15)
	require.Equal(t, (*Node)(nil), actual)

	actual = s.Search(20)
	require.Equal(t, (*Node)(nil), actual)
}

func TestDelete(t *testing.T) {
	t.Run("nomal insert", func(t *testing.T) {
		s := New(0)

		s.Insert(1)
		s.Insert(3)
		s.Insert(4)
		s.Insert(5)
		s.Insert(7)
		s.Insert(8)
		s.Insert(9)
		s.Insert(10)
		s.Insert(13)
		s.Insert(16)
		s.Insert(17)
		s.Insert(18)

		fmt.Println(s.Print(s.height))
		fmt.Println("------------------------------------------------")

		input := 3
		s.Delete(input)
		require.Equal(t, (*Node)(nil), s.Search(input))

		input = 18
		s.Delete(input)
		require.Equal(t, (*Node)(nil), s.Search(input))

		input = 8
		s.Delete(input)
		require.Equal(t, (*Node)(nil), s.Search(input))

		input = 20
		s.Delete(input)
		require.Equal(t, (*Node)(nil), s.Search(input))
	})

	t.Run("nomal insert", func(t *testing.T) {
		s := New(0)

		s.Insert(1)

		fmt.Println(s.Print(s.height))
		fmt.Println("------------------------------------------------")

		input := 1
		s.Delete(input)
		require.Equal(t, (*Node)(nil), s.Search(input))
	})
}
