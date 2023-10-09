package leetcode121

import "testing"

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		prices   []int
		expected int
	}{
		{
			prices:   []int{},
			expected: 0,
		},
		{
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			prices:   []int{7, 2, 5, 6, 1, 3},
			expected: 4,
		},
		{
			prices:   []int{7, 2, 3, 6, 8, 1, 5},
			expected: 6,
		},
		{
			prices:   []int{7, 4, 6, 3, 7},
			expected: 4,
		},
	}

	for i, c := range cases {
		result := maxProfit(c.prices)
		if c.expected != result {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.expected, result)
		}
	}
}
