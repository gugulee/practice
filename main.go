package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func main() {
	now := time.Now()
	wait.PollImmediate(1*time.Second, 4*time.Second, func() (done bool, err error) {

		fmt.Println("in")
		// time.Sleep(10 * time.Second)

		return false, nil
	})

	fmt.Println("elapse", time.Since(now))
}
