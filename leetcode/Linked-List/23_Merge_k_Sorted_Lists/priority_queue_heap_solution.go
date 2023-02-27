package leetcode23

import (
	"container/heap"
)

type minHeap []*ListNode

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i int, j int) bool {
	return h[i].Val < h[j].Val
}

func (h minHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(item interface{}) {
	*h = append(*h, item.(*ListNode))
}

func (h *minHeap) Pop() interface{} {
	n := len(*h)
	min := (*h)[n-1] // the last one is the minimum because swap() and heapifyDown() have already been executed by the built-in Pop()
	*h = (*h)[:n-1]
	return min
}

func mergeKLists(lists []*ListNode) *ListNode {
	// Initialise my heap
	h := &minHeap{}
	heap.Init(h)

	// Add the first node of each list into the heap in the first iteration only
	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
			l = l.Next
		}
	}

	// Extract the minimum and add it into the merged list
	dummy := &ListNode{}
	curr := dummy
	for h.Len() > 0 {
		// Get the minimum
		minimum := heap.Pop(h)

		// Add minimum to the merged list
		curr.Next = minimum.(*ListNode)

		// Push the next node of minimum
		if minimum.(*ListNode).Next != nil {
			heap.Push(h, minimum.(*ListNode).Next)
		}

		curr = curr.Next
	}

	return dummy.Next
}
