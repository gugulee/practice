package horse_race

import (
	"fmt"
)

// the total time of a horse complete
type horseTime float64

var (
	// the time of every horse of q
	qHorseTime = map[string]horseTime{"q1": 1.0, "q2": 2.0, "q3": 3.0}
	// the time of every horse of t
	tHorseTime  = map[string]horseTime{"t1": 1.5, "t2": 2.5, "t3": 3.5}
	qHorseOrder = []string{"q1", "q2", "q3"} // the horse order of q
)

func compare(t, q []string) {
	// t won count
	wonCount := 0
	for i := range t {
		if tHorseTime[t[i]] < qHorseTime[q[i]] {
			wonCount++
		}
	}
	if len(t)/2 < wonCount {
		fmt.Println("t won!")
	} else {
		fmt.Println("q won!")
	}
}

func permutate(horses []string, result []string) {
	if 0 == len(horses) {
		fmt.Printf("result is %s\n", result)
		compare(result, qHorseOrder)
		fmt.Println()
	}

	for i, horse := range horses {
		var restHorses []string
		restHorses = append(restHorses, horses[0:i]...)
		restHorses = append(restHorses, horses[i+1:]...)

		permutate(restHorses, append(result, horse))
	}

	// for i, horse := range horses {
	// 	var newResult []string
	// 	newResult = append(newResult, result...) // copy result
	// 	newResult = append(newResult, horse)

	// 	var restHorses []string
	// 	restHorses = append(restHorses, horses[0:i]...)
	// 	restHorses = append(restHorses, horses[i+1:]...)

	// 	permutate(restHorses, newResult)
	// }

}
