package leetcode76

import (
	"testing"
)

func TestMinWindow(t *testing.T) {
	cases := []struct {
		input1 string
		input2 string
		output string
	}{
		{
			input1: "",
			input2: "",
			output: "",
		},
		{
			input1: "ADOBECODEBANC",
			input2: "ABC",
			output: "BANC",
		},
		{
			input1: "a",
			input2: "a",
			output: "a",
		},
		{
			input1: "a",
			input2: "aa",
			output: "",
		},
		{
			input1: "ADCECDZIBZBECDCAGFB",
			input2: "ABCC",
			output: "BECDCA",
		},
		{
			input1: "ABCD",
			input2: "BCD",
			output: "BCD",
		},
		{
			input1: "AAAABCDEF",
			input2: "ACD",
			output: "ABCD",
		},
		{
			input1: "ABCDEF",
			input2: "GH",
			output: "",
		},
		{
			input1: "ABCDEF",
			input2: "CD",
			output: "CD",
		},
		{
			input1: "ABCDEGHABC",
			input2: "AC",
			output: "ABC",
		},
		{
			input1: "ABCAADBCDEBC",
			input2: "ABC",
			output: "ABC",
		},
		{
			input1: "ABAACDA",
			input2: "AAB",
			output: "ABA",
		},
		{
			input1: "AAB",
			input2: "BA",
			output: "AB",
		},
		{
			input1: ", ",
			input2: "A",
			output: "",
		},
	}

	for i, c := range cases {
		result := minWindow(c.input1, c.input2)
		if c.output != result {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
