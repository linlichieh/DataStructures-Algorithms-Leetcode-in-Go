package leetcode15

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	cases := []struct {
		input  []int
		output [][]int
	}{
		{
			input:  []int{},
			output: [][]int{},
		},
		{
			input:  []int{-1, 0, 1, 2, -1, -4},
			output: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			input:  []int{0, 1, 1},
			output: [][]int{},
		},
		{
			input:  []int{0, 0, 0},
			output: [][]int{{0, 0, 0}},
		},
		{
			input:  []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			output: [][]int{{0, 0, 0}}, // Notice that the solution set must not contain duplicate triplets.
		},
		{
			input:  []int{-1, 0, 1, 0},
			output: [][]int{{-1, 0, 1}},
		},
		{
			input:  []int{1, -1, -1, 0},
			output: [][]int{{-1, 0, 1}},
		},
		{
			input:  []int{3, -2, 1, 0},
			output: [][]int{},
		},
		{
			input:  []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
			output: [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}},
		},
		{
			input:  []int{-2, 0, 0, 2, 2},
			output: [][]int{{-2, 0, 2}},
		},
		{
			input:  []int{-2, 0, 1, 1, 2},
			output: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			input:  []int{-2, 0, 0, 0, 1, 1, 1, 1, 2},
			output: [][]int{{-2, 0, 2}, {-2, 1, 1}, {0, 0, 0}},
		},
	}

	for i, c := range cases {
		result := threeSum(c.input)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
