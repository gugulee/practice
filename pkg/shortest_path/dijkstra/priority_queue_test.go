package dijkstra

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_IsFull(t *testing.T) {
	p := NewPriorityQueue(5)
	require.False(t, p.IsFull())

	p.used = 5
	require.True(t, p.IsFull())

	p.used = 6
	require.True(t, p.IsFull())
}

func Test_add(t *testing.T) {
	values := []int{7, 5, 8, 9, 10, 4, 6, 3}
	p := NewPriorityQueue(len(values))
	for i, v := range values {
		p.add(i, v)
	}
	expect := []vertex{
		{7, 3},
		{5, 4},
		{1, 5},
		{0, 7},
		{4, 10},
		{2, 8},
		{6, 6},
		{3, 9},
	}
	require.Equal(t, expect, p.slice())
}

func Test_pull(t *testing.T) {
	values := []int{7, 5, 8, 9, 10, 4, 6, 3}
	p := NewPriorityQueue(len(values))
	for i, v := range values {
		p.add(i, v)
	}

	expect := []vertex{
		{7, 3},
		{5, 4},
		{1, 5},
		{0, 7},
		{4, 10},
		{2, 8},
		{6, 6},
		{3, 9},
	}
	require.Equal(t, expect, p.slice())

	out := p.pull()
	expect = []vertex{
		{5, 4},
		{0, 7},
		{1, 5},
		{3, 9},
		{4, 10},
		{2, 8},
		{6, 6},
	}
	require.Equal(t, out.distance, 3)
	require.Equal(t, out.id, 7)
	require.Equal(t, expect, p.slice())
}

func Test_update(t *testing.T) {
	values := []int{7, 5, 8, 9, 10, 4, 6, 3}
	t.Run("the updated values is less than the parent node", func(t *testing.T) {
		p := NewPriorityQueue(len(values))
		for i, v := range values {
			p.add(i, v)
		}
		p.update(1, 2)
		expect := []vertex{
			{1, 2},
			{5, 4},
			{7, 3},
			{0, 7},
			{4, 10},
			{2, 8},
			{6, 6},
			{3, 9},
		}
		require.Equal(t, expect, p.slice())
	})

	t.Run("the updated values is greater than the parent node and the child node", func(t *testing.T) {
		p := NewPriorityQueue(len(values))
		for i, v := range values {
			p.add(i, v)
		}
		p.update(1, 15)
		expect := []vertex{
			{7, 3},
			{5, 4},
			{6, 6},
			{0, 7},
			{4, 10},
			{2, 8},
			{1, 15},
			{3, 9},
		}
		require.Equal(t, expect, p.slice())
	})

	t.Run("the updated values is greater than the parent node and less than the child node", func(t *testing.T) {
		p := NewPriorityQueue(len(values))
		for i, v := range values {
			p.add(i, v)
		}
		p.update(1, 4)
		expect := []vertex{
			{7, 3},
			{5, 4},
			{1, 4},
			{0, 7},
			{4, 10},
			{2, 8},
			{6, 6},
			{3, 9},
		}
		require.Equal(t, expect, p.slice())
	})

	t.Run("no data exist", func(t *testing.T) {
		p := NewPriorityQueue(len(values))
		for i, v := range values {
			p.add(i, v)
		}
		p.update(10, 4)
		expect := []vertex{
			{7, 3},
			{5, 4},
			{1, 5},
			{0, 7},
			{4, 10},
			{2, 8},
			{6, 6},
			{3, 9},
		}
		require.Equal(t, expect, p.slice())
	})
}
