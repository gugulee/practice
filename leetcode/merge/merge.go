package merge

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := 0, 0, 0

	tmp := make([]int, m+n)

	for (i < m) && j < n {
		if nums1[i] >= nums2[j] {
			tmp[k] = nums2[j]
			j++
		} else {
			tmp[k] = nums1[i]
			i++
		}

		k++
	}

	for ; i < m; i++ {
		tmp[k] = nums1[i]
		k++
	}

	for ; j < n; j++ {
		tmp[k] = nums2[j]
		k++
	}

	copy(nums1, tmp)
}
