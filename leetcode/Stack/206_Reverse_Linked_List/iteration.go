package leetcode206

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// iterative solution
func reverseList(head *ListNode) *ListNode {
	var prev, tmp *ListNode
	curr := head
	for curr != nil {
		tmp = curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}

// for debug
func print(root *ListNode) {
	for root != nil {
		fmt.Printf("%d -> ", root.Val)
		root = root.Next
	}
	fmt.Printf("nil\n")
}
