package findthenearest

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

type xSort [][]float64

func (s xSort) Len() int {
	return len(s)
}

func (s xSort) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s xSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ySort [][]float64

func (s ySort) Len() int {
	return len(s)
}

func (s ySort) Less(i, j int) bool {
	return s[i][1] < s[j][1]
}

func (s ySort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// findthenearest return the index of the nearest point between two points
func findTheNearest(in [][]float64) ([]float64, []float64, float64) {
	sort.Sort(xSort(in))

	return nearestDistance(in)
}

//	in := [][]float64{
//		{1, 3},
//		{4, 10},
//		{5, 6},
//		{7, 7},
//		{10, 2},
//		{30, 10},
//	}
func nearestDistance(in [][]float64) (a []float64, b []float64, dmin float64) {
	if 2 == len(in) {
		return in[0], in[1], distance(in[0], in[1])
	} else if len(in) < 2 {
		return nil, nil, math.MaxFloat64
	}

	p := 0
	r := len(in) - 1
	// q is the middle point
	q := (p + r) / 2

	left := make([][]float64, q-p+1)
	sort.Sort(xSort(in))
	copy(left, in[p:q+1])

	right := make([][]float64, r-q)
	copy(right, in[q+1:r+1])
	sort.Sort(ySort(right))

	// calculate the distance in the left side
	a1, b1, d1 := nearestDistance(left)
	// calculate the distance in the right side
	a2, b2, d2 := nearestDistance(right)

	a, b, dmin = a1, b1, d1
	if d1 > d2 {
		a, b, dmin = a2, b2, d2
	}

	/* consider points in the left side which the distance is less than dmin */
	xidx := bsearch(left, in[q][0]-dmin, 0)
	yidx := bsearch(right, in[q][1]-dmin, 1)

	if xidx < 0 || yidx < 0 {
		return a, b, dmin
	}

	for i := xidx; i < len(left); i++ {
		// there are only 6 point at most in the right side of in
		for j := yidx; j < 7 && j < len(right); j++ {
			d3 := distance(left[i], right[j])
			if d3 < dmin {
				a, b, dmin = left[i], right[j], d3
			}
		}
	}

	return a, b, Decimal(dmin)
}

// bruteForce return the index of the nearest point between two points
func bruteForce(in [][]float64) ([]float64, []float64, float64) {
	a, b := -1, -1
	tmp := math.MaxFloat64
	length := len(in)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			d := distance(in[i], in[j])
			if d < tmp {
				tmp = d
				a, b = i, j
			}
		}
	}

	return in[a], in[b], Decimal(tmp)
}

func distance(x, y []float64) float64 {
	return math.Sqrt(math.Pow(x[0]-y[0], 2) + math.Pow(x[1]-y[1], 2))
}

// bsearch lookup the first value which is greater than value in a, return the index of target
func bsearch(a [][]float64, value float64, dimension int) int {
	length := len(a)

	if 0 == length {
		return -1
	}

	if dimension >= len(a[0]) {
		return -1
	}

	low, high := 0, length-1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if value >= a[mid][dimension] {
			if mid == length-1 {
				return -1
			}

			if a[mid+1][dimension] > value {
				return mid + 1
			}

			low = mid + 1
		} else {
			if 0 == mid || a[mid-1][dimension] < value {
				return mid
			}
			high = mid - 1
		}
	}

	return -1
}

// Decimal return decimal value
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", value), 64)
	return value
}
