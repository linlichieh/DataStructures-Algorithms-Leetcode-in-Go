package leetcode1

// One-pass hash table solution instead of two-pass which fills the hash with all elements first
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		complement := target - v
		if j, ok := m[complement]; ok {
			return []int{i, j}
		}
		m[v] = i
	}
	return []int{}
}
