package main

import (
	"fmt"
	"io/ioutil"
	"runtime"

	"golang.org/x/sys/unix"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	pid := unix.Getpid()
	fmt.Println("pid:", pid)
	fmt.Println("total cpu:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())

	for {
		_, err := ioutil.ReadFile("/tmp/1")
		if nil != err {
			fmt.Println("readFile: ", err)
		}
		tid := unix.Gettid()
		fmt.Println("tid:", tid)
		// time.Sleep(1 * time.Second)
	}
}
