package leetcode70

import "testing"

func TestClimbStairs(t *testing.T) {
	cases := []struct {
		input  int
		output int
	}{
		{
			input:  0,
			output: 0,
		},
		{
			input:  1,
			output: 1,
		},
		{
			input:  2,
			output: 2,
		},
		{
			input:  3,
			output: 3,
		},
		{
			input:  4,
			output: 5,
		},
		{
			input:  5,
			output: 8,
		},
		{
			input:  6,
			output: 13,
		},
	}

	for i, c := range cases {
		result := climbStairs(c.input)
		if c.output != result {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
