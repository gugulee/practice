package v2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	out := hash("hello", 16)
	fmt.Println(out)
}

func TestNewDictHt(t *testing.T) {
	out := NewDictHt(0)
	require.Equal(t, 16, len(out.node))

	out = NewDictHt(16)
	require.Equal(t, 16, len(out.node))

	out = NewDictHt(31)
	require.Equal(t, 32, len(out.node))

	out = NewDictHt(32)
	require.Equal(t, 32, len(out.node))
}

func TestNew(t *testing.T) {
	out := New(0)
	require.Equal(t, 16, len(out.ht[0].node))
	require.Equal(t, -1, out.rehashIndex)
	require.Nil(t, out.ht[1])
}

func TestNeedRehash(t *testing.T) {
	out := New(0)
	out.ht[0].used = 14
	require.True(t, out.NeedExpand())

	out.ht[0].used = 9
	require.False(t, out.NeedExpand())
}

func TestDhInsert(t *testing.T) {
	d := New(0)
	d.ht[0].Insert("hello")
	require.Equal(t, d.ht[0].used, 1)

	d.ht[0].Insert("helloa")
	require.Equal(t, d.ht[0].used, 1)

	d.ht[0].Insert("helloaa")
	require.Equal(t, d.ht[0].used, 1)

	d.ht[0].Insert("world")
	require.Equal(t, d.ht[0].used, 2)

	d.ht[0].Insert("worlda")
	require.Equal(t, d.ht[0].used, 2)

	d.ht[0].Insert("practice")
	require.Equal(t, d.ht[0].used, 3)

	d.ht[0].Insert("practicae")
	require.Equal(t, d.ht[0].used, 4)
}

func TestMigrateData(t *testing.T) {
	d := New(0)
	d.ht[1] = NewDictHt(2 * d.ht[0].capacity)

	d.ht[0].Insert("hello")
	d.ht[0].Insert("world")
	d.ht[0].Insert("worlda")
	d.ht[0].Insert("practice")

	d.MigrateData()
	require.Equal(t, hash("world", d.ht[0].capacity), d.rehashIndex)
	require.Equal(t, 2, d.ht[0].used)
	require.Equal(t, 1, d.ht[1].used)
	require.Equal(t, 32, d.ht[1].capacity)

	word := "world"
	hashValue := hash(word, d.ht[1].capacity)
	require.NotNil(t, d.ht[1].node[hashValue].Search(word))

	word = "worlda"
	hashValue = hash(word, d.ht[1].capacity)
	require.NotNil(t, d.ht[1].node[hashValue].Search(word))

	d.MigrateData()
	require.Equal(t, hash("hello", d.ht[0].capacity), d.rehashIndex)
	require.Equal(t, 1, d.ht[0].used)
	require.Equal(t, 2, d.ht[1].used)

	d.MigrateData()
	require.Equal(t, -1, d.rehashIndex)
	require.Equal(t, 3, d.ht[0].used)
	require.Nil(t, d.ht[1])
}

func TestWordSpellcheck(t *testing.T) {
	d := New(0)

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
		d.Insert(word)
	}

	fi.Seek(0, 0)

	// word not in hash table can not be found
	require.Equalf(t, false, d.Search("fafd"), "fafd")
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

		require.Equalf(t, true, d.Search(word), word)
	}
}
