package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrint(t *testing.T) {
	s := New(0)
	s.height = 3

	s.head.next[0] = NewNode(1, "1")
	s.head.next[0].next[0] = NewNode(1, "3")
	s.head.next[0].next[0].next[0] = NewNode(1, "5")
	s.head.next[0].next[0].next[0].next[0] = NewNode(1, "7")

	s.head.next[1] = NewNode(2, "1")
	s.head.next[1].next[1] = NewNode(2, "5")

	fmt.Println(s.Print())
}

func TestNew(t *testing.T) {
	s := New(0)
	require.Equal(t, defaultCapacity, s.head.layer)
}

func TestAllElement(t *testing.T) {
	t.Run("nomal insert", func(t *testing.T) {
		s := New(0)

		s.Insert("1")
		s.Insert("3")
		s.Insert("4")
		s.Insert("5")
		s.Insert("7")
		s.Insert("8")
		s.Insert("9")
		s.Insert("10")
		s.Insert("13")
		s.Insert("16")
		s.Insert("17")
		s.Insert("18")

		require.ElementsMatch(t, []string{"1", "3", "4", "5", "7", "8", "9", "10", "13", "16", "17", "18"}, s.AllElement())
	})

}

func TestInsert(t *testing.T) {
	t.Run("nomal insert", func(t *testing.T) {
		s := New(0)

		s.Insert("1")
		s.Insert("3")
		s.Insert("4")
		s.Insert("5")
		s.Insert("7")
		s.Insert("8")
		s.Insert("9")
		s.Insert("10")
		s.Insert("13")
		s.Insert("16")
		s.Insert("17")
		s.Insert("18")

		fmt.Println(s.Print())
	})

}

func TestSearch(t *testing.T) {
	s := New(0)

	s.Insert("1")
	s.Insert("3")
	s.Insert("4")
	s.Insert("5")
	s.Insert("7")
	s.Insert("8")
	s.Insert("9")
	s.Insert("10")
	s.Insert("13")
	s.Insert("16")
	s.Insert("17")
	s.Insert("18")

	fmt.Println(s.Print())
	fmt.Println("------------------------------------------------")

	expect := s.head.Search("17", 0)
	actual := s.Search("17")
	require.NotNil(t, actual)
	require.Equal(t, expect, actual)

	expect = s.head.Search("10", 0)
	actual = s.Search("10")
	require.NotNil(t, actual)
	require.Equal(t, expect, actual)

	expect = s.head.Search("1", 0)
	actual = s.Search("1")
	require.NotNil(t, actual)
	require.Equal(t, expect, actual)

	actual = s.Search("23")
	require.Nil(t, actual)

	actual = s.Search("6")
	require.Nil(t, actual)

	actual = s.Search("15")
	require.Nil(t, actual)
}

func TestDelete(t *testing.T) {
	s := New(0)

	s.Insert("1")
	s.Insert("3")
	s.Insert("4")
	s.Insert("17")
	s.Insert("18")

	fmt.Println(s.Print())
	fmt.Println("------------------------------------------------")

	s.Delete("17")
	actual := s.Search("17")
	require.Nil(t, actual)

	s.Delete("18")
	actual = s.Search("18")
	require.Nil(t, actual)

	s.Delete("1")
	actual = s.Search("1")
	require.Nil(t, actual)

	s.Delete("3")
	actual = s.Search("3")
	require.Nil(t, actual)

	s.Delete("4")
	actual = s.Search("4")
	require.Nil(t, actual)
	
	require.True(t, s.IsEmpty())
}
