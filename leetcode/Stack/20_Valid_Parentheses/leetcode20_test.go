package leetcode20

import "testing"

func TestIsValid(t *testing.T) {
	cases := []struct {
		input string
		valid bool
	}{
		{
			input: "(",
			valid: false,
		},
		{
			input: ")",
			valid: false,
		},
		{
			input: "AA",
			valid: false,
		},
		{
			input: "((",
			valid: false,
		},
		{
			input: "))",
			valid: false,
		},
		{
			input: ")(",
			valid: false,
		},
		{
			input: "()",
			valid: true,
		},
		{
			input: "()[]{}",
			valid: true,
		},
		{
			input: "(]",
			valid: false,
		},
		{
			input: "[)",
			valid: false,
		},
		{
			input: "[[[]",
			valid: false,
		},
		{
			input: "([{}])",
			valid: true,
		},
		{
			input: "([{]])",
			valid: false,
		},
		{
			input: "([{]})",
			valid: false,
		},
		{
			input: "([(())])",
			valid: true,
		},
		{
			input: "[[()]{()}]",
			valid: true,
		},
		{
			input: "[[[[[]]]]}",
			valid: false,
		},
		{
			input: "{[[[[]]]]]",
			valid: false,
		},
		{
			input: "[[[[[]]]]]",
			valid: true,
		},
		{
			input: "{([{([])}])}",
			valid: true,
		},
	}

	for i, c := range cases {
		result := isValid(c.input)
		if c.valid != result {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.valid, result)
		}
	}
}
