package ld

import "math"

var minimum = math.MaxInt16

// a := []byte("mitcmu")
// b := []byte("mtacnu")
// levenshteinDistance get the levenshtein distance
func levenshteinDistance(a, b []byte, i, j, distance int) {
	lena := len(a)
	lenb := len(b)

	if i == lena || j == lenb {
		if i < lena {
			distance += lena - i
		}
		if j < lenb {
			distance += lenb - j
		}

		if minimum > distance {
			minimum = distance
		}
		return
	}

	if a[i] == b[j] {
		levenshteinDistance(a, b, i+1, j+1, distance)
	} else {
		// delete a[i] or insert a character before b[j]
		levenshteinDistance(a, b, i+1, j, distance+1)

		// delete b[j] or insert a character before a[i]
		levenshteinDistance(a, b, i, j+1, distance+1)

		// replace
		levenshteinDistance(a, b, i+1, j+1, distance+1)
	}

}
