package removek

import "github.com/practice/pkg/tools"

// removeK remove k digit from a
func removeK(a string, k int) string {
	defer tools.FuncTime()()
	tmp := 0

	for i := 0; i < k; i++ {
		for j := tmp; j < len(a); {
			// it is the end of a, remove the last character
			if j == len(a)-1 {
				a = a[:j]
				tmp = j - 1
				break
			}

			if a[j] > a[j+1] {
				a = a[:j] + a[j+1:]
				tmp = 0
				break
			}

			j++
		}

	}

	// remove the prefix 0
	for i := 0; i < len(a); {
		if '0' != a[i] {
			break
		}

		a = a[i+1:]
	}

	if 0 == len(a) {
		a = "0"
	}

	return a
}

func removeKdigits(num string, k int) string {
	digits := len(num) - k
	stack := make([]byte, len(num))
	top := 0
	for i := range num {
		for top > 0 && stack[top-1] > num[i] && k > 0 {
			top--
			k--
		}
		stack[top] = num[i]
		top++
	}

	i := 0

	for i < digits && stack[i] == '0' {
		i++
	}

	if i == digits {
		return "0"
	}

	return string(stack[i:digits])
}
