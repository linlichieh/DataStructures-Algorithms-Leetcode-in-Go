package leetcode53

import "testing"

func TestMaxSubArray(t *testing.T) {
	cases := []struct {
		arr      []int
		expected int
	}{
		{
			arr:      []int{-1},
			expected: -1,
		},
		{
			arr:      []int{1},
			expected: 1,
		},
		{
			arr:      []int{-2, -1},
			expected: -1,
		},
		{
			arr:      []int{5, 4, -1, 7, 8},
			expected: 23,
		},
		{
			arr:      []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			arr:      []int{-1, 5, -2, 4, -7, 1},
			expected: 7,
		},
	}

	for i, c := range cases {
		result := maxSubArray(c.arr)
		if result != c.expected {
			t.Errorf("case(%d) expected %v, but got %d", i, c.expected, result)
		}
	}
}
