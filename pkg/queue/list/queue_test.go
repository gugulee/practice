package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsEmpty(t *testing.T) {
	q := New()
	out := q.IsEmpty()
	require.Equal(t, true, out)
}

func TestQueue(t *testing.T) {
	t.Run("the queue should nil", func(t *testing.T) {
		q := New()
		require.Equal(t, []interface{}(nil), q.Print())
	})

	t.Run("enqueue and dequeue normally", func(t *testing.T) {
		q := New()
		q.Enqueue("a")
		q.Enqueue("b")
		q.Enqueue("c")
		q.Enqueue("d")
		q.Enqueue("e")

		out := q.Dequeue()
		require.Equal(t, "a", out)
		require.Equal(t, []interface{}{"b", "c", "d", "e"}, q.Print())

		out = q.Dequeue()
		require.Equal(t, "b", out)
		require.Equal(t, []interface{}{"c", "d", "e"}, q.Print())

		q.Enqueue("f")
		require.Equal(t, []interface{}{"c", "d", "e", "f"}, q.Print())

		q.Enqueue("g")
		require.Equal(t, []interface{}{"c", "d", "e", "f", "g"}, q.Print())

		q.Dequeue()
		q.Dequeue()
		q.Dequeue()
		q.Dequeue()
		q.Dequeue()
		q.Dequeue()
		q.Dequeue()
		require.Equal(t, []interface{}(nil), q.Print())
	})
}

func TestPrint(t *testing.T) {
	q := New()
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")
	q.Enqueue("d")
	require.Equal(t, []interface{}{"a", "b", "c", "d"}, q.Print())
}
