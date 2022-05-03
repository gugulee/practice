package bruteforce

// aaaaaab
// aab
func bfSearch(main, sub string) bool {
	lenMain := len(main)
	lenSub := len(sub)

	for i := 0; i <= lenMain-lenSub; i++ {
		j := 0
		idx := i
		for j < lenSub {
			if main[idx] != sub[j] {
				break
			}

			idx++
			j++
		}

		if j == lenSub {
			return true
		}
	}

	return false
}

//BF search pattern
func bfSearch1(main string, pattern string) bool {
	//defensive
	if len(main) == 0 || len(pattern) == 0 || len(main) < len(pattern) {
		return false
	}

	for i := 0; i <= len(main)-len(pattern); i++ {
		subStr := main[i : i+len(pattern)]
		if subStr == pattern {
			return true
		}
	}

	return false
}

// bfSearchMatrixString search pattern from matrix string
func bfSearchMatrixString(main, pattern [][]string) bool {
	lenMainRow := len(main)
	lenMainColumn := len(main[0])
	lenPatternRow := len(pattern)
	lenPatternColumn := len(pattern[0])

	if 0 == lenMainRow || 0 == lenPatternRow {
		return false
	}

	if lenMainRow < lenPatternRow || lenMainColumn < lenPatternColumn {
		return false
	}

	for i := 0; i <= lenMainRow-lenPatternRow; i++ {
		for j := 0; j <= lenMainColumn-lenPatternColumn; j++ {
			sub := truncate(main, i, i+lenPatternRow-1, j, j+lenPatternColumn-1)
			if isEqual(sub, pattern) {
				return true
			}
		}
	}

	return false
}

// truncate truncate the string from a,
// e.g., a=
//	{"d", "a", "b", "c", "d"},
//	{"e", "f", "a", "d", "g"},
//	{"c", "c", "a", "f", "h"},
//	{"d", "e", "f", "c", "i"},
// if rowStart=0, rowEnd=1, columnStart=0, columnEnd =2, result=
//	{"d", "a", "b"},
//	{"e", "f", "a"},
func truncate(a [][]string, rowStart, rowEnd, columnStart, columnEnd int) (result [][]string) {
	if 0 == len(a) || rowEnd < rowStart || columnEnd < columnStart || len(a) < rowEnd+1 || len(a[0]) < columnEnd+1 {
		return nil
	}

	for i := rowStart; i <= rowEnd; i++ {
		result = append(result, a[i][columnStart:columnEnd+1])
	}

	return result
}

// isEqual return true if first==second, else return false
func isEqual(first, second [][]string) bool {
	if 0 == len(first) && 0 == len(second) {
		return false
	}

	if len(first) != len(second) {
		return false
	}

	for i := range first {
		if !isSliceEqual(first[i], second[i]) {
			return false
		}
	}

	return true
}

// isSliceEqual return true if first==second, else return false
func isSliceEqual(first, second []string) bool {
	if 0 == len(first) && 0 == len(second) {
		return false
	}

	if len(first) != len(second) {
		return false
	}

	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}
