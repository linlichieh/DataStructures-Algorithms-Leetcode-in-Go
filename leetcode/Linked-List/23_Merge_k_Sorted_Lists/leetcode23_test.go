package leetcode23

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	cases := []struct {
		input  []*ListNode
		output *ListNode
	}{
		{
			input: []*ListNode{
				&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
				&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
				&ListNode{Val: 2, Next: &ListNode{Val: 6}},
			},
			output: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}},
		},
		{
			input:  []*ListNode{},
			output: nil,
		},
		{
			input:  []*ListNode{&ListNode{}},
			output: &ListNode{},
		},
		{
			input:  []*ListNode{&ListNode{}, &ListNode{}, &ListNode{}},
			output: &ListNode{Next: &ListNode{Next: &ListNode{}}},
		},
		{
			input:  []*ListNode{nil},
			output: nil,
		},
		{
			input:  []*ListNode{nil, nil},
			output: nil,
		},
		{
			input:  []*ListNode{&ListNode{}, &ListNode{Val: 1}},
			output: &ListNode{Val: 0, Next: &ListNode{Val: 1}},
		},
	}

	for _, c := range cases {
		result := mergeKLists(c.input)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("expect '%v', but got '%v'", c.output, result)
		}
	}
}
