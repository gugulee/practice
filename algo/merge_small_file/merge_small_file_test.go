package mergesmallfile

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func generateTestData() (data [][]int) {
	data = make([][]int, 5)
	data[0] = []int{1, 11, 21, 31, 41, 51, 61, 71}
	data[1] = []int{2, 12, 22, 32, 42, 52, 62, 72, 82}
	data[2] = []int{3, 13, 23, 33, 43, 53, 63, 73, 83, 93, 103}
	data[3] = []int{4, 14, 24, 34, 44, 54, 64, 74, 84, 94}
	data[4] = []int{5, 15, 25, 35, 45, 55, 65, 75, 85, 95, 105, 115}
	return data
}

func Test_mergeSmallFile(t *testing.T) {
	req := require.New(t)
	data := [][]int{
		{1, 2, 4},
		{3, 5, 6},
	}
	out := mergeSmallFile(data)
	req.Equal([]int{1, 2, 3, 4, 5, 6}, out)

	data = generateTestData()
	var expect []int
	for _, d := range data {
		expect = append(expect, d...)
	}

	sort.Ints(expect)

	out = mergeSmallFile(data)
	req.Equal(expect, out)
}

func Test_mergeSmallFile1(t *testing.T) {
	req := require.New(t)

	data := [][]int{
		{3, 5, 6},
		{1, 2, 4},
	}
	out := mergeSmallFile1(data)
	req.Equal([]int{1, 2, 3, 4, 5, 6}, out)

	data = generateTestData()
	var expect []int
	for _, d := range data {
		expect = append(expect, d...)
	}

	sort.Ints(expect)

	out = mergeSmallFile1(data)
	req.Equal(expect, out)
}
