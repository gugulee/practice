package builder_test

import (
	"fmt"
	"testing"

	"github.com/practice/design_pattern/builder"
)

func Test_New(t *testing.T) {
	r := builder.NewBuilder().
		SetName("lee").
		SetMaxTotal(100).
		SetMaxIdle(100).
		SetMinIdle(20).
		Build()

	fmt.Println(r)
}
