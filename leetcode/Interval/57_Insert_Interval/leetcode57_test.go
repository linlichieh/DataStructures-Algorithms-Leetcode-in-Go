package leetcode57

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	cases := []struct {
		param1   [][]int
		param2   []int
		expected [][]int
	}{
		{
			param1:   [][]int{},
			param2:   []int{2, 5},
			expected: [][]int{{2, 5}},
		},
		{
			param1:   [][]int{{1, 3}, {6, 9}},
			param2:   []int{2, 5},
			expected: [][]int{{1, 5}, {6, 9}},
		},
		{
			param1:   [][]int{{3, 4}, {6, 9}},
			param2:   []int{2, 5},
			expected: [][]int{{2, 5}, {6, 9}},
		},
		{
			param1:   [][]int{{6, 9}, {13, 15}},
			param2:   []int{2, 5},
			expected: [][]int{{2, 5}, {6, 9}, {13, 15}},
		},
		{
			param1:   [][]int{{6, 9}, {13, 15}},
			param2:   []int{2, 7},
			expected: [][]int{{2, 9}, {13, 15}},
		},
		{
			param1:   [][]int{{6, 9}, {13, 15}},
			param2:   []int{2, 14},
			expected: [][]int{{2, 15}},
		},
		{
			param1:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			param2:   []int{4, 8},
			expected: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			param1:   [][]int{{1, 3}, {7, 10}},
			param2:   []int{4, 6},
			expected: [][]int{{1, 3}, {4, 6}, {7, 10}},
		},
		{
			param1:   [][]int{{1, 3}, {7, 10}},
			param2:   []int{14, 16},
			expected: [][]int{{1, 3}, {7, 10}, {14, 16}},
		},
	}

	for i, c := range cases {
		result := insert(c.param1, c.param2)
		if !reflect.DeepEqual(c.expected, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.expected, result)
		}
	}
}
