package lcs

import (
	"fmt"

	"github.com/practice/pkg/tools"
)

/*
a := []byte("mitcmu")
b := []byte("mtacnu")
-------------------------------------------------------
	''	m	t	a	c	n	u
''	0	0	0	0	0	0	0
m	0	1	1	1	1	1	1
i	0	1	1	1	1	1	1
t	0	1	2	2	2	2	2
c	0	1	2	2	3	3	3
m	0	1	2	2	3	3	3
u	0	1	2	2	3	3	4
*/

func lcs(a, b []byte) int {
	lena := len(a)
	lenb := len(b)
	states := make([][]int, lena+1)
	for i := range states {
		states[i] = make([]int, lenb+1)
	}

	for i := 1; i < lena+1; i++ {
		for j := 1; j < lenb+1; j++ {
			if a[i-1] == b[j-1] {
				states[i][j] = states[i-1][j-1] + 1
			} else {
				states[i][j] = tools.Max(states[i-1][j], states[i][j-1])
			}
		}
	}

	printData(states)
	return states[lena][lenb]
}

func printData(data [][]int) {
	for row := 0; row < len(data); row++ {
		for column := 0; column < len(data[row]); column++ {
			fmt.Print(data[row][column], " ")
		}
		fmt.Println()
	}
}
