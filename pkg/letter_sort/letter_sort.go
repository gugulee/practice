package lettersort

func letterSort(a []string) {
	length := len(a)

	for i, j := 0, length-1; i < j; {
		switch {
		case "a" <= a[i]: // minuscule
			i++
			fallthrough
		case "A" <= a[j] && a[j] < "a": // capital
			j--
			continue
		}

		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

func letterSort1(a []string) {
	// there are three types in a, so need three buckets.
	// store minuscule in bucket[0], digit in bucket[1] and capital in bucket[2]
	bucket := make([][]string, 3)

	for _, letter := range a {
		switch {
		case "a" <= letter: // minuscule
			bucket[0] = append(bucket[0], letter)
		case letter < "A": // digit
			bucket[1] = append(bucket[1], letter)
		case "A" <= letter && letter < "a": // capital
			bucket[2] = append(bucket[2], letter)
		}
	}

	position := 0
	for _, b := range bucket {
		copy(a[position:], b)
		position += len(b)
	}
}
