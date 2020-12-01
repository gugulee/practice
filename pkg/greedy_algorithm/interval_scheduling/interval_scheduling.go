package intervalscheduling

func intervalScheduling(intvs [][]int) (result [][]int) {
	length := len(intvs)

	// sort
	sortDyadicSlice(intvs)
	minEnd := intvs[0][1]
	result = append(result, intvs[0])

	for i := 1; i < length; i++ {
		if intvs[i][0] >= minEnd {
			result = append(result, intvs[i])
			minEnd = intvs[i][1]
		}
	}

	return result
}

// s := [][]int{{3, 6}, {1, 3}, {2, 4}}
func sortDyadicSlice(s [][]int) {
	length := len(s)
	for i := 0; i < length; i++ {
		flag := false
		for j := 0; j < length-i-1; j++ {
			if s[j][1] > s[j+1][1] {
				s[j], s[j+1] = s[j+1], s[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}
}
