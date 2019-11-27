package conbine

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	in := []string{"t1", "t2", "t3", "t4"}
	combine(in, []string{}, 2)

	fmt.Println("---------------------")

	in = []string{"t1", "t2", "t3"}
	combine(in, []string{}, 2)
}
