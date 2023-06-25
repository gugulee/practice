package main

import (
	"fmt"
)

func producer(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func consumer(id int, ch <-chan int, done chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("id: %d, recv: %d\n", id, v)
	}

	done <- struct{}{}
}

func main() {
	ch := make(chan int)
	n := 2
	done := make(chan struct{}, n)
	
	for i := 1; i <= n; i++ {
		go consumer(i, ch, done)
	}
	
	go producer(ch)

	for i := 1; i <= n; i++ {
		<-done
	}
}
