package removeelement

func removeElement(nums []int, val int) int {
	i, j := len(nums)-1, len(nums)-1

	for ; i >= 0; i-- {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}

	return j + 1
}
