package leetcode21

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	cases := []struct {
		l1     *ListNode
		l2     *ListNode
		output *ListNode
	}{
		{
			l1:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
			l2:     &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
			output: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}},
		},
		{
			l1:     nil,
			l2:     nil,
			output: nil,
		},
		{
			l1:     nil,
			l2:     &ListNode{},
			output: &ListNode{},
		},
	}

	for _, c := range cases {
		result := mergeTwoLists(c.l1, c.l2)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("expect '%v', but got '%v'", c.output, result)
		}
	}
}
