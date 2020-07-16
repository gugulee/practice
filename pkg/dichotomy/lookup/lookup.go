package lookup

import "fmt"

// Lookup ...
func Lookup(sortedDict []string, target string) (bool, error) {
	if 0 == len(sortedDict) {
		return false, fmt.Errorf("dict can not be empty")
	}

	if 0 == len(target) {
		return false, fmt.Errorf("target can not be empty")
	}

	left, right := 0, len(sortedDict)-1
	for left <= right {
		middle := left + (right-left)>>1
		if target == sortedDict[middle] {
			return true, nil
		} else if target > sortedDict[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return false, fmt.Errorf("can no find target:%q in dict", target)
}

// Lookup1 ...
func Lookup1(sortedDict []string, target string, left, right int) bool {
	if left > right {
		return false
	}

	middle := left + (right-left)>>1

	if target == sortedDict[middle] {
		return true
	} else if target > sortedDict[middle] {
		return Lookup1(sortedDict, target, middle+1, right)
	} else {
		return Lookup1(sortedDict, target, left, middle-1)
	}
}
