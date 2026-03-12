package change

import (
	"fmt"
	"testing"
)

func Test_change(t *testing.T) {
	total := 577
	out := change(total)
	fmt.Println(out)

	total = 9
	out = change(total)
	fmt.Println(out)
}
