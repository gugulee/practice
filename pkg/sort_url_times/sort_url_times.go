package sorturltimes

func getMax(urls map[string]int) int {
	max := -1

	for _, times := range urls {
		if times > max {
			max = times
		}
	}

	return max
}

func sortURLTimes(urls map[string]int) (r []string) {
	r = make([]string, len(urls))
	max := getMax(urls)
	buckets := make([][]string, max+1)
	for url, times := range urls {
		buckets[times] = append(buckets[times], url)
	}

	position := 0
	for _, b := range buckets {
		bLen := len(b)
		if bLen == 0 {
			continue
		}

		copy(r[position:], b)
		position += bLen
	}

	return r
}
