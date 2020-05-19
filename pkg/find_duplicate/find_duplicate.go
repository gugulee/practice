package find_duplicate

import (
	"github.com/practice/pkg/sort/merge_sort"
)

// FindScopeDuplicate find ths duplicate digital if the digital is continuous
// attention: the digital must be continuous
func FindScopeDuplicate(origin []int, n int) int {
	originXor := 0
	nXor := 0
	out := 0

	for _, t := range origin {
		originXor ^= t
	}

	for i := 1; i <= n; i++ {
		nXor ^= i
	}

	for i := 1; i <= n; i++ {
		if originXor^nXor == i {
			out = i
			break
		}
	}

	return out
}

// FindDuplicate find the duplicate digital
func FindDuplicate(origin []int) (int, error) {
	var out = -1

	origin = merge_sort.MergeSort(origin)
	maxLen := origin[len(origin)-1]/64 + 1
	storeBit := make([]int, maxLen)

	// if digital is 5, store in the fifth of storeBit with bit:1(storeBit[0] |= 0b10000)
	for _, i := range origin {
		idx := i / 64
		place := 1 << (i % 64)

		// "0==storeBit[x] & place" means that not find the digital, store the digital in the place
		if 0 == storeBit[idx]&place {
			storeBit[idx] |= place
		} else { // find the digital in the place
			out = i
			break
		}
	}

	return out, nil
}
