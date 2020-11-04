package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	for i := 0; i < 10; i++ {
		targetVertex := rand.Intn(10)
		fmt.Println(targetVertex)
	}
}
