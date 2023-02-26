package leetcode23

func CompareOneByOne(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	minIdx := -1
	for i, l := range lists {
		if l == nil {
			continue
		}
		if minIdx < 0 {
			minIdx = i
			continue
		}
		if lists[minIdx].Val > l.Val {
			minIdx = i
		}
	}
	if minIdx < 0 {
		return nil
	}
	minVal := lists[minIdx].Val
	lists[minIdx] = lists[minIdx].Next

	return &ListNode{
		Val:  minVal,
		Next: mergeKLists(lists),
	}
}
