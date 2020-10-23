package lookup

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookup(t *testing.T) {
	dict := []string{"i", "am", "a", "programer", "of", "c", "and", "go"}
	sort.Strings(dict)

	out, err := Lookup(dict, "go")
	assert.NoError(t, err)
	assert.Equal(t, true, out)

	out, err = Lookup(dict, "cc")
	assert.EqualError(t, err, `can no find target:"cc" in dict`)
	assert.Equal(t, false, out)

	out, err = Lookup([]string{}, "i")
	assert.EqualError(t, err, `dict can not be empty`)

	out, err = Lookup(dict, "")
	assert.EqualError(t, err, `target can not be empty`)
}

func TestLookup1(t *testing.T) {
	dict := []string{"i", "am", "a", "programer", "of", "c", "and", "go"}
	sort.Strings(dict)
	fmt.Println("dict", dict)

	out := Lookup1(dict, "go", 0, len(dict)-1)
	assert.Equal(t, true, out)

	out = Lookup1(dict, "cc", 0, len(dict)-1)
	assert.Equal(t, false, out)
}

func TestBsearch(t *testing.T) {
	a := []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}
	out := Bsearch(a, 8)
	assert.Equal(t, 5, out)

	a = []int{1, 3, 4, 5, 8, 8, 8, 8, 11, 18}
	out = Bsearch(a, 8)
	assert.Equal(t, 4, out)

	a = []int{1, 3, 4, 5, 8, 8, 8, 8, 11, 18}
	out = Bsearch(a, 2)
	assert.Equal(t, -1, out)

	a = []int{1, 3, 4, 5, 6, 7, 8, 8, 11, 18}
	out = Bsearch(a, 8)
	assert.Equal(t, 6, out)

	a = []int{1, 2, 2, 3, 4, 6, 7, 8, 11, 18}
	out = Bsearch(a, 2)
	assert.Equal(t, 1, out)

	a = []int{2, 2, 2, 3, 4, 6, 7, 8, 11, 18}
	out = Bsearch(a, 2)
	assert.Equal(t, 0, out)
}

func TestBsearch1(t *testing.T) {
	a := []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}
	out := bsearch1(a, 8)
	assert.Equal(t, 7, out)

	a = []int{1, 3, 4, 5, 6, 7, 8, 8, 11, 18}
	out = bsearch1(a, 8)
	assert.Equal(t, 7, out)

	a = []int{1, 2, 2, 3, 4, 6, 7, 8, 11, 18}
	out = bsearch1(a, 2)
	assert.Equal(t, 2, out)

	a = []int{2, 2, 2, 3, 4, 6, 7, 8, 11, 11}
	out = bsearch1(a, 11)
	assert.Equal(t, 9, out)
}

func TestBsearch2(t *testing.T) {
	a := []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}
	out := bsearch2(a, 8)
	assert.Equal(t, 5, out)

	a = []int{2, 3, 4, 5, 6, 8, 8, 8, 11, 18}
	out = bsearch2(a, 1)
	assert.Equal(t, 0, out)
}

func TestBsearch3(t *testing.T) {
	a := []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}
	out := bsearch3(a, 8)
	assert.Equal(t, 7, out)

	a = []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}
	out = bsearch3(a, 20)
	assert.Equal(t, 9, out)
}

func TestBsearch4(t *testing.T) {
	a := []int{4, 5, 6, 1, 2, 3}
	out := bsearch4(a, 4)
	assert.Equal(t, 0, out)

	out = bsearch4(a, 5)
	assert.Equal(t, 1, out)

	out = bsearch4(a, 6)
	assert.Equal(t, 2, out)

	out = bsearch4(a, 2)
	assert.Equal(t, 4, out)

	a = []int{10, 15, 1, 3, 5, 7}
	out = bsearch4(a, 10)
	assert.Equal(t, 0, out)

	out = bsearch4(a, 15)
	assert.Equal(t, 1, out)

	out = bsearch4(a, 1)
	assert.Equal(t, 2, out)

	out = bsearch4(a, 7)
	assert.Equal(t, 5, out)
}
