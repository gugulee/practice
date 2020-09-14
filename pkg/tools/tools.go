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
