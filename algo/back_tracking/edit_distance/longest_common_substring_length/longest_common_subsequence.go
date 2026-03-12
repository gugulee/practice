package lcs

import (
	"math"
)

var maxLcs = math.MinInt16

/*
a := []byte("mitcmu")
b := []byte("mtacnu")
*/
func lcs(a, b []byte, i, j, curLcs int) {
	lena := len(a)
	lenb := len(b)

	if i == lena || j == lenb {
		if maxLcs < curLcs {
			maxLcs = curLcs
		}
		return
	}

	if a[i] == b[j] {
		lcs(a, b, i+1, j+1, curLcs+1)
	} else {
		lcs(a, b, i+1, j, curLcs)
		lcs(a, b, i, j+1, curLcs)
	}
}
