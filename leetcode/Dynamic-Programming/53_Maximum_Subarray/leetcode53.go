package leetcode53

func maxSubArray(nums []int) int {
	// the lengh of nums is always more than 1
	// set default maximum as the first integer of the nums
	maximum := nums[0]
	sum := 0

	for _, val := range nums {
		// If sum is negative, reset the sum
		if sum < 0 {
			sum = 0
		}
		// Add up each item whil moving forward
		sum += val
		// Set the maximum
		if sum > maximum {
			maximum = sum
		}
	}

	return maximum
}
