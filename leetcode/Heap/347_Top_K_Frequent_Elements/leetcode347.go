package leetcode347

func topKFrequent(nums []int, k int) []int {
	// Hash map to count the frequency of each element
	counter := make(map[int]int)
	for _, num := range nums {
		// check if num exists
		if counter[num] == 0 {
			counter[num] = 1
		} else {
			counter[num] += 1
		}
	}

	// Create buckets with size `n`
	buckets := make(map[int][]int, len(nums))
	for num, freq := range counter {
		if len(buckets[freq]) == 0 {
			buckets[freq] = []int{num}
		} else {
			buckets[freq] = append(buckets[freq], num)
		}
	}

	// Get top K frequent elements
	result := []int{}
	for i := len(nums); i > 0; i-- {
		if len(buckets[i]) == 0 {
			continue
		}
		if len(buckets[i]) <= k-len(result) {
			result = append(result, buckets[i]...)
			continue
		}
		for j := 0; j < k-len(result); j++ {
			result = append(result, buckets[i][j])
		}

		if len(result) == k {
			break
		}
	}
	return result
}
