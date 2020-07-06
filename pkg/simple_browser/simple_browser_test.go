package simple_browser

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestForward(t *testing.T) {
	b := New()
	out := b.Forward("a")
	require.Equal(t, "a", out)
	require.Equal(t, []interface{}{"a"}, b.forward.Print())

	out = b.Forward("b")
	require.Equal(t, "b", out)
	require.Equal(t, []interface{}{"b", "a"}, b.forward.Print())

	out = b.Forward("c")
	require.Equal(t, "c", out)
	require.Equal(t, []interface{}{"c", "b", "a"}, b.forward.Print())

	out = b.Forward("")
	require.Equal(t, "", out)
	require.Equal(t, []interface{}{"c", "b", "a"}, b.forward.Print())

	b.backward.Push("d")
	out = b.Forward("e")
	require.Equal(t, "e", out)
	require.Equal(t, []interface{}{"e", "c", "b", "a"}, b.forward.Print())
	require.Equal(t, true, b.backward.IsEmpty())

	b.backward.Push("f")
	out = b.Forward("")
	require.Equal(t, "f", out)
	require.Equal(t, []interface{}{"f", "e", "c", "b", "a"}, b.forward.Print())
	require.Equal(t, true, b.backward.IsEmpty())
}

func TestBackward(t *testing.T) {
	b := New()
	b.Forward("a")
	b.Forward("b")
	b.Forward("c")

	out := b.Backward()
	require.Equal(t, "c", out)
	require.Equal(t, []interface{}{"b", "a"}, b.forward.Print())
	require.Equal(t, []interface{}{"c"}, b.backward.Print())

	out = b.Backward()
	require.Equal(t, "b", out)
	require.Equal(t, []interface{}{"a"}, b.forward.Print())
	require.Equal(t, []interface{}{"b", "c"}, b.backward.Print())

	out = b.Backward()
	require.Equal(t, "a", out)
	require.Equal(t, true, b.forward.IsEmpty())
	require.Equal(t, []interface{}{"a", "b", "c"}, b.backward.Print())

	out = b.Backward()
	require.Equal(t, "", out)
	require.Equal(t, true, b.forward.IsEmpty())
	require.Equal(t, []interface{}{"a", "b", "c"}, b.backward.Print())
}

func TestBrowser(t *testing.T) {
	b := New()
	b.Forward("a")
	b.Forward("b")
	b.Forward("c")

	out := b.Backward()
	require.Equal(t, "c", out)
	require.Equal(t, []interface{}{"b", "a"}, b.forward.Print())
	require.Equal(t, []interface{}{"c"}, b.backward.Print())

	out = b.Forward("d")
	require.Equal(t, "d", out)
	require.Equal(t, []interface{}{"d", "b", "a"}, b.forward.Print())
	require.Equal(t, true, b.backward.IsEmpty())
}
