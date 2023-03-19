package leetcode283

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	output := make([]int, len(nums), len(nums))

	// Calculate the prefix product up to each element
	pre := 1
	for i := 0; i < len(nums); i++ {
		if i-1 < 0 {
			output[i] = 1
			continue
		}
		pre *= nums[i-1]
		output[i] = pre
	}

	// Calculate the suffix product and multiply corresponding prefix product
	post := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+1 >= len(nums) {
			output[i] *= post
			continue
		}
		post *= nums[i+1]
		output[i] *= post
	}

	return output
}
