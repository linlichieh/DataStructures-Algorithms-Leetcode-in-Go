package leetcode560

import "testing"

func TestSubarraySum(t *testing.T) {
	cases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2,
		},
		{
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 2,
		},
		{
			nums:     []int{1, -1, 1, 1, 1, 1},
			k:        3,
			expected: 4,
		},
		{
			nums:     []int{1, -1, 2, -1, 2, -3, 3},
			k:        3,
			expected: 5,
		},
	}

	for i, c := range cases {
		result := subarraySum(c.nums, c.k)
		if c.expected != result {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.expected, result)
		}
	}
}
