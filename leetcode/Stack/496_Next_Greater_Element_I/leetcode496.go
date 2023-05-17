package leetcode496

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// Create a map, so that we can check if the values from nums2 are in the nums1
	h := make(map[int]int, len(nums1))
	for _, n := range nums1 {
		h[n] = -1
	}

	// stack
	stack := []int{}
	for _, n := range nums2 {
		// if there is any element in the stack, comapre the top of the stack with the current number `n` from nums2
		// if the `n` it's greater than the value from stack, than `n` is the next greater value of the value from stack
		// , and update the map by the correspoding key
		for len(stack) > 0 && stack[len(stack)-1] < n {
			// Pop the value from the stack
			val := stack[len(stack)-1]
			// Resize the stack by removing the last one
			stack = stack[:len(stack)-1]
			// Record the next greater value in the map
			h[val] = n
		}

		// If the value exists in nums1, add it to the stack for later use.
		if _, ok := h[n]; ok {
			stack = append(stack, n)
		}
	}

	// Convert the map into an array and maintain the original order of nums1.
	result := make([]int, len(nums1))
	for i, n := range nums1 {
		result[i] = h[n]
	}

	return result
}
