package sqrt

import (
	"math"
)

// AchieveSqrt achieve square root  of source
// deltaThreshold is threshold of precision
// maxTry prevent program from endless Loop
func AchieveSqrt(source, deltaThreshold float64, maxTry int) float64 {
	var left, right, middle float64 = 1, source, 0
	for i := 0; i < maxTry; i++ {
		middle := left + (right-left)/2
		middleSquare := middle * middle
		if delta := math.Abs(middleSquare/source - 1); delta <= deltaThreshold {
			break
		}

		if source < middleSquare {
			right = middle
		} else {
			left = middle
		}
	}
	return middle
}
