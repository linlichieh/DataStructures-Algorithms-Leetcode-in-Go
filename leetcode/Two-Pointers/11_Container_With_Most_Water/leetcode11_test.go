package leetcode11

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	cases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			output: 49,
		},
		{
			input:  []int{1, 1},
			output: 1,
		},
		{
			input:  []int{4, 3, 2, 1, 4},
			output: 16,
		},
		{
			input:  []int{1, 2, 1},
			output: 2,
		},
		{
			input:  []int{2, 3, 4, 5, 18, 17, 6},
			output: 17,
		},
		{
			input:  []int{1, 1, 1, 1, 1, 1, 1, 1},
			output: 7,
		},
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			output: 20,
		},
		{
			input:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			output: 20,
		},
		{
			input:  []int{1, 3, 2, 5, 25, 24, 5},
			output: 24,
		},
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1},
			output: 24,
		},
	}

	for _, c := range cases {
		result := maxArea(c.input)
		if c.output != result {
			t.Errorf("expect '%v', but got '%v'", c.output, result)
		}
	}
}
