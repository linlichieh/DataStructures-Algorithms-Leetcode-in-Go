package leetcode1

// One-pass hash table solution instead of two-pass which fills the hash with all elements first
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		complement := target - v
		// `i != j` to skip using the same element twice
		if j, ok := m[complement]; ok && i != j {
			return []int{i, j}
		}
		m[v] = i
	}
	return []int{}
}
