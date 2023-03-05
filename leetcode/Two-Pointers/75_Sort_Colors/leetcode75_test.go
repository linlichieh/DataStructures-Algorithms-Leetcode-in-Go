package leetcode75

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	cases := []struct {
		input  []int
		result []int
	}{
		{
			input:  []int{},
			result: []int{},
		},
		{
			input:  []int{0},
			result: []int{0},
		},
		{
			input:  []int{1, 1, 1},
			result: []int{1, 1, 1},
		},
		{
			input:  []int{2, 1, 0},
			result: []int{0, 1, 2},
		},
		{
			input:  []int{2, 0, 1},
			result: []int{0, 1, 2},
		},
		{
			input:  []int{0, 1, 2},
			result: []int{0, 1, 2},
		},
		{
			input:  []int{2, 0, 2, 1, 1, 0},
			result: []int{0, 0, 1, 1, 2, 2},
		},
		{
			input:  []int{2, 2, 1, 1, 0, 0},
			result: []int{0, 0, 1, 1, 2, 2},
		},
		{
			input:  []int{1, 0, 1, 2, 1, 0, 2},
			result: []int{0, 0, 1, 1, 1, 2, 2},
		},
		{
			input:  []int{1, 2, 0, 2, 0, 1},
			result: []int{0, 0, 1, 1, 2, 2},
		},
	}

	for i, c := range cases {
		// in-place
		sortColors(c.input)
		if !reflect.DeepEqual(c.result, c.input) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.result, c.input)
		}
	}
}
