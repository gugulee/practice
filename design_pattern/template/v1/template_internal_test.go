package template

import (
	"fmt"
	"testing"
)

func Test_Runner(t *testing.T) {
	l := NewLogger()
	l.Start()
	fmt.Println("Hello, playground")
}
