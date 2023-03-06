package leetcode3

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		input  string
		output int
	}{
		{
			input:  "",
			output: 0,
		},
		{
			input:  " ",
			output: 1,
		},
		{
			input:  "  ",
			output: 1,
		},
		{
			input:  " a",
			output: 2,
		},
		{
			input:  "dvdf",
			output: 3,
		},
		{
			input:  "abcabcbb",
			output: 3,
		},
		{
			input:  "bbbbb",
			output: 1,
		},
		{
			input:  "pwwkew",
			output: 3,
		},
		{
			input:  "pcwpek",
			output: 5,
		},
		{
			input:  "ppcwpck",
			output: 4,
		},
		{
			input:  "abcccdab",
			output: 4,
		},
		{
			input:  "32483331469527666",
			output: 8,
		},
		{
			input:  "##$%%^^&&$#@@!",
			output: 4,
		},
		{
			input:  "444434 sd3222d",
			output: 6,
		},
	}

	for i, c := range cases {
		result := lengthOfLongestSubstring(c.input)
		if c.output != result {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
