package leetcode80

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		input    []int
		output   int
		expected []int
	}{
		{
			input:    []int{},
			output:   0,
			expected: []int{}, // It does not matter what you leave beyond the returned k
		},
		{
			input:    []int{1},
			output:   1,
			expected: []int{1},
		},
		{
			input:    []int{1, 1, 1, 2, 2, 3},
			output:   5,
			expected: []int{1, 1, 2, 2, 3, 1},
		},
		{
			input:    []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			output:   7,
			expected: []int{0, 0, 1, 1, 2, 3, 3, 1, 1},
		},
		{
			input:    []int{1, 1, 1},
			output:   2,
			expected: []int{1, 1, 1},
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			output:   5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			input:    []int{1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3},
			output:   6,
			expected: []int{1, 1, 2, 2, 3, 3, 2, 2, 3, 3, 3},
		},
		{
			input:    []int{1, 1, 1, 1, 1, 1},
			output:   2,
			expected: []int{1, 1, 1, 1, 1, 1},
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6},
			output:   6,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			input:    []int{1, 1, 2, 2, 3},
			output:   5,
			expected: []int{1, 1, 2, 2, 3},
		},
	}

	for i, c := range cases {
		// in-place
		output := removeDuplicates(c.input)
		if c.output != output {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, output)
		}
		if !reflect.DeepEqual(c.expected, c.input) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.expected, c.input)
		}
	}
}
