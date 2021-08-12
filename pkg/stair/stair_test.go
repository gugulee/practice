package stair

import (
	"fmt"
	"testing"
)

func TestStair(t *testing.T) {
	stair(1, []int{})
	fmt.Println("----------------------------")

	stair(2, []int{})
	fmt.Println("----------------------------")

	stair(3, []int{})
	fmt.Println("----------------------------")

	stair(4, []int{})
	fmt.Println("----------------------------")

	stair(5, []int{})
	fmt.Println("----------------------------")
}

func TestCountStair(t *testing.T) {
	fmt.Println(countStair(1))
	fmt.Println("----------------------------")

	fmt.Println(countStair(2))
	fmt.Println("----------------------------")

	fmt.Println(countStair(3))
	fmt.Println("----------------------------")

	fmt.Println(countStair(4))
	fmt.Println("----------------------------")

	fmt.Println(countStair(5))
	fmt.Println("----------------------------")
}
