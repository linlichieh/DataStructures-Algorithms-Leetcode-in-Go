package leetcode283

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	cases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{},
			output: []int{},
		},
		{
			input:  []int{1, 2, 3, 4},
			output: []int{24, 12, 8, 6},
		},
		{
			input:  []int{1, -2, 3, -4},
			output: []int{24, -12, 8, -6},
		},
		{
			input:  []int{-1, 1, 0, -3, 3},
			output: []int{0, 0, 9, 0, 0},
		},
		{
			input:  []int{1, 0, 3, 0},
			output: []int{0, 0, 0, 0},
		},
	}

	for i, c := range cases {
		result := productExceptSelf(c.input)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
