package get_mul_factor

import (
	"fmt"
	"testing"
)

func TestIterate(t *testing.T) {
	GetMulFactot(4, 4, []int{})
	fmt.Println("----------------------------")
	GetMulFactot(6, 6, []int{})
	fmt.Println("----------------------------")
	GetMulFactot(8, 8, []int{})
}

func TestIterate1(t *testing.T) {
	GetMulFactot1(4, []int{})
	fmt.Println("----------------------------")
	GetMulFactot1(6, []int{})
	fmt.Println("----------------------------")
	GetMulFactot1(8, []int{})
}
