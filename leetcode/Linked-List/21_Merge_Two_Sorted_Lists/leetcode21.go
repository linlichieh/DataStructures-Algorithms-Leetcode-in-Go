package leetcode21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	var min *ListNode
	if list1 != nil && (list2 == nil || list1.Val < list2.Val) {
		min = list1
		list1 = list1.Next
	} else {
		min = list2
		list2 = list2.Next
	}

	return &ListNode{
		Val:  min.Val,
		Next: mergeTwoLists(list1, list2),
	}
}
