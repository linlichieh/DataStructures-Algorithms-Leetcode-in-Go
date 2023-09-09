package leetcode56

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input: [][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			expected: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			input: [][]int{
				{1, 4},
				{4, 5},
			},
			expected: [][]int{
				{1, 5},
			},
		},
		{
			input: [][]int{
				{1, 4},
				{2, 3},
			},
			expected: [][]int{
				{1, 4},
			},
		},
	}

	for i, c := range cases {
		result := merge(c.input)
		if !reflect.DeepEqual(c.expected, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.expected, result)
		}
	}
}
