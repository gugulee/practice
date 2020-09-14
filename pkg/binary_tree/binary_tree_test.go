package binarytree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinaryTree(t *testing.T) {
	bt := New()
	a := newNode("A")
	b := newNode("B")
	c := newNode("C")
	d := newNode("D")
	e := newNode("E")
	f := newNode("F")
	g := newNode("G")

	bt.head.left = a
	bt.head.right = a

	a.left = b
	a.right = c

	b.left = d
	b.right = e

	c.left = f
	c.right = g

	/* the recursive traversal */
	out := bt.PreOrder()
	require.Equal(t, "A->B->D->E->C->F->G", out)

	out = bt.PostOrder()
	require.Equal(t, "D->E->B->F->G->C->A", out)

	out = bt.InOrder()
	require.Equal(t, "D->B->E->A->F->C->G", out)

	/* the non-recursive traversal */
	out = bt.PreOrder1()
	require.Equal(t, "A->B->D->E->C->F->G", out)

	out = bt.PostOrder1()
	require.Equal(t, "D->E->B->F->G->C->A", out)

	out = bt.InOrder1()
	require.Equal(t, "D->B->E->A->F->C->G", out)

	out = bt.BFS()
	require.Equal(t, "A->B->C->D->E->F->G", out)

	/* optimize */
	out = bt.InOrder2()
	require.Equal(t, "D->B->E->A->F->C->G", out)

	out = bt.PostOrder2()
	require.Equal(t, "D->E->B->F->G->C->A", out)
}
