package leetcode82

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// The first node and the next could be the same, in order to deal with it, it would be easier not to assign the head to dummy but its Next, e.g. 1->1->2->3
	dummy := &ListNode{Next: head}
	// prev only does point to the curr that has been processed
	// curr's job is to find the non-duplicate one
	prev, curr := dummy, dummy.Next
	for curr != nil {
		var dup bool
		// To skip the duplicate node
		for curr.Next != nil && curr.Val == curr.Next.Val {
			dup = true
			curr = curr.Next
		}
		if dup {
			// If it's duplicated, point to the next one, but don't move the prev to the next one in order to handle the case that
			// node is duplicated at the end e.g. 1->2->3->3, just point the prev to nil if needed
			// If the case is 1->2->3->3->4->4 that curr.Next is 4 that is also a duplicated, it's okay, because we didn't move the
			// prev to the next one, so the value would be overwritten to nil after this iteration
			prev.Next = curr.Next
		} else {
			// The best case is that the curr node isn't duplicated, then just move the prev to the curr
			prev = curr
		}
		curr = curr.Next
	}
	return dummy.Next
}

// for debug
func print(l *ListNode) {
	for l != nil {
		fmt.Printf("%d -> ", l.Val)
		l = l.Next
	}
	fmt.Printf("nil\n")
}
