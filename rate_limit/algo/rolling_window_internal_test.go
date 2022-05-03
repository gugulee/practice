package algo

import (
	"testing"
	"time"

	"github.com/practice/rate_limit/config"
	"github.com/stretchr/testify/require"
)

func Test_Limit(t *testing.T) {
	limit := &config.Limit{
		Limit:    5,
		Interval: 1,
	}
	rw := NewRollingWindow(limit)
	var pass bool
	for i := 0; i < 5; i++ {
		pass = rw.Limit()
		require.True(t, pass)
	}

	pass = rw.Limit()
	require.False(t, pass)

	time.Sleep(1 * time.Second)

	for i := 0; i < 5; i++ {
		out := rw.Limit()
		require.True(t, out)
	}

	pass = rw.Limit()
	require.False(t, pass)
}
