package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func add() error {
	return fmt.Errorf("test error")
}

func test() error {
	err := add()
	return errors.Wrap(err, "call test")
}

func main() {
	if err := test(); err != nil {
		fmt.Println(err)
	}
}
