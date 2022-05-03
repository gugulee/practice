package ld

import (
	"fmt"

	"github.com/practice/pkg/tools"
)

/*
a := []byte("mitcmu")
b := []byte("mtacnu")
-------------------------------------------------------
	m	t	a	c	n	u
m	0	1	2	3	4	5
i	1	1	2	3	4	5
t	2	1	2	3	4	5
c	3	2	2	2	3	4
m	4	3	3	3	3	4
u	5	4	4	4	4	3
*/
func levenshteinDistance(a, b []byte) (distance int) {
	lena := len(a)
	lenb := len(b)
	states := make([][]int, lena)
	for i := range states {
		states[i] = make([]int, lenb)
	}

	// fill the states[0][0]
	if a[0] != b[0] {
		states[0][0] = 1
	}

	// fill the states[0][j]
	for j := 1; j < lenb; j++ {
		if a[0] == b[j] {
			states[0][j] = j
		} else {
			states[0][j] = states[0][j-1] + 1
		}
	}

	// fill the states[i][0]
	for i := 1; i < lena; i++ {
		if b[0] == a[i] {
			states[i][0] = i
		} else {
			states[i][0] = states[i-1][0] + 1
		}
	}

	for i := 1; i < lena; i++ {
		for j := 1; j < lenb; j++ {
			r := 0
			if a[i] != b[j] {
				r = 1
			}
			states[i][j] = tools.Min(states[i-1][j]+1, states[i][j-1]+1, states[i-1][j-1]+r)
		}
	}

	printData(states)
	return states[lena-1][lenb-1]
}

func printData(data [][]int) {
	for row := 0; row < len(data); row++ {
		for column := 0; column < len(data[row]); column++ {
			fmt.Print(data[row][column], " ")
		}
		fmt.Println()
	}
}
