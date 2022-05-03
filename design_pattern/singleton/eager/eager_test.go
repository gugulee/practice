package eager_test

import (
	"sync"
	"testing"

	"github.com/practice/design_pattern/singleton/eager"
	"github.com/stretchr/testify/require"
)

var lock sync.Mutex

func Test_eager(t *testing.T) {
	count := 200
	expect := make([]uint64, count)
	for i := 0; i < count; i++ {
		expect[i] = uint64(i + 1)
	}

	// run it many times, be sure to get id which is correct
	for j := 0; j < 10000; j++ {
		eager.Instance().Clear()
		ids := make([]uint64, 0, count)

		var wg sync.WaitGroup
		wg.Add(count)
		for i := 0; i < count; i++ {
			go func() {
				defer wg.Done()
				id := eager.Instance().GetId()
				lock.Lock()
				ids = append(ids, id)
				lock.Unlock()
			}()
		}
		wg.Wait()
		require.ElementsMatch(t, expect, ids)
	}
}
