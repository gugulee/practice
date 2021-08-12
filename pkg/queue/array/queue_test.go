package array

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsEmpty(t *testing.T) {
	q := New(10)
	out := q.IsEmpty()
	require.Equal(t, true, out)
}

func TestQueue(t *testing.T) {
	t.Run("the length of queue is 3, only 3 element in queue", func(t *testing.T) {
		q := New(3)
		q.Enqueue("a")
		q.Enqueue("b")
		q.Enqueue("c")
		q.Enqueue("d")
		q.Enqueue("e")
		require.Equal(t, []interface{}{"a", "b", "c"}, q.Print())
	})

	t.Run("enqueue and dequeue normally", func(t *testing.T) {
		q := New(10)
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
		q.Enqueue("e")

		out := q.Dequeue()
		require.Equal(t, "a", out)
		require.Equal(t, []interface{}{"b", "c", "d", "e"}, q.Print())

		out = q.Dequeue()
		require.Equal(t, "b", out)
		require.Equal(t, []interface{}{"c", "d", "e"}, q.Print())

		q.Enqueue("f")
		require.Equal(t, []interface{}{"c", "d", "e", "f"}, q.Print())
		require.Equal(t, 0, q.head)

		
		q.Enqueue("g")
		require.Equal(t, []interface{}{"c", "d", "e", "f","g"}, q.Print())
		require.Equal(t, 0, q.head)

		q.Enqueue("h")
		require.Equal(t, []interface{}{"c", "d", "e", "f","g"}, q.Print())
		require.Equal(t, 0, q.head)
	})
}

func TestPrint(t *testing.T) {
	q := New(10)
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")
	q.Enqueue("d")
	require.Equal(t, []interface{}{"a", "b", "c", "d"}, q.Print())
}
