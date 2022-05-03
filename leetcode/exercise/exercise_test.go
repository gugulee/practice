package exercise

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_exercise(t *testing.T) {
	root := &TreeNode{
		Val: -10,
	}
	node1 := &TreeNode{
		Val: 9,
	}

	node2 := &TreeNode{
		Val: 20,
	}

	node3 := &TreeNode{
		Val: 15,
	}

	node4 := &TreeNode{
		Val: 7,
	}

	root.Left = node1
	root.Right = node2

	node2.Left = node3
	node2.Right = node4

	out := maxPathSum(root)
	require.Equal(t, 42, out)
}
