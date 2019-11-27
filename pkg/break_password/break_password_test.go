package break_password

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBreakPassword(t *testing.T) {
	out := breakPassword("jgfdgh", "")
	assert.Equal(t, true, out)
	out = breakPassword("zzzz", "")
	assert.Equal(t, true, out)

	out = breakPassword("1zzz", "")
	assert.Equal(t, false, out)
}
