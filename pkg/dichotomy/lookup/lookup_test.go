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
