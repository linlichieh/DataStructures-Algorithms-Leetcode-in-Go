package leetcode283

func divisionSolution(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	total := 1

	// counter is for dealing with 0
	zeroCounter := 0

	// Multiply all elements
	for _, num := range nums {
		// If there's zero in it, skip it
		if num == 0 {
			zeroCounter++
			continue
		}
		total *= num
	}
	for i, num := range nums {
		// If the element is 0, but there are other 0 in the array. The value is still be 0.
		// If the elemnt isn't 0 and there is any 0 in the array, its value should be 0.
		if (num == 0 && zeroCounter > 1) || (num != 0 && zeroCounter > 0) {
			nums[i] = 0
		} else if num == 0 {
			// If the element is 0 and it is the only 0 in the array, so the value should be the cumulative multiplication.
			nums[i] = total
		} else {
			// If there is no 0 involved, divied itself with the cumulative multiplication.
			nums[i] = total / num
		}
	}
	return nums
}
