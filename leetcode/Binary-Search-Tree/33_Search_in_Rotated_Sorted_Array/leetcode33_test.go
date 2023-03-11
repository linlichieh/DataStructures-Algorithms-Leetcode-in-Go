package leetcode33

import (
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		output int
	}{
		{
			nums:   []int{},
			target: 0,
			output: -1,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 4,
			output: 0,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 5,
			output: 1,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 6,
			output: 2,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 7,
			output: 3,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			output: 4,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 1,
			output: 5,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 6,
			output: 2,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			output: -1,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1},
			target: 4,
			output: 0,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1},
			target: 5,
			output: 1,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1},
			target: 6,
			output: 2,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1},
			target: 7,
			output: 3,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1},
			target: 0,
			output: 4,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1},
			target: 1,
			output: 5,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1},
			target: 3,
			output: -1,
		},
		{
			nums:   []int{1},
			target: 0,
			output: -1,
		},
		{
			nums:   []int{1},
			target: 1,
			output: 0,
		},
	}

	for i, c := range cases {
		result := search(c.nums, c.target)
		if c.output != result {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
