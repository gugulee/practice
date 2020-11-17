package kmp

// kmp search pattern in main with kmp algorithm
func kmp(main string, pattern string) bool {
	m := len(main)
	p := len(pattern)

	//defensive
	if m == 0 || p == 0 || m < p {
		return false
	}

	next := getNext(pattern)

	//	    main := "ababaeababacd"
	//  patternp := "ababacd"
	for i := 0; i <= m-p; {
		j := 0
		for ; j < p; j++ {
			if main[i+j] != pattern[j] {
				break
			}
		}

		// found
		if j == p {
			return true
		}

		if 0 == j {
			i++
		} else {
			i += j - 1 - next[j-1]
		}
	}

	return false
}

func getNext(pattern string) (next []int) {
	m := len(pattern)
	next = make([]int, m-1)

	for i := range next {
		next[i] = -1
	}

	// aba bacdd
	for i := 0; i < m-1; i++ {
		k := 0
		j := 1

		for j <= i {
			if pattern[k] == pattern[j] {
				next[i] = k
				k++
			} else {
				k = 0
			}

			j++
		}

		if 0 == k {
			next[i] = -1
		}
	}

	return next
}
