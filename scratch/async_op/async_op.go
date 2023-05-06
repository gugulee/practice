package asyncop

import (
	"fmt"
)

type AsyncOp struct {
	inProgress chan struct{}
}

func New(maxConcurrent int) *AsyncOp {
	return &AsyncOp{
		inProgress: make(chan struct{}, maxConcurrent),
	}
}

func (ao *AsyncOp) Execute(op func() error) error {
	select {
	case ao.inProgress <- struct{}{}:
	default:
		return fmt.Errorf("no more operation is allowed")
	}

	go func() {
		defer func() { <-ao.inProgress }()

		if err := op(); err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}()

	return nil
}
