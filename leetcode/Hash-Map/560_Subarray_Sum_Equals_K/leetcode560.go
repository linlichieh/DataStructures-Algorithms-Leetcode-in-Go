package leetcode560

func subarraySum(nums []int, k int) int {
	result, sum := 0, 0
	sumFreq := map[int]int{
		0: 1, // the default is one
	}
	for _, v := range nums {
		sum += v
		// it indicates that there are one or more subarrays ending at the current index with a sum equal to k.
		if freq, ok := sumFreq[sum-k]; ok {
			result += freq
		}
		sumFreq[sum]++
	}
	return result
}
