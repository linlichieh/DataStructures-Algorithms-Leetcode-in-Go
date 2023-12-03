package leetcode435

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {
	cases := []struct {
		intervals [][]int
		expected  int
	}{
		{
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			expected:  1,
		},
		{
			intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
			expected:  2,
		},
		{
			intervals: [][]int{{1, 2}, {2, 3}},
			expected:  0,
		},
		{
			intervals: [][]int{{1, 4}, {2, 3}, {3, 5}},
			expected:  1,
		},
		{
			intervals: [][]int{{1, 5}, {2, 3}, {4, 7}},
			expected:  1,
		},
		{
			intervals: [][]int{{1, 4}, {2, 3}, {5, 7}},
			expected:  1,
		},
	}

	for i, c := range cases {
		result := eraseOverlapIntervals(c.intervals)
		if c.expected != result {
			t.Errorf("(%d) expected '%v', but got '%v'", i, c.expected, result)
		}
	}
}
