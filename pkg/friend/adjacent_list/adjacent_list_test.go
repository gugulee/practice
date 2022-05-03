package adjacentlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// 1's friend is 2,
// 2's friend is 3,5,4,
// 4's friend is 1,2,
// 5's friend is 4,3
func TestInsert(t *testing.T) {
	r := require.New(t)
	f := New(0)
	f.Insert(1, 2)
	f.Insert(2, 3)
	f.Insert(2, 5)
	f.Insert(2, 4)
	f.Insert(4, 1)
	f.Insert(4, 2)
	f.Insert(5, 4)
	f.Insert(5, 3)

	out := f.PrintFollow()
	r.Equal("1-->2,2-->3,2-->4,2-->5,4-->1,4-->2,5-->3,5-->4", out)

	out = f.PrintFans()
	r.Equal("1-->4,2-->1,2-->4,3-->2,3-->5,4-->2,4-->5,5-->2", out)
}

func TestDelete(t *testing.T) {
	r := require.New(t)
	f := New(0)
	f.Insert(1, 2)
	f.Insert(2, 3)
	f.Insert(2, 5)
	f.Insert(2, 4)
	f.Insert(4, 1)
	f.Insert(4, 2)
	f.Insert(5, 4)
	f.Insert(5, 3)

	f.Delete(2, 3)
	out := f.PrintFollow()
	r.Equal("1-->2,2-->4,2-->5,4-->1,4-->2,5-->3,5-->4", out)

	out = f.PrintFans()
	r.Equal("1-->4,2-->1,2-->4,3-->5,4-->2,4-->5,5-->2", out)

	f.Delete(1, 3)
	out = f.PrintFollow()
	r.Equal("1-->2,2-->4,2-->5,4-->1,4-->2,5-->3,5-->4", out)

	out = f.PrintFans()
	r.Equal("1-->4,2-->1,2-->4,3-->5,4-->2,4-->5,5-->2", out)
}

func TestIsFollow(t *testing.T) {
	r := require.New(t)
	f := New(0)
	f.Insert(1, 2)
	f.Insert(2, 3)
	f.Insert(2, 5)
	f.Insert(2, 4)
	f.Insert(4, 1)
	f.Insert(4, 2)
	f.Insert(5, 4)
	f.Insert(5, 3)

	out := f.IsFollow(1, 2)
	r.True(out)

	out = f.IsFollow(2, 1)
	r.False(out)

	out = f.IsFollow(3, 1)
	r.False(out)

	out = f.IsFollow(5, 3)
	r.True(out)
}

func TestIsFans(t *testing.T) {
	r := require.New(t)
	f := New(0)
	f.Insert(1, 2)
	f.Insert(2, 3)
	f.Insert(2, 5)
	f.Insert(2, 4)
	f.Insert(4, 1)
	f.Insert(4, 2)
	f.Insert(5, 4)
	f.Insert(5, 3)

	out := f.IsFans(1, 2)
	r.True(out)

	out = f.IsFans(2, 3)
	r.True(out)

	out = f.IsFans(3, 2)
	r.False(out)
}
