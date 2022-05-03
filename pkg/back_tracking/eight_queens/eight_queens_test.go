package eightqueens

import (
	"fmt"
	"testing"
)

func Test_eightQueens(t *testing.T) {
	in := make([][]int, 8)
	for i := 0; i < 8; i++ {
		in[i] = make([]int, 8)
	}

	eightQueens(in, 0)
	fmt.Println(count)
}
