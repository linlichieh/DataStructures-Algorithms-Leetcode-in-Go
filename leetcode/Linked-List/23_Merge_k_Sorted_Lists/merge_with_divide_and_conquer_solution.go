package leetcode23

type ListNode struct {
	Val  int
	Next *ListNode
}

type heap struct {
	root *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	return merge(lists, 0, len(lists)-1)
}

// For example, slice size = 8

// merge(0,7)
//
//	  = l1 merge(0,3)
//						= l1 merge(0,1)
//										= l1 merge(0,0) return 0
//										+ l2 merge(1,1) return 1
//						+ l2 merge(2,3)
//										= l1 merge(2,2) return 2
//										+ l2 merge(3,3) return 3
//	  + l2 merge(4,7)
//						= l1 merge(4,5)
//										= l1 merge(4,4) return 4
//										+ l2 merge(5,5) return 5
//						+ l2 merge(6,7)
//										= l1 merge(6,6) return 6
//										+ l2 merge(7,7) return 7
func merge(lists []*ListNode, left int, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := (left + right) / 2
	l1 := merge(lists, left, mid)
	l2 := merge(lists, mid+1, right)
	return mergeTwoLists(l1, l2)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		// if l1 is less than l2, append l1 to merged list
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	// If l1 is not empty, append the rest of l1 to merged list
	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}
	return dummy.Next
}
