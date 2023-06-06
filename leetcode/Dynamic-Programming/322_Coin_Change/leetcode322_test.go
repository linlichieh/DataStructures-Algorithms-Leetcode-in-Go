package leetcode322

import "testing"

func TestCoinChange(t *testing.T) {
	cases := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{
			coins:    []int{2, 4},
			amount:   5,
			expected: -1,
		},
		{
			coins:    []int{2, 3},
			amount:   4,
			expected: 2,
		},
		{
			coins:    []int{1, 3, 4, 5},
			amount:   7,
			expected: 2,
		},
		{
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3,
		},
		{
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			coins:    []int{1},
			amount:   0,
			expected: 0,
		},
		{
			coins:    []int{1},
			amount:   2,
			expected: 2,
		},
		{
			coins:    []int{2},
			amount:   1,
			expected: -1,
		},
	}

	for i, c := range cases {
		result := coinChange(c.coins, c.amount)
		if c.expected != result {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.expected, result)
		}
	}
}
