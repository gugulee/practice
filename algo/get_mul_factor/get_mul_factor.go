package get_mul_factor

import "fmt"

func contains(s []int, t int) bool {
	for _, k := range s {
		if k == t {
			return true
		}
	}

	return false
}

func equal(s []int, t int) bool {
	product := 1
	for _, k := range s {
		product *= k
	}
	return t == product
}

func GetMulFactot(original, num int, result []int) {
	for i := 1; i <= num; i++ {
		if contains(result, 1) && 1 == i {
			if equal(result, original) {
				fmt.Println(result)
			}
			continue
		}
		var newResult []int
		newResult = append(newResult, result...) // copy result
		if 0 == num%i {
			newResult = append(newResult, i)
			GetMulFactot(original, num/i, newResult)
		}
	}
}

func GetMulFactot1(num int, result []int) {
	if 1 == num {
		if !contains(result, 1) {
			result = append(result, 1)
		}
		fmt.Println(result)
		return
	}

	for i := 1; i <= num; i++ {
		if contains(result, 1) && 1 == i {
			continue
		}

		var newResult []int
		newResult = append(newResult, result...) // copy result
		if 0 == num%i {
			newResult = append(newResult, i)
			GetMulFactot1(num/i, newResult)
		}
	}
}
