package merge_sort

// merge_sort merge_sort two slice
// notice: first and two must be ordered
func merge(first, two []int) []int {
	var newList []int
	for {
		// two is empty
		if 0 == len(two) {
			return append(newList, first...)
		}

		// first is empty
		if 0 == len(first) {
			return append(newList, two...)
		}

		// first[0] is less, remove first[0] from first slice
		if first[0] < two[0] {
			newList = append(newList, first[0])
			first = first[1:]
		} else { // two[0] is less, remove two[0] from two slice
			newList = append(newList, two[0])
			two = two[1:]
		}
	}
}

func MergeSort(original []int) []int {
	if 0 == len(original) {
		return nil
	}

	// if only one in original, return the slice
	if 1 == len(original) {
		return original
	}

	length := len(original)
	left := MergeSort(original[0 : length/2])
	right := MergeSort(original[length/2:])
	return merge(left, right)
}

func MergeSort1(original []int) []int {
	if 0 == len(original) {
		return nil
	}

	var r []int
	// push original
	var stack [][]int
	stack = append(stack, original)

	for len(stack) > 0 {
		// pop
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// if only one in current, merge it to result
		if 1 == len(current) {
			r = merge(r, current)
			continue
		}

		// push
		length := len(current)
		stack = append(stack, current[length/2:])
		stack = append(stack, current[0:length/2])
	}
	return r
}
