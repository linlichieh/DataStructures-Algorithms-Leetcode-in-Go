package leetcode141

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// The solution called "slow and fast pointers" which is to have 2 pointers traversing through the list at different speed
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	// We don't need to worry about whether slow is nil because if it's not a cycle list, the fast will break the loop first
	for fast != nil && fast.Next != nil {
		// If it's a cycle list, these 2 pointers will be the same eventually
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

// for debug
func print(l *ListNode) {
	m := map[*ListNode]bool{}

	for l != nil {
		// Used for check if it's cycle lists
		if _, ok := m[l]; !ok {
			m[l] = true
		} else {
			// Cycle
			fmt.Printf("%d (cycle)\n", l.Val)
			return
		}

		fmt.Printf("%d -> ", l.Val)
		l = l.Next
	}
	fmt.Printf("nil (non-cycle)\n")
}
