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
	// []int{-1, -1, 0, 1, 2, -1},
	j := 0
	for i := 0; i < m; i++ {
		for j > 0 && main[i] != pattern[j] { // 一直找到a[i]和b[j]
			j = next[j-1] + 1
		}

		if main[i] == pattern[j] {
			j++
		}

		if j == p { // 找到匹配模式串的了
			return true
		}
	}

	return false
}

// kmp search pattern in main with kmp algorithm
func kmp1(main string, pattern string) bool {
	m := len(main)
	p := len(pattern)

	//defensive
	if m == 0 || p == 0 || m < p {
		return false
	}

	next := getNext1(pattern)

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
	next = make([]int, m)
	next[0] = -1

	// pattern = "bcaebcaedbc"
	// []int{-1, -1, -1, -1, 0, 1, 2, 3, -1, 0, 1}
	k := -1
	for i := 1; i < m; i++ {
		for k != -1 && pattern[k+1] != pattern[i] {
			k = next[k]
		}

		if pattern[k+1] == pattern[i] {
			k++
		}

		next[i] = k
	}

	return next
}

func getNext1(pattern string) (next []int) {
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
