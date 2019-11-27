package conbine

import "fmt"

func combine(origin, result []string, m int) {
	if 0 == m {
		fmt.Println(result)
		return
	}

	for i := range origin {
		var newResult []string
		newResult = append(newResult, result...) // copy result
		newResult = append(newResult, origin[i])

		combine(origin[i+1:], newResult, m-1)
	}
}

