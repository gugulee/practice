package tools

import (
	"fmt"
	"runtime"
	"time"
)

func FuncTime() func() {
	t1 := time.Now() // get current time
	f, _, _, _ := runtime.Caller(1)
	return func() {
		elapsed := time.Since(t1)
		fmt.Printf("function %q elapsed: %v\n", runtime.FuncForPC(f).Name(), elapsed)
	}
}

func FuncTime1(f func(), name string) {
	t1 := time.Now() // get current time
	f()
	fmt.Printf("function %s elapsed: %v\n", name, time.Since(t1))
}

func Remove(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}
