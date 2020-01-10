package search_dir

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDirBfs(t *testing.T) {
	dir := `D:\doc`
	err := dirBfs(dir)
	require.NoError(t, err)
}

func TestDirDfs(t *testing.T) {
	dir := `D:\doc`
	err := dirDfs(dir)
	require.NoError(t, err)
}
