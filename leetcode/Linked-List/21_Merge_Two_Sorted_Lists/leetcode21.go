package leetcode21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var min *ListNode
	if list1.Val < list2.Val {
		min = list1
		list1 = list1.Next
	} else {
		min = list2
		list2 = list2.Next
	}
	min.Next = mergeTwoLists(list1, list2)
	return min
}
