package leetcode153

import (
	"reflect"
	"testing"
)

func TestFindMin(t *testing.T) {
	cases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{-1},
			output: -1,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			output: 1,
		},
		{
			input:  []int{2, 3, 4, 5, 1},
			output: 1,
		},
		{
			input:  []int{3, 4, 5, 1, 2},
			output: 1,
		},
		{
			input:  []int{4, 5, 1, 2, 3},
			output: 1,
		},
		{
			input:  []int{5, 1, 2, 3, 4},
			output: 1,
		},
		{
			input:  []int{0, 1, 2, 3, 4, 5},
			output: 0,
		},
		{
			input:  []int{5, 0, 1, 2, 3, 4},
			output: 0,
		},
		{
			input:  []int{4, 5, 0, 1, 2, 3},
			output: 0,
		},
		{
			input:  []int{3, 4, 5, 0, 1, 2},
			output: 0,
		},
		{
			input:  []int{2, 3, 4, 5, 0, 1},
			output: 0,
		},
		{
			input:  []int{1, 2, 3, 4, 5, 0},
			output: 0,
		},
		{
			input:  []int{4, 5, 6, 7, 0, 1, 2},
			output: 0,
		},
		{
			input:  []int{11, 13, 15, 17},
			output: 11,
		},
	}

	for i, c := range cases {
		result := findMin(c.input)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
