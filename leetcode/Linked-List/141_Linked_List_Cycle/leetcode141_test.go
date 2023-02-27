package leetcode141

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	cases := []struct {
		input    *ListNode
		pos      int
		hasCycle bool
	}{
		{
			input:    &ListNode{},
			pos:      -1,
			hasCycle: false,
		},
		{
			input:    &ListNode{},
			pos:      0,
			hasCycle: false,
		},
		{
			input:    &ListNode{Val: 3},
			pos:      -1,
			hasCycle: false,
		},
		{
			input:    &ListNode{Val: 3},
			pos:      0,
			hasCycle: false,
		},
		{
			input:    &ListNode{Val: 3, Next: &ListNode{Val: 4}},
			pos:      -1,
			hasCycle: false,
		},
		{
			input:    &ListNode{Val: 3, Next: &ListNode{Val: 4}},
			pos:      0,
			hasCycle: true,
		},
		{
			input:    &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
			pos:      -1,
			hasCycle: false,
		},
		{
			input:    &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
			pos:      0,
			hasCycle: true,
		},
		{
			input:    &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 4}}}},
			pos:      -1,
			hasCycle: false,
		},
		{
			input:    &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 4}}}},
			pos:      0,
			hasCycle: true,
		},
		{
			input:    &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 4}}}},
			pos:      1,
			hasCycle: true,
		},
		{
			input:    &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 4}}}},
			pos:      2,
			hasCycle: true,
		},
		{
			input:    &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 4}}}},
			pos:      3,
			hasCycle: false,
		},
	}

	for _, c := range cases {
		// Can't make a cycle linked list with one declaration
		// Call this method for this pupose
		l := makeCycleList(c.input, c.pos)

		result := hasCycle(l)
		if result != c.hasCycle {
			t.Errorf("expect '%v', but got '%v'", c.hasCycle, result)
		}
	}
}

func makeCycleList(l *ListNode, pos int) *ListNode {
	if pos < 0 {
		return l
	}

	var posCounter int
	var posNode *ListNode
	var lastNode *ListNode

	curr := l
	for curr != nil {
		// Get the pos node
		if posCounter == pos {
			posNode = curr
		}

		// Get the last node for later use
		if curr.Next == nil {
			lastNode = curr
		}
		curr = curr.Next
		posCounter++
	}

	// Point last.Next to posNode
	if posNode != lastNode {
		lastNode.Next = posNode
	}

	return l
}
