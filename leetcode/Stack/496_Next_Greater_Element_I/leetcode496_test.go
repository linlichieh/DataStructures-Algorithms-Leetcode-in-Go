package leetcode496

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	cases := []struct {
		nums1  []int
		nums2  []int
		result []int
	}{
		{
			nums1:  []int{4, 1, 2},
			nums2:  []int{1, 3, 4, 2},
			result: []int{-1, 3, -1},
		},
		{
			nums1:  []int{4, 1, 2},
			nums2:  []int{2, 1, 3, 4},
			result: []int{-1, 3, 3},
		},
		{
			nums1:  []int{4, 1, 2},
			nums2:  []int{2, 4, 1, 3},
			result: []int{-1, 3, 4},
		},
		{
			nums1:  []int{4, 1, 2},
			nums2:  []int{2, 4, 1, 6},
			result: []int{6, 6, 4},
		},
		{
			nums1:  []int{2, 4},
			nums2:  []int{1, 2, 3, 4},
			result: []int{3, -1},
		},
		{
			nums1:  []int{2, 0, 1},
			nums2:  []int{2, 0, 3, 4},
			result: []int{3, 3, -1},
		},
	}

	for _, c := range cases {
		result := nextGreaterElement(c.nums1, c.nums2)
		if !reflect.DeepEqual(c.result, result) {
			t.Errorf("expect '%v', but got '%v'", c.result, result)
		}
	}
}
