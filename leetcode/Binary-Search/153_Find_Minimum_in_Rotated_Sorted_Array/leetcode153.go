package leetcode153

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return binarySearch(nums, 0, len(nums)-1)
}

func binarySearch(nums []int, left int, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	// `mid == left` only doesn't work in the test case `2 3 4 5 1` that outputs `5`
	if mid == left && mid == right {
		return nums[mid]
	}

	// If the middle element is greater than the last element, then the minimum element is in the right half
	if nums[mid] <= nums[right] {
		return binarySearch(nums, left, mid)
	}
	return binarySearch(nums, mid+1, right)
}
