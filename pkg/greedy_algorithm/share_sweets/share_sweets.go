package sharesweets

import (
	"fmt"

	"github.com/practice/pkg/tools"
)

func shareSweets(sweetsSize, childrenNeed []int) {
	for i := range childrenNeed {
		j := 0
		for j < len(sweetsSize) {
			if childrenNeed[i] <= sweetsSize[j] {
				break
			}
			j++
		}

		if j == len(sweetsSize) {
			fmt.Printf("child: %d's need(%d) are not met\n", i, childrenNeed[i])
		} else {
			fmt.Printf("child: %d's need(%d) are met, sweet's size is %d\n", i, childrenNeed[i], sweetsSize[j])
			sweetsSize = tools.Remove(sweetsSize, j)
		}
	}
}
