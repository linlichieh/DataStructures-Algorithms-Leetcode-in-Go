package leetcode217

func containsDuplicate(nums []int) bool {
	numMap := make(map[int]bool)

	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			return true
		}
		numMap[num] = true
	}

	return false
}
