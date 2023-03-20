package leetcode347

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	cases := []struct {
		nums   []int
		k      int
		output []int
	}{
		{
			nums:   []int{1, 1, 1, 2, 2, 3},
			k:      2,
			output: []int{1, 2},
		},
		{
			nums:   []int{1},
			k:      1,
			output: []int{1},
		},
		{
			nums:   []int{3, 3, 4, 5, 5, 5, 6, 6, 6, 7},
			k:      3,
			output: []int{5, 6, 3},
		},
		{
			nums:   []int{3, 3, 4, 5, 5, 5, 6, 6, 6, 7},
			k:      2,
			output: []int{5, 6},
		},
		{
			nums:   []int{1, 2, 3, 1, 2, 1, 4, 3},
			k:      2,
			output: []int{1, 2},
		},
		{
			nums:   []int{-1, -1},
			k:      1,
			output: []int{-1},
		},
	}

	for i, c := range cases {
		result := topKFrequent(c.nums, c.k)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}
