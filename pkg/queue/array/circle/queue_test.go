package circle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {

	t.Run("the length of queue is 3, only 3 element in queue", func(t *testing.T) {
		q := New(3)
		q.Enqueue("a")
		q.Enqueue("b")
		q.Enqueue("c")
		q.Enqueue("d")
		q.Enqueue("e")
		require.Equal(t, []interface{}{"a", "b"}, q.Print())
	})

	t.Run("enqueue and dequeue normally", func(t *testing.T) {
		q := New(5)
		q.Enqueue("a")
		q.Enqueue("b")
		q.Enqueue("c")
		q.Enqueue("d")
		require.Equal(t, []interface{}{"a", "b", "c", "d"}, q.Print())

		out := q.Dequeue()
		require.Equal(t, "a", out)
		require.Equal(t, []interface{}{"b", "c", "d"}, q.Print())

		out = q.Dequeue()
		require.Equal(t, "b", out)
		require.Equal(t, []interface{}{"c", "d"}, q.Print())

		out = q.Dequeue()
		require.Equal(t, "c", out)
		require.Equal(t, []interface{}{"d"}, q.Print())

		out = q.Dequeue()
		require.Equal(t, "d", out)
		require.Equal(t, []interface{}(nil), q.Print())
	})

	t.Run("enqueue and dequeue abnormally", func(t *testing.T) {
		q := New(5)
		q.Enqueue("a")
		q.Enqueue("b")
		q.Enqueue("c")
		q.Enqueue("d")

		out := q.Dequeue()
		require.Equal(t, "a", out)
		require.Equal(t, []interface{}{"b", "c", "d"}, q.Print())
		require.Equal(t, 1, q.head)
		require.Equal(t, 4, q.tail)

		out = q.Dequeue()
		require.Equal(t, "b", out)
		require.Equal(t, []interface{}{"c", "d"}, q.Print())
		require.Equal(t, 2, q.head)
		require.Equal(t, 4, q.tail)

		q.Enqueue("e")
		require.Equal(t, []interface{}{"c", "d", "e"}, q.Print())
		require.Equal(t, 2, q.head)
		require.Equal(t, 0, q.tail)

		q.Enqueue("f")
		require.Equal(t, []interface{}{"c", "d", "e", "f"}, q.Print())
		require.Equal(t, 2, q.head)
		require.Equal(t, 1, q.tail)

		q.Enqueue("g")
		require.Equal(t, []interface{}{"c", "d", "e", "f"}, q.Print())
		require.Equal(t, 2, q.head)
		require.Equal(t, 1, q.tail)
	})
}

func TestPrint(t *testing.T) {
	q := New(5)
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")
	q.Enqueue("d")
	q.Enqueue("e")
	q.Print()
	require.Equal(t, []interface{}{"a", "b", "c", "d"}, q.Print())
}
