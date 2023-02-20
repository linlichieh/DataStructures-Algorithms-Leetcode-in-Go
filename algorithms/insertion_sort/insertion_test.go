package insertionSort

import (
	"reflect"
	"testing"
)

func TestInsertion(t *testing.T) {
	cases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{},
			output: []int{},
		},
		{
			input:  []int{0, 0, 0},
			output: []int{0, 0, 0},
		},
		{
			input:  []int{0, 1, 0},
			output: []int{0, 0, 1},
		},
		{
			input:  []int{6, 8, 1, 4, 9, 7, 2},
			output: []int{1, 2, 4, 6, 7, 8, 9},
		},
		{
			input:  []int{2, -4, 8, 3, 5, 7, 0},
			output: []int{-4, 0, 2, 3, 5, 7, 8},
		},
		{
			input:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			output: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			input:  []int{6, 5, 3, 1, 8, 7, 2, 4},
			output: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, c := range cases {
		result := insertion(c.input)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("expect '%v', but got '%v'", c.output, result)
		}
	}
}
