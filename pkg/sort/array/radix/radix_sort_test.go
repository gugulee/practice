package radix

import (
	"math"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/practice/pkg/sort/array/quick"
	"github.com/practice/pkg/tools"
	"github.com/stretchr/testify/require"
)

var digits = 4

// generateNum generate n number which has XX digits
func generateNum(n int) (r []int) {
	rand.Seed(time.Now().Unix())
	r = make([]int, n)
	for i := 0; i < n; i++ {
		r[i] += 1 * int(math.Pow10(digits-1))
		for j := digits - 2; j >= 0; j-- {
			r[i] += rand.Intn(10) * int(math.Pow10(j))
		}
	}

	return
}

func TestRadixSort(t *testing.T) {
	length := 1000000
	in := generateNum(length)
	standard := make([]int, length)
	in1 := make([]int, length)
	copy(standard, in)
	copy(in1, in)

	// fmt.Println("in", in)
	// fmt.Println()

	tools.FuncTime1(func() { quick.Sort(standard) }, "quick sort")

	tools.FuncTime1(func() { sort.Ints(in) }, "Sort")

	// tools.FuncTime1(func() { radixSort(in) }, "radix sort")

	tools.FuncTime1(func() { radixSort1(in1) }, "radix1 sort")


	require.Equal(t, standard, in)
	require.Equal(t, standard, in1)
}

func BenchmarkRadixSort(b *testing.B) {
	length := 10000000
	in := generateNum(length)
	in1 := make([]int, length)
	copy(in1, in)

	tools.FuncTime1(func() { radixSort(in) }, "radix sort")
	tools.FuncTime1(func() { radixSort1(in1) }, "radix1 sort")
}
