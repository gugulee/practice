package combine

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	in := []string{"t1", "t2", "t3", "t4", "t5", "t6"}
	combine(in, []string{}, 3)

	fmt.Println("---------------------")

	in = []string{"t1", "t2", "t3"}
	combine(in, []string{}, 2)
}
