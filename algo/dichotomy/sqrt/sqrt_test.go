package sqrt

import "testing"

const (
	deltaThreshold = 0.0000001
	maxTry         = 100000
)

func TestAchieveSqrt(t *testing.T) {
	var source float64 = 2
	out := AchieveSqrt(source, deltaThreshold, maxTry)
	t.Logf("in=%f,out=%f", source, out)

	source = 99.191293
	out = AchieveSqrt(source, deltaThreshold, maxTry)
	t.Logf("in=%f,out=%f", source, out)
}
