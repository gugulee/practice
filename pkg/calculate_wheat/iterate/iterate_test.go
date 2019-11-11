package iterate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

func traceFunc(f func(), times int) {
	t1 := time.Now() // get current time
	for i := 0; i < times; i++ {
		f()
	}
	fmt.Println("func elapsed: ", time.Since(t1))
}

func TestIterate(t *testing.T) {
	out := Iterate(-1)
	assert.Equal(t, out, uint(0))

	out = Iterate(0)
	assert.Equal(t, out, uint(0))

	grid := 63
	traceFunc(func() {
		out := Iterate(grid)
		assert.Equal(t, out, uint(0x7fffffffffffffff))
	}, 1)
}

func TestIterate1(t *testing.T) {
	grid := 63
	result := math.Pow(2, float64(grid)) - 1
	traceFunc(func() {
		out, _, correct := Iterate1(grid)
		assert.True(t, correct)
		assert.Equal(t, out, result)
	}, 1)
}
