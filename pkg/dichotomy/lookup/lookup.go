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

// Bsearch lookup the first value in the a, return the index of target
func Bsearch(a []int, value int) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] > value {
			high = mid - 1
		} else if a[mid] < value {
			low = mid + 1
		} else {
			if mid == 0 || a[mid-1] != value {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// bsearch1 lookup the last value in the a, return the index of target
func bsearch1(a []int, value int) int {
	length := len(a)
	low, high := 0, length-1

	for low <= high {
		mid := low + (high-low)>>1
		if value == a[mid] {
			if mid == length-1 || a[mid+1] != value {
				return mid
			}
			low = mid + 1
		} else if value > a[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

// bsearch2 lookup the first value which is equal or greater than target in the a, return the index of target
func bsearch2(a []int, value int) int {
	length := len(a)
	low, high := 0, length-1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if value <= a[mid] {
			if mid == 0 || a[mid-1] < value {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

// bsearch3 lookup the last value which is equal or less than target in the a, return the index of target
func bsearch3(a []int, value int) int {
	length := len(a)
	low, high := 0, length-1

	for low <= high {
		mid := low + (high-low)>>1
		if a[mid] <= value {
			if mid == length-1 || a[mid+1] > value {
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

// bsearch4 lookup the specified value in the circular array, return the index of target
func bsearch4(a []int, value int) int {
	length := len(a)
	low, high := 0, length-1

	for low <= high {
		mid := low + (high-low)>>1
		
		if value == a[mid] {
			return mid
		}

		// a[low:mid+1] is ordered
		if a[low] <= a[mid] {
			// value is in ordered area
			if a[low] <= value && value < a[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}

		} else { // a[mid:high+1] is ordered
			// value is in ordered area
			if a[mid] < value && value <= a[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}
