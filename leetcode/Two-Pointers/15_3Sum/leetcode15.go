package leetcode15

import "sort"

// Two pointers + hash
func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}

	// Sort the array first in order to make skipping the duplicate value easier
	sort.Ints(nums)

	// loop the element (locked) in the outer loop
	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicate value
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Use 2 pointers to traverse the array
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			switch {
			case sum > 0:
				right--

			case sum < 0:
				left++

			case sum == 0:
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// Skip the duplicate value, for example `[-2, 0, 0, 0, 1, 1, 1, 1, 2]`
				// NOTE the new left and right compare to their previous one
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return result
}
