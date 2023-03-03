package leetcode1

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		output []int
	}{
		// NOTE the order of the indices in output doesn't matter
		// NOTE some solutions might be more than 2, one of them can pass
		{
			nums:   []int{},
			target: 0,
			output: []int{},
		},
		{
			nums:   []int{},
			target: 0,
			output: []int{},
		},
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			output: []int{1, 0},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			output: []int{2, 1},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			output: []int{1, 0},
		},
		{
			nums:   []int{1, 2, 3, 4, 5},
			target: 9,
			output: []int{4, 3},
		},
		{
			nums:   []int{1, 2, 3, 4, 5},
			target: 10,
			output: []int{},
		},
		{
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			output: []int{4, 2}},
		{
			nums:   []int{0, 0, 0},
			target: 0,
			output: []int{1, 0},
		},
		{
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
			output: []int{1, 0},
		},
		{
			nums:   []int{1, 2, 3, 4, 5},
			target: 6,
			output: []int{3, 1},
		},
	}

	for i, c := range cases {
		result := twoSum(c.nums, c.target)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
