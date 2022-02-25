package longestcommonprefix

import (
	"strings"

	"github.com/practice/pkg/tools"
)

func longestCommonPrefix(strs []string) string {
	defer tools.FuncTime()()
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, k := range strs {
		for strings.Index(k, prefix) != 0 {
			if 0 == len(prefix) {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}

	return prefix
}

// divide and conquer
func longestCommonPrefix1(strs []string) string {
	defer tools.FuncTime()()
	length := len(strs)
	if length < 1 {
		return ""
	}

	return commonPrefix(strs, 0, length-1)
}

func commonPrefix(strs []string, left, right int) string {
	if left == right {
		return strs[left]
	}

	mid := (left + right) / 2
	leftPrefix := commonPrefix(strs, left, mid)
	rightPrefix := commonPrefix(strs, mid+1, right)

	return prefix(leftPrefix, rightPrefix)
}

func prefix(left, right string) string {
	minlen := len(left)
	lenRight := len(right)
	if minlen > lenRight {
		minlen = lenRight
	}
	index := -1
	for i := 0; i < minlen; i++ {
		if left[i] != right[i] {
			break
		}
		index = i
	}
	return left[:index+1]
}
