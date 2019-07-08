package sqrt

import "testing"

const (
	deltaThreshold = 0.0001
	maxTry         = 100000
)

func TestAchieveSqrt(t *testing.T) {
	var source float64 = 10
	out := AchieveSqrt(source, deltaThreshold, maxTry)
	t.Logf("out=%f", out)

	source  = 99
	out = AchieveSqrt(source, deltaThreshold, maxTry)
	t.Logf("out=%f", out)

}
