package leetcode417

import (
	"reflect"
	"testing"
)

func TestPacificAtlantic(t *testing.T) {
	cases := []struct {
		input  [][]int
		output [][]int
	}{
		{
			input:  [][]int{},
			output: [][]int{},
		},
		{
			input:  [][]int{{1}},
			output: [][]int{{0, 0}},
		},
		{
			input: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			output: [][]int{
				{0, 4},
				{1, 3},
				{1, 4},
				{2, 2},
				{3, 0},
				{3, 1},
				{4, 0},
			},
		},
		{
			input: [][]int{
				{1, 3, 5, 5},
				{7, 5, 6, 5},
				{2, 4, 8, 4},
				{3, 1, 2, 6},
			},
			output: [][]int{
				{0, 2},
				{0, 3},
				{1, 0},
				{1, 1},
				{1, 2},
				{1, 3},
				{2, 1},
				{2, 2},
				{3, 0},
			},
		},
	}

	for i, c := range cases {
		result := pacificAtlantic(c.input)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
