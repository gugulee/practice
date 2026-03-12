package get_str_distance

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// getStrDistance get edit distance between a and b
func getStrDistance(a, b string) int {
	var d = make([][]int, len(a)+1)
	for i := range d {
		d[i] = make([]int, len(b)+1)
	}

	for j := 0; j <= len(b); j++ {
		d[0][j] = j
	}

	for i := 0; i <= len(a); i++ {
		d[i][0] = i
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			r := 0
			if a[i-1] != b[j-1] {
				r = 1
			}
			// d[i-1][j]+1: insert a
			// d[i][j-1]+1: insert b
			// d[i-1][j-1]+r: replace
			d[i][j] = min(min(d[i-1][j]+1, d[i][j-1]+1), d[i-1][j-1]+r)
		}
	}

	return d[len(a)][len(b)]
}

// getStrDistance1 get edit distance between a and b
func getStrDistance1(a, b string, aLen, bLen int) int {
	if 0 == aLen {
		return bLen
	}

	if 0 == bLen {
		return aLen
	}

	r := 0
	if a[aLen-1] != b[bLen-1] {
		r = 1
	}

	aInsert := getStrDistance1(a, b, aLen-1, bLen) + 1
	bInsert := getStrDistance1(a, b, aLen, bLen-1) + 1
	replace := getStrDistance1(a, b, aLen-1, bLen-1) + r

	return min(min(aInsert, bInsert), replace)
}
