package asyncop

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAsyncOp(t *testing.T) {
	t.Parallel()
	ao := New(1)
	err := ao.Execute(func() error {
		fmt.Println("op")
		time.Sleep(1 * time.Second)
		return nil
	})
	require.NoError(t, err)

	err = ao.Execute(func() error {
		fmt.Println("op1")
		return nil
	})
	require.EqualError(t, err, "no more operation is allowed")

	fmt.Println("waiting...")
	time.Sleep(3 * time.Second)
}
