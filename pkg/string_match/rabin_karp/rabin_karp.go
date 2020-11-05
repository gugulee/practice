package rabinkarp

import (
	"math"
)

const defaultBase = 26

// exponent get the value(base^exponent) and store the value in the result,
// e.g., result[2]=base^2, result[1]=base^1, result[0]=base^0
func exponent(base, exponent int) (result []int) {
	result = make([]int, exponent+1)
	for i := range result {
		result[i] = int(math.Pow(float64(base), float64(i)))
	}
	return result
}

func rkSearch(main string, pattern string) bool {
	lenMain := len(main)
	lenPattern := len(pattern)

	exponent := exponent(defaultBase, lenPattern)

	// store the hash of main string
	hashMain := make([]int, lenMain-lenPattern+1)
	// store the hash of pattern
	hashPattern := 0

	// calculate the hash of pattern and hashMain[0]
	for i := range pattern {
		idx := lenPattern - 1 - i
		hashPattern += exponent[idx] * int(pattern[i]-'a')
		hashMain[0] += exponent[idx] * int(main[i]-'a')
	}

	// calculate the hash of hashMain[1:]
	for i := 1; i < lenMain-lenPattern+1; i++ {
		hashMain[i] = defaultBase*(hashMain[i-1]-exponent[lenPattern-1]*int(main[i-1]-'a')) + exponent[0]*int(main[i+lenPattern-1]-'a')
	}

	// compare the hash of main and pattern
	for i := range hashMain {
		if hashPattern == hashMain[i] {
			return true
		}
	}
	return false
}

// rkSearchMatrixString search pattern from matrix string,
// h[0][0] represent the hash of s[0][0] -> s[1][2],
// h[0][1] represent the hash of s[0][1] -> s[1][3],
// h[1][0]	represent the hash of s[1][0] -> s[2][2],
// r=lenPatternRow, c=lenPatternColumn,
// e.g., if pattern=
//	{"d", "a", "b"},
//	{"e", "f", "a"},
// r=2, c=3,
// h[i][j]=10^(r-1)*(26^(c-1)*(s[i][j]-'a')+26^(c-2)*(s[i][j+1]-'a')+...+26^0*(s[i][j+c-1]-'a'))+
//	10^(r-2)*(26^(c-1)*(s[i+1][j]-'a')+26^(c-2)*(s[i+1][j+1]-'a')+...+26^0*(s[i+1][j+c-1]-'a'))+...
// 	10^0*(26^(c-1)*(s[i+r-1][j]-'a')+26^(c-2)*(s[i+r-1][j+1]-'a')+...+26^0*(s[i+r-1][j+c-1]-'a'))
// h[i][j+1]=
//	26*(h[i][j]-10^(r-1)*26^(c-1)*(s[i][j]-'a')-10^(r-2)*26^(c-1)*(s[i+1][j]-'a')-...-10^0*26^(c-1)*(s[i+r-1][j]-'a'))+
//	10^(r-1)*26^0*(s[i][j+c]-'a')+10^(r-2)*26^0*(s[i+1][j+c]-'a')+...+10^0*26^0*(s[i+r-1][j+c]-'a')
// h[i+1][j]=
//	10*h[i][j]-10^(r)*(26^(c-1)*(s[i][j]-'a')+26^(c-2)*(s[i][j+1]-'a')+...+26^0*(s[i][j+c-1]-'a'))+
//	10^0*(26^(c-1)*(s[i+r][j]-'a')+26^(c-2)*(s[i+r][j+1]-'a')+...+26^0*(s[i+r][j+c-1]-'a'))
func rkSearchMatrixString(main, pattern [][]string) bool {
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

	exponent26 := exponent(defaultBase, lenPatternColumn)
	exponent10 := exponent(10, lenPatternRow)

	// store the hash of main string
	hashMain := make([][]int, lenMainRow-lenPatternRow+1)
	for i := range hashMain {
		hashMain[i] = make([]int, lenMainColumn-lenPatternColumn+1)
	}

	// store the hash of pattern
	hashPattern := 0

	// calculate the hash of pattern and hashMain[0][0]
	for i := range pattern {
		tmpPattern := 0
		tmpMain := 0
		for j := range pattern[i] {
			idxColumn := lenPatternColumn - 1 - j
			tmpPattern += exponent26[idxColumn] * int([]byte(pattern[i][j])[0]-'a')
			tmpMain += exponent26[idxColumn] * int([]byte(main[i][j])[0]-'a')
		}
		idxRow := lenPatternRow - 1 - i
		hashPattern += exponent10[idxRow] * tmpPattern
		hashMain[0][0] += exponent10[idxRow] * tmpMain
	}

	// calculate the hash of hashMain[1:][0]
	for i := 1; i <= lenMainRow-lenPatternRow; i++ {
		tmp := 0
		idxColumn := lenPatternColumn - 1
		for j := 0; j <= lenPatternColumn-1; j++ {
			tmp += exponent26[idxColumn]*int([]byte(main[i-1+lenPatternRow][j])[0]-'a') -
				exponent10[lenPatternRow]*(exponent26[idxColumn]*int([]byte(main[i-1][j])[0]-'a'))
			idxColumn--
		}

		hashMain[i][0] = 10*hashMain[i-1][0] + tmp
	}

	// calculate the hash of hashMain[i][j]
	for i := 0; i <= lenMainRow-lenPatternRow; i++ {
		for j := 1; j <= lenMainColumn-lenPatternColumn; j++ {
			tmp := 0
			idxRow := lenPatternRow - 1
			for k := i; k <= i+lenPatternRow-1; k++ {
				tmp += exponent10[idxRow]*int([]byte(main[k][j+lenPatternColumn-1])[0]-'a') -
					exponent26[lenPatternColumn]*exponent10[idxRow]*int([]byte(main[k][j-1])[0]-'a')
				idxRow--
			}

			hashMain[i][j] = 26*hashMain[i][j-1] + tmp
		}
	}

	for i := range hashMain {
		for j := range hashMain[i] {
			if hashPattern == hashMain[i][j] {
				return true
			}
		}
	}
	return false
}
