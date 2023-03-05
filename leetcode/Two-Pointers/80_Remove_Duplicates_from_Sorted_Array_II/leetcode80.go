package leetcode80

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	// Use 2-pointer approach
	// We only need current element to compare to the element 2 position back, so that we start from index `2`
	i, j := 2, 2
	for j < len(nums) {
		// If the value of index j doesn't equal to the i-2, after we swap the value, it won't be the third same vlaue in a row
		if nums[j] != nums[i-2] {
			// Sawp these 2 and move i and j forward by 1
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
			continue
		}
		// If they are the same, we want to skip it
		j++
	}
	return i
}
