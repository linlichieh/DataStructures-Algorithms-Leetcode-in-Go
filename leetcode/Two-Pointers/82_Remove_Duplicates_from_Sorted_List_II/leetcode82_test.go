package leetcode82

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	cases := []struct {
		input  *ListNode
		result *ListNode
	}{
		{
			input:  &ListNode{},
			result: &ListNode{},
		},
		{
			input:  &ListNode{Val: 1, Next: &ListNode{Val: 1}},
			result: nil,
		},
		{
			input:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}},
			result: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			input:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}},
			result: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			input:  &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}}},
			result: &ListNode{Val: 2},
		},

		{
			input:  &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			result: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			input:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}}},
			result: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}},
		},
		{
			input:  &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}},
			result: &ListNode{Val: 2, Next: &ListNode{Val: 3}},
		},
		{
			input:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}}}},
			result: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}},
		},
		{
			input:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}},
			result: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}},
		},
	}

	for i, c := range cases {
		result := deleteDuplicates(c.input)
		if !reflect.DeepEqual(c.result, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.result, result)
		}
	}
}
