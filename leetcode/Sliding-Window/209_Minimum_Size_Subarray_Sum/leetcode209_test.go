package leetcode209

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	cases := []struct {
		target   int
		nums     []int
		expected int
	}{
		{
			target:   7,
			nums:     []int{2, 3, 1, 2, 4, 3},
			expected: 2,
		},
		{
			target:   4,
			nums:     []int{1, 4, 4},
			expected: 1,
		},
		{
			target:   11,
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			expected: 0,
		},
	}

	for i, c := range cases {
		result := minSubArrayLen(c.target, c.nums)
		if c.expected != result {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.expected, result)
		}
	}
}
