package removeduplicates

func removeDuplicates(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	slow, fast := 1, 1
	for ; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

// preserve two duplicate num at most
func removeDuplicates1(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	slow, fast := 1, 1
	for ; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}