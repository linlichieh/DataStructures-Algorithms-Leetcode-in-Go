package leetcode33

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target int, left int, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	switch {
	case target == nums[mid]:
		return mid
	case nums[left] <= nums[mid]: // if only 2 numbers, the left will be the same as mid, that's why we need `<=` here instead of  `<` only
		// For example, find 1 in `4 5 6 7 0 1 2`
		if target > nums[mid] || target < nums[left] {
			return binarySearch(nums, target, mid+1, right)
		}
		return binarySearch(nums, target, left, mid-1)
	case nums[left] > nums[mid]:
		// For example, find 3 in `4 5 6 7 0 1 2 3 4 5`
		if target < nums[mid] || target > nums[right] {
			return binarySearch(nums, target, left, mid-1)
		}
		return binarySearch(nums, target, mid+1, right)
	default:
		// Won't enter here
		return -1
	}
}
