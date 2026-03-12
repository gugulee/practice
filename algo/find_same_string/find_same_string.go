package findsamestring

func findSameString(one, two []string) (r []string) {
	mapString := make(map[string]struct{}, len(one))
	for _, s := range one {
		mapString[s] = struct{}{}
	}

	for _, s := range two {
		if _, ok := mapString[s]; ok {
			r = append(r, s)
		}
	}

	return r
}
