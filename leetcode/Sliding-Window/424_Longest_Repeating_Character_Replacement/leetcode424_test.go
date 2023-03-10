package leetcode424

import (
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	cases := []struct {
		input1 string
		input2 int
		output int
	}{
		{
			input1: "",
			input2: 0,
			output: 0,
		},
		{
			input1: "A",
			input2: 1,
			output: 1,
		},
		{
			input1: "AA",
			input2: 1,
			output: 2,
		},
		{
			input1: "ABC",
			input2: 3,
			output: 3,
		},
		{
			input1: "ABCD",
			input2: 3,
			output: 4,
		},
		{
			input1: "ABAB",
			input2: 2,
			output: 4,
		},
		{
			input1: "AAABB",
			input2: 2,
			output: 5,
		},
		{
			input1: "ECEBA",
			input2: 2,
			output: 4,
		},
		{
			input1: "AABABBA",
			input2: 1,
			output: 4,
		},
		{
			input1: "AAAAAA",
			input2: 2,
			output: 6,
		},
		{
			input1: "ABABAB",
			input2: 2,
			output: 5,
		},
		{
			input1: "ABACDEF",
			input2: 2,
			output: 4,
		},
		{
			input1: "ABACCCC",
			input2: 2,
			output: 6,
		},
		{
			input1: "ABCABCABC",
			input2: 3,
			output: 5,
		},
		{
			input1: "AABABBBBBA",
			input2: 1,
			output: 7,
		},
		{
			input1: "AABACCCCCCAB",
			input2: 1,
			output: 7,
		},
		{
			input1: "SDFFDSASSDSAASD",
			input2: 3,
			output: 7,
		},
	}

	for i, c := range cases {
		result := characterReplacement(c.input1, c.input2)
		if c.output != result {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
