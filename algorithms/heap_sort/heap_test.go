package heapSort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	cases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{},
			output: []int{},
		},
		{
			input:  []int{1},
			output: []int{1},
		},
		{
			input:  []int{0, 1},
			output: []int{0, 1},
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
			input:  []int{7, 1, 2, 8, 4},
			output: []int{1, 2, 4, 7, 8},
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
		{
			input:  []int{10, 11, 23, 44, 8, 15, 3, 9, 12, 45, 56, 45, 45},
			output: []int{3, 8, 9, 10, 11, 12, 15, 23, 44, 45, 45, 45, 56},
		},
		{
			input:  []int{11, 10, 23, 44, 8, 15, 3, 9, 12, 35, 56, 45},
			output: []int{3, 8, 9, 10, 11, 12, 15, 23, 35, 44, 45, 56},
		},
	}

	for _, c := range cases {
		result := sort(c.input)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("expect '%v', but got '%v'", c.output, result)
		}
	}
}
