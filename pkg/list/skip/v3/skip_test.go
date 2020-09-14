package v3

import (
	"fmt"
	"testing"

	"github.com/practice/pkg/constants"
	vs "github.com/practice/pkg/utils/value_store"
	"github.com/stretchr/testify/require"
)

func TestNewNode(t *testing.T) {
	v := &vs.Ivalue{
		Value: 100,
	}

	n := NewNode(16, v)

	r := n.value.(*vs.Ivalue)
	fmt.Println(r.Value)
}

func Test_Skip_Print(t *testing.T) {
	s := New(0, &vs.Ivalue{Value: constants.MinInt})
	s.height = 3

	s.head.next[0] = NewNode(1, &vs.Ivalue{Value: 1})
	s.head.next[0].next[0] = NewNode(1, &vs.Ivalue{Value: 3})
	s.head.next[0].next[0].next[0] = NewNode(1, &vs.Ivalue{Value: 5})
	s.head.next[0].next[0].next[0].next[0] = NewNode(1, &vs.Ivalue{Value: 7})

	s.head.next[1] = NewNode(2, &vs.Ivalue{Value: 1})
	s.head.next[1].next[1] = NewNode(2, &vs.Ivalue{Value: 5})

	fmt.Println(s.Print())
}

func TestNew(t *testing.T) {
	s := New(0, &vs.Ivalue{Value: constants.MinInt})
	require.Equal(t, defaultCapacity, s.head.layer)
	require.Equal(t, &vs.Ivalue{Value: constants.MinInt}, s.head.value)
}

func TestInsert(t *testing.T) {
	t.Run("nomal insert", func(t *testing.T) {
		s := New(0, &vs.Ivalue{Value: constants.MinInt})

		in := []int{1, 3, 4, 4, 5, 7, 8, 9, 10, 13, 16, 17, 18}
		for _, i := range in {
			s.Insert(&vs.Ivalue{Value: i})
		}

		fmt.Println(s.Print())
	})

}

func TestSearch(t *testing.T) {
	s := New(0, &vs.Ivalue{Value: constants.MinInt})

	in := []int{1, 3, 4, 5, 7, 8, 9, 10, 13, 16, 17, 18}
	for _, i := range in {
		s.Insert(&vs.Ivalue{Value: i})
	}

	fmt.Println(s.Print())
	fmt.Println("------------------------------------------------")

	v := &vs.Ivalue{Value: 17}
	expect := s.head.Search(v, 0)
	actual := s.Search(v)
	require.NotNil(t, actual)
	require.Equal(t, expect, actual)

	v = &vs.Ivalue{Value: 10}
	expect = s.head.Search(v, 0)
	actual = s.Search(v)
	require.NotNil(t, actual)
	require.Equal(t, expect, actual)

	v = &vs.Ivalue{Value: 1}
	expect = s.head.Search(v, 0)
	actual = s.Search(v)
	require.NotNil(t, actual)
	require.Equal(t, expect, actual)

	v = &vs.Ivalue{Value: 23}
	actual = s.Search(v)
	require.Nil(t, actual)

	v = &vs.Ivalue{Value: 6}
	actual = s.Search(v)
	require.Nil(t, actual)

	v = &vs.Ivalue{Value: 15}
	actual = s.Search(v)
	require.Nil(t, actual)
}

func TestSearchRange(t *testing.T) {
	t.Run("no same element", func(t *testing.T) {
		s := New(0, &vs.Ivalue{Value: constants.MinInt})

		in := []int{1, 3, 4, 5, 7, 8, 9, 10, 13, 16, 17, 18}
		for _, i := range in {
			s.Insert(&vs.Ivalue{Value: i})
		}

		fmt.Println(s.Print())
		fmt.Println("------------------------------------------------")

		var outputs []int
		outs := s.SearchRange(&vs.Ivalue{Value: 6}, &vs.Ivalue{Value: 14})
		expect := []int{7, 8, 9, 10, 13}
		for i := range outs {
			outputs = append(outputs, outs[i].(*vs.Ivalue).Value)
		}
		require.ElementsMatch(t, expect, outputs)

		outs = s.SearchRange(&vs.Ivalue{Value: 5}, &vs.Ivalue{Value: 13})
		outputs = nil
		expect = []int{5, 7, 8, 9, 10, 13}
		for i := range outs {
			outputs = append(outputs, outs[i].(*vs.Ivalue).Value)
		}
		require.ElementsMatch(t, expect, outputs)
	})

	t.Run("has same element", func(t *testing.T) {
		s := New(0, &vs.Ivalue{Value: constants.MinInt})

		in := []int{1, 3, 4, 5, 5, 7, 8, 9, 10, 13, 13, 16, 17, 18}
		for _, i := range in {
			s.Insert(&vs.Ivalue{Value: i})
		}

		fmt.Println(s.Print())
		fmt.Println("------------------------------------------------")

		var outputs []int
		outs := s.SearchRange(&vs.Ivalue{Value: 5}, &vs.Ivalue{Value: 13})
		expect := []int{5, 5, 7, 8, 9, 10, 13, 13}
		for i := range outs {
			outputs = append(outputs, outs[i].(*vs.Ivalue).Value)
		}
		require.ElementsMatch(t, expect, outputs)

		outs = s.SearchRange(&vs.Ivalue{Value: 6}, &vs.Ivalue{Value: 14})
		outputs = nil
		expect = []int{7, 8, 9, 10, 13, 13}
		for i := range outs {
			outputs = append(outputs, outs[i].(*vs.Ivalue).Value)
		}
		require.ElementsMatch(t, expect, outputs)

		outs = s.SearchRange(&vs.Ivalue{Value: 6}, &vs.Ivalue{Value: 16})
		outputs = nil
		expect = []int{7, 8, 9, 10, 13, 13, 16}
		for i := range outs {
			outputs = append(outputs, outs[i].(*vs.Ivalue).Value)
		}
		require.ElementsMatch(t, expect, outputs)
	})
}

func TestSearchRank(t *testing.T) {
	t.Run("no same element", func(t *testing.T) {
		s := New(0, &vs.Ivalue{Value: constants.MinInt})

		in := []int{1, 3, 4, 5, 7, 8, 9, 10, 13, 16, 17, 18}
		for _, i := range in {
			s.Insert(&vs.Ivalue{Value: i})
		}

		fmt.Println(s.Print())
		fmt.Println("------------------------------------------------")

		var outputs []int
		outs := s.SearchRank(&vs.Ivalue{Value: 6}, 0, 4)
		expect := []int{7, 8, 9, 10}
		for i := range outs {
			outputs = append(outputs, outs[i].(*vs.Ivalue).Value)
		}
		require.ElementsMatch(t, expect, outputs)

		outs = s.SearchRank(&vs.Ivalue{Value: 5}, 0, 6)
		outputs = nil
		expect = []int{5, 7, 8, 9, 10, 13}
		for i := range outs {
			outputs = append(outputs, outs[i].(*vs.Ivalue).Value)
		}
		require.ElementsMatch(t, expect, outputs)

		outs = s.SearchRank(&vs.Ivalue{Value: 16}, 0, 10)
		outputs = nil
		expect = []int{16, 17, 18}
		for i := range outs {
			outputs = append(outputs, outs[i].(*vs.Ivalue).Value)
		}
		require.ElementsMatch(t, expect, outputs)
	})

	t.Run("has same element", func(t *testing.T) {
		s := New(0, &vs.Ivalue{Value: constants.MinInt})

		in := []int{1, 3, 4, 5, 5, 7, 8, 9, 10, 13, 13, 16, 17, 18}
		for _, i := range in {
			s.Insert(&vs.Ivalue{Value: i})
		}

		fmt.Println(s.Print())
		fmt.Println("------------------------------------------------")

		var outputs []int
		outs := s.SearchRank(&vs.Ivalue{Value: 5}, 0, 8)
		expect := []int{5, 5, 7, 8, 9, 10, 13, 13}
		for i := range outs {
			outputs = append(outputs, outs[i].(*vs.Ivalue).Value)
		}
		require.ElementsMatch(t, expect, outputs)

		outs = s.SearchRank(&vs.Ivalue{Value: 5}, 1, 7)
		outputs = nil
		expect = []int{5, 7, 8, 9, 10, 13, 13}
		for i := range outs {
			outputs = append(outputs, outs[i].(*vs.Ivalue).Value)
		}
		require.ElementsMatch(t, expect, outputs)

		outs = s.SearchRank(&vs.Ivalue{Value: 6}, 0, 7)
		outputs = nil
		expect = []int{7, 8, 9, 10, 13, 13, 16}
		for i := range outs {
			outputs = append(outputs, outs[i].(*vs.Ivalue).Value)
		}
		require.ElementsMatch(t, expect, outputs)
	})
}

func TestDelete(t *testing.T) {
	s := New(0, &vs.Ivalue{Value: constants.MinInt})

	in := []int{1, 3, 4, 5, 7, 8, 9, 10, 13, 16, 17, 18}
	for _, i := range in {
		s.Insert(&vs.Ivalue{Value: i})
	}

	fmt.Println(s.Print())
	fmt.Println("------------------------------------------------")

	v := &vs.Ivalue{Value: 17}
	s.Delete(v)
	actual := s.Search(v)
	require.Nil(t, actual)

	v = &vs.Ivalue{Value: 18}
	s.Delete(v)
	actual = s.Search(v)
	require.Nil(t, actual)

	v = &vs.Ivalue{Value: 1}
	s.Delete(v)
	actual = s.Search(v)
	require.Nil(t, actual)
}
