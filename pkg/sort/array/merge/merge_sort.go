package merge

// merge merge two slice
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

// merge1 merge a[start:mid+1] and a[mid+1:end+1] into a[start:end+1],
// notice: a[start:mid+1] and a[mid+1:end+1] must be ordered
func merge1(a []int, start, mid, end int) {
	var newList = make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0

	for ; i <= mid && j <= end; k++ {
		if a[i] <= a[j] {
			newList[k] = a[i]
			i++
		} else {
			newList[k] = a[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		newList[k] = a[i]
		k++
	}

	for ; j <= end; j++ {
		newList[k] = a[j]
		k++
	}

	copy(a[start:end+1], newList)
}

// merge2 merge original[start:mid+1] and original[mid+1:end+1] into original[start:end+1],
// notice: original[start:mid+1] and original[mid+1:end+1] must be ordered
func merge2(original []int, start, mid, end int) {
	first := original[start : mid+1]
	second := original[mid+1 : end+1]

	var newList []int
	for {
		// second is empty
		if 0 == len(second) {
			newList = append(newList, first...)
			break
		}

		// first is empty
		if 0 == len(first) {
			newList = append(newList, second...)
			break
		}

		// first[0] is less, remove first[0] from first slice
		if first[0] <= second[0] {
			newList = append(newList, first[0])
			first = first[1:]
		} else { // second[0] is less, remove second[0] from second slice
			newList = append(newList, second[0])
			second = second[1:]
		}
	}

	copy(original[start:end+1], newList)
}

// merge3 that has sentinel merge original[start:mid+1] and original[mid+1:end+1] into original[start:end+1],
// notice: original[start:mid+1] and original[mid+1:end+1] must be ordered
func merge3(original []int, start, mid, end int) {
	max := 999
	first := make([]int, mid-start+1)
	second := make([]int, end-mid)
	copy(first, original[start:mid+1])
	copy(second, original[mid+1:end+1])

	first = append(first, max)
	second = append(second, max)
	i, j, k := 0, 0, start

	for first[i] != max {
		if first[i] <= second[j] {
			original[k] = first[i]
			i++
		} else {
			original[k] = second[j]
			j++
		}
		k++
	}
}

// MergeSort ...
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

// MergeSort1 ...
func MergeSort1(original []int, start, end int) {
	if end <= start {
		return
	}

	mid := (start + end) / 2

	MergeSort1(original, start, mid)
	MergeSort1(original, mid+1, end)
	merge1(original, start, mid, end)
}
