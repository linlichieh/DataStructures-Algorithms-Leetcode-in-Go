package leetcode2

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return add(l1, l2, 0)
}

func add(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	// if there is no node from l1 and l2 and no carryover
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	// Get the carryover from the previous call and add it to sum
	sum := carry
	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}

	// Calculate the carryover value and the current sum value.
	carry = 0
	if sum > 9 {
		carry = 1
		sum -= 10
	}

	// Create a new node and call the add function recursively for the `Next` node
	return &ListNode{
		Val:  sum,
		Next: add(l1, l2, carry),
	}
}

func print(n *ListNode) {
	if n == nil {
		return
	}
	fmt.Printf("%+v\n", n)
	if n.Next != nil {
		print(n.Next)
	}
}
