package lottery

import (
	"fmt"
	"testing"
)

func TestDecrease(t *testing.T) {
	in := []int{3, 2, 1}
	decrease(in)
	fmt.Println(in)

}

func TestDelElement(t *testing.T) {
	in := []string{"3", "2", "1"}
	out := delElement(in, "1")
	fmt.Println(out)

	in = []string{"3", "2", "1"}
	out = delElement(in, "4")
	fmt.Println(out)
}

func TestLottery(t *testing.T) {
	in := []string{"m0", "m1", "m2", "m3", "m4"}
	lottery(in, in, []string{}, []int{2, 1, 1}) //output: 60
	fmt.Println("-------------------------")
	in = []string{"m0", "m1", "m2", "m3", "m4", "m5", "m6", "m6", "m8", "m9"}
	lottery(in, in, []string{}, []int{3, 2, 1}) //output: 12600
}