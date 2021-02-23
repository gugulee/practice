package tools

import (
	"fmt"
	"runtime"
	"time"
)

// FuncTime ...
func FuncTime() func() {
	t1 := time.Now() // get current time
	f, _, _, _ := runtime.Caller(1)
	return func() {
		elapsed := time.Since(t1)
		fmt.Printf("function %q elapsed: %v\n", runtime.FuncForPC(f).Name(), elapsed)
	}
}

// FuncTime1 ...
func FuncTime1(f func(), name string) {
	t1 := time.Now() // get current time
	f()
	fmt.Printf("function %s elapsed: %v\n", name, time.Since(t1))
}

// Remove remove one element from slice
func Remove(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

// Reverse reverse slice
func Reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// MapDeepCopy copy a to b
func MapDeepCopy(b, a map[int]int) {
	if nil == a || b == nil {
		return
	}

	for k := range a {
		b[k] = a[k]
	}
}

// GetMinIdx return the index of the minimum
func GetMinIdx(a []int) int {
	length := len(a)
	if 0 == length {
		return -1
	}

	minIdx := -1
	// find the minimum index which a[minIdx]>0
	for i := 0; i < length; i++ {
		if a[i] > 0 {
			minIdx = i
			break
		}
	}

	if -1 == minIdx {
		return -1
	}

	for i := minIdx + 1; i < length; i++ {
		// neglect the negative
		if a[i] < 0 {
			continue
		}

		if a[minIdx] > a[i] {
			minIdx = i
		}
	}

	return minIdx
}

// Min get the minimum
func Min(a ...int) int {
	if 0 == len(a) {
		return -1
	}

	min := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}
	return min
}

// Max get the maximum
func Max(a ...int) int {
	if 0 == len(a) {
		return -1
	}
	max := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}
