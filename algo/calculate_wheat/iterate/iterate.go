package iterate

import (
	"fmt"
	"math"
	"time"
)

func trace() func() {
	t1 := time.Now() // get current time
	return func() {
		fmt.Println("func elapsed: ", time.Since(t1).Nanoseconds())
	}
}

// Iterate calculate the number of wheat by iteration
func Iterate(grid int) uint {
	if grid <= 0 {
		return 0
	}

	var sum uint = 0
	var numberOfGridWheat uint = 0

	// the first grid
	sum += 1
	numberOfGridWheat = 1

	for i := 2; i <= grid; i++ {
		numberOfGridWheat *= 2
		sum += numberOfGridWheat
	}

	return sum
}

// Iterate calculate the number of wheat by iteration
func Iterate1(grid int) (float64, float64, bool) {
	if grid <= 0 {
		return 0, 0, false
	}

	var sum float64 = 0
	var numberOfGridWheat float64 = 0
	correct := false

	if 1 == grid {
		if 1 == math.Pow(2, 1)-1 {
			sum = 1
			numberOfGridWheat = 1
			correct = true
			return sum, numberOfGridWheat, correct
		}
	}

	sum, numberOfGridWheat, correct = Iterate1(grid - 1)
	if !correct {
		return 0, 0, false
	}

	numberOfGridWheat *= 2
	sum += numberOfGridWheat

	if sum == math.Pow(2, float64(grid))-1 && correct {
		return sum, numberOfGridWheat, correct
	}

	return 0, 0, false
}
