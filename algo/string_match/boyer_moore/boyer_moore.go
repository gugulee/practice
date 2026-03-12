package boyermoore

import (
	"strings"

	"github.com/practice/pkg/tools"
)

// SIZE ...
const SIZE = 256

// generateBC calculate the location of ss, the subscript of bc is ASCII of character in ss, the value of bc is the location of character
// e.g., ss="abda",
// bc[97]=3, bc[98]=1, bc[100]=2
func generateBC(ss string) (bc []int) {
	bc = make([]int, SIZE)
	for i := range bc {
		bc[i] = -1
	}

	for i, s := range ss {
		bc[s] = i
	}

	return bc
}

// generateGS calculate the location of ss,
// the subscript of suffix is the length of the suffix substring of ss, the value of suffix is location of the suffix substring(it is -1 if not found),
// the subscript of prefix is the length of the suffix substring, the value of prefix represent if appear at the beggin of ss,
// e.g., ss="cabcab",
// the suffix substring is "b", the prefix substring is "cabca", the length(len(b)) is 1, suffix[1]=2, prefix[1]=false
// the suffix substring is "ab", the prefix substring is "cabc", the length is 2, suffix[2]=1, prefix[2]=false
// the suffix substring is "cab", the prefix substring is "cab", the length is 3, suffix[3]=0, prefix[3]=true
// the suffix substring is "bcab", the prefix substring is "ca", the length is 4, suffix[4]=-1, prefix[4]=false
// the suffix substring is "abcab", the prefix substring is "c", the length is 5, suffix[5]=-1, prefix[5]=false
func generateGS(pattern string) (prefix []bool, suffix []int) {
	m := len(pattern)
	suffix = make([]int, m)
	prefix = make([]bool, m)

	//init
	for i := 0; i < m; i++ {
		suffix[i] = -1
		prefix[i] = false
	}
	// "cabcab"
	for i := 0; i < m-1; i++ {
		j := i
		k := 0
		for j >= 0 && pattern[j] == pattern[m-1-k] {
			j--
			k++
			suffix[k] = j + 1
		}

		if j == -1 {
			prefix[k] = true
		}
	}

	return prefix, suffix
}

func generateGS1(ss string) (prefix []bool, suffix []int) {
	lenSS := len(ss)
	prefix = make([]bool, lenSS)
	suffix = make([]int, lenSS)

	for i := lenSS - 1; 1 <= i; i-- {
		// s is the suffix substring of ss, p is the prefix substring of ss
		s := ss[i:]
		p := ss[:i]
		idx := len(s)

		location := strings.LastIndex(p, s)
		suffix[idx] = location
		if 0 == location {
			prefix[idx] = true
		}
	}

	return prefix, suffix
}

// bmSearch search main with pattern
func bmSearch(main string, pattern string) bool {
	lenMain := len(main)
	lenPattern := len(pattern)

	//defensive
	if lenMain == 0 || lenPattern == 0 || lenMain < lenPattern {
		return false
	}

	bc := generateBC(pattern)
	prefix, suffix := generateGS(pattern)

	for i := 0; i <= lenMain-lenPattern; {
		/* check from right to left */
		j := lenPattern - 1
		for ; 0 <= j; j-- {
			if main[i+j] != pattern[j] {
				break
			}
		}

		// found
		if j < 0 {
			return true
		}

		// cabcab
		var bch, gsh int
		// bad character heuristic
		bch = j - bc[main[i+j]]

		// has good suffix, apply to good suffix heuristic
		if j < lenPattern-1 {
			gsh = moveByGS(j, lenPattern, suffix, prefix)
		}

		i += tools.Max(bch, gsh)
	}

	return false
}

func moveByGS(j, m int, suffix []int, prefix []bool) int {
	// the length of good suffix
	k := m - 1 - j
	if -1 != suffix[k] {
		return j + 1 - suffix[k]
	}

	for r := j + 2; r <= m-1; r++ {
		if prefix[m-r] {
			return r
		}
	}

	return m
}
