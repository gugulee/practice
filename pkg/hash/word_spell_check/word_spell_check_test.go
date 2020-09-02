package wordspellcheck

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	require.Equal(t, 0, hash('a'))
	require.Equal(t, 1, hash('b'))
	require.Equal(t, 25, hash('z'))
}

func TestSearch(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		w := New()
		w.Insert("app")
		require.Equal(t, true, w.Search("app"))
		require.Equal(t, false, w.Search("a"))

		w.Insert("a")
		require.Equal(t, true, w.Search("a"))
		require.Equal(t, false, w.Search("aaaaaa"))
	})

	t.Run("test two", func(t *testing.T) {
		w := New()
		w.Insert("a")
		require.Equal(t, true, w.Search("a"))
		require.Equal(t, false, w.Search("app"))

		w.Insert("app")
		require.Equal(t, true, w.Search("app"))
	})
}

func TestWordSpellcheck(t *testing.T) {
	w := New()

	fi, err := os.Open(`word`)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)

	// insert
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		word := string(a)
		if 0 == len(word) {
			continue
		}
		w.Insert(word)
	}

	fi.Seek(0, 0)
	// search
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		word := string(a)
		if 0 == len(word) {
			continue
		}
		require.Equalf(t, true, w.Search(word), word)
	}
}
