package leetcode209

func minSubArrayLen(target int, nums []int) int {
	// Set the minimal subarray length to len(nums)+1 because even all nums are sumed up, the total length still less than it
	// And it indicates that it there is no such subarray
	min := len(nums) + 1

	total, left := 0, 0
	for right := 0; right < len(nums); right++ {
		total += nums[right]

		// If taget is less than total, move left pointer forward
		for total >= target {
			// Update the minimal subarray length
			length := right - left + 1
			if length < min {
				min = length
			}
			total -= nums[left]
			left++
		}
	}
	if min == len(nums)+1 {
		return 0
	}
	return min
}
