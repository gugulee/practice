package get_min_money

import (
	"fmt"
	"testing"
)

func TestGetMinMoney(t *testing.T) {
	out := getMinMoney(50)
	fmt.Println(out)

}

func TestGetMinMoney1(t *testing.T) {
	out := getMinMoney1(50)
	fmt.Println(out)
}

func TestGetMinMoney2(t *testing.T) {
	out := getMinMoney2(50)
	fmt.Println(out)
}
