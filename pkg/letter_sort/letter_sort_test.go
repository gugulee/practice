package lettersort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLetterSort(t *testing.T) {
	in := []string{"D", "a", "F", "B", "c", "A", "z"}
	letterSort(in)
	require.Equal(t, []string{"z", "a", "c", "B", "F", "A", "D"}, in)

	in = []string{"c", "a", "d", "e", "c", "z"}
	letterSort(in)
	require.Equal(t, []string{"c", "a", "d", "e", "c", "z"}, in)
}

func TestLetterSort1(t *testing.T) {
	in := []string{"0", "D", "a", "F", "3", "B", "c", "2", "A", "z", "7"}
	letterSort1(in)
	require.Equal(t, []string{"a", "c", "z", "0", "1", "3", "2", "7", "D", "F", "B", "A"}, in)
}
