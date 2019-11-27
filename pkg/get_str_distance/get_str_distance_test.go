package get_str_distance

import (
	"fmt"
	"testing"
)

func TestGetStrDistance(t *testing.T) {
	out := getStrDistance("mouuse", "mouse")
	fmt.Println("distance:",out)
}

func TestGetStrDistance1(t *testing.T) {
	a := "mouuse"
	b := "mouse"
	out := getStrDistance1(a, b, len(a), len(b))
	fmt.Println(out)
}
