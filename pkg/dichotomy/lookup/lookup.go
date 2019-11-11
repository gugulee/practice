package lookup

import "fmt"

func Lookup(sortedDict []string, target string) (bool, error) {
	if 0 == len(sortedDict) {
		return false, fmt.Errorf("dict can not be empty")
	}

	if 0 == len(target) {
		return false, fmt.Errorf("target can not be empty")
	}

	left, right := 0, len(sortedDict)-1
	for left <= right {
		middle := (left + right) / 2
		if target == sortedDict[middle] {
			return true, nil
		} else if target > sortedDict[middle] {
			left += 1
		} else {
			right -= 1
		}
	}

	return false, fmt.Errorf("can no find target:%q in dict", target)
}
