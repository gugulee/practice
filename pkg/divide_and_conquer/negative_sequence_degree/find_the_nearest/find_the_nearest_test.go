package findthenearest

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func generateTestData(n int) (result [][]float64) {
	rand.Seed(time.Now().Unix())
	result = make([][]float64, n)

	for i := 0; i < n; i++ {
		result[i] = make([]float64, 2)
		result[i][0] = float64(rand.Intn(1000))
		result[i][1] = float64(rand.Intn(500))
	}

	return result
}

func Test_xSort(t *testing.T) {
	in := [][]float64{
		{7, 7},
		{10, 2},
		{30, 10},
		{1, 3},
		{4, 10},
		{5, 6},
	}

	sort.Sort(xSort(in))
	require.Equal(t, [][]float64{
		{1, 3},
		{4, 10},
		{5, 6},
		{7, 7},
		{10, 2},
		{30, 10},
	}, in)
}

func Test_ySort(t *testing.T) {
	in := [][]float64{
		{7, 7},
		{10, 2},
		{30, 10},
		{1, 3},
		{4, 10},
		{5, 6},
	}

	sort.Sort(ySort(in))
	require.Equal(t, [][]float64{
		{10, 2},
		{1, 3},
		{5, 6},
		{7, 7},
		{30, 10},
		{4, 10},
	}, in)
}

func Test_findTheNearest(t *testing.T) {
	in := [][]float64{
		{7, 7},
		{10, 2},
		{30, 10},
		{1, 3},
		{4, 10},
		{5, 6},
	}
	a, b, d := findTheNearest(in)
	require.Equal(t, []float64{5, 6}, a)
	require.Equal(t, []float64{7, 7}, b)
	require.Equal(t, 2.236, d)

	in = generateTestData(10)
	a1, b1, d1 := bruteForce(in)
	a2, b2, d2 := findTheNearest(in)
	fmt.Println(a1, b1, d1)
	fmt.Println(a2, b2, d2)
	require.Equal(t, d1, d2)
}

func Test_bruteForce(t *testing.T) {
	in := [][]float64{
		{1, 3},
		{4, 10},
		{5, 6},
		{7, 7},
		{10, 2},
		{30, 10},
	}
	a, b, d := bruteForce(in)
	require.Equal(t, []float64{5, 6}, a)
	require.Equal(t, []float64{7, 7}, b)
	require.Equal(t, 2.236, d)
}

func Test_bsearch(t *testing.T) {
	t.Run("search x coordinate", func(t *testing.T) {
		in := [][]float64{
			{1, 3},
			{4, 10},
			{5, 6},
			{5, 10},
			{7, 7},
			{10, 2},
			{30, 10},
		}

		out := bsearch(in, 5, 0)
		require.Equal(t, 4, out)

		out = bsearch(in, 15, 0)
		require.Equal(t, 6, out)

		out = bsearch(in, 30, 0)
		require.Equal(t, -1, out)

		out = bsearch(in, 0.5, 0)
		require.Equal(t, 0, out)
	})

	t.Run("search y coordinate", func(t *testing.T) {
		in := [][]float64{
			{10, 2},
			{1, 3},
			{5, 6},
			{7, 7},
			{4, 10},
			{30, 10},
		}

		out := bsearch(in, 2, 1)
		require.Equal(t, 1, out)

		out = bsearch(in, 10, 1)
		require.Equal(t, -1, out)

		out = bsearch(in, 7, 1)
		require.Equal(t, 4, out)

		out = bsearch(in, 30, 1)
		require.Equal(t, -1, out)

		out = bsearch(in, 1, 1)
		require.Equal(t, 0, out)

		out = bsearch(in, 6, 1)
		require.Equal(t, 3, out)

		out = bsearch(in, 0.5, 1)
		require.Equal(t, 0, out)
	})

	t.Run("abnormal scene", func(t *testing.T) {
		in := [][]float64{}
		out := bsearch(in, 5, 0)
		require.Equal(t, -1, out)

		in = [][]float64{{5, 6}}
		out = bsearch(in, 2.38, 1)
		require.Equal(t, 0, out)
	})

}
