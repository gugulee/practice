package variableparameters

import (
	"fmt"
	"testing"
)

func Test_VariableParameters(t *testing.T) {
	options := NewOptions(WithIdOption(10), WithNameOption("lee"), WithaddressOption("cq"))
	fmt.Println(options)
}
