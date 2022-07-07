package change

import (
	"fmt"
	"strings"
)

var value = []int{1, 2, 5, 10, 20, 50, 100}
var moneyNumber = map[int]int{
	1:   100,
	2:   40,
	5:   20,
	10:  10,
	20:  7,
	50:  3,
	100: 3,
}

func change(total int) string {
	var r []string
	for i := len(value) - 1; i >= 0; i-- {
		j := 0
		for ; j < moneyNumber[value[i]] && total >= value[i]; j++ {
			total -= value[i]
		}

		if 0 != j {
			r = append(r, fmt.Sprintf("%d * %d", value[i], j))
		}
	}
	return strings.Join(r, " + ")
}
