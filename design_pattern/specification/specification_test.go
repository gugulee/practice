package specification

import (
	"fmt"
	"testing"
)

func Test_Rule(t *testing.T) {
	rule := NewAndSpecification(
		NewKV("ip4.src", "==", "$test.allow.as", ""),
		NewKV("ip4.src", "!=", "$test.except.as", ""),
		NewKV("tcp.dst", "<=", "12345", "12500"),
		NewKV("outport", "==", "@ovn.sg.test_sg", ""),
		NewKV("ip", "", "", ""),
	)

	fmt.Println(rule.Rule())
}
