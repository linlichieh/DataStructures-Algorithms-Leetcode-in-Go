package leetcode206

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	cases := []struct {
		input  *ListNode
		output *ListNode
	}{
		{
			input:  &ListNode{},
			output: &ListNode{},
		},
		{
			input:  &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			output: &ListNode{Val: 2, Next: &ListNode{Val: 1}},
		},
		{
			input:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			output: &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}},
		},
	}

	for _, c := range cases {
		result := reverseList(c.input)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("expect '%v', but got '%v'", c.output, result)
		}
	}
}
