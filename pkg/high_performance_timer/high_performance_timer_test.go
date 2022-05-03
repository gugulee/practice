package highperformancetimer

import (
	"testing"
	"time"
)

func generateTestData() (result map[int64]string) {
	result = make(map[int64]string)
	for i := 1; i <= 5; i++ {
		a := time.Now().Unix() + int64(i*5)
		s := time.Unix(a, 0).Format("2006-01-02 15:04:05")
		result[a] = "the task will be run at " + s
	}

	return result
}

func Test_highPerformanceTimer(t *testing.T) {
	in := generateTestData()
	h:=New(in)
	h.highPerformanceTimer()
}
