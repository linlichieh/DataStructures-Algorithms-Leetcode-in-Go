package heapSort

import (
	"container/heap"
)

type Node struct {
	Val  int
	Next *Node
}

type Nodes []*Node

func (h Nodes) Len() int           { return len(h) }
func (h Nodes) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h Nodes) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Nodes) Push(x any) {
	*h = append(*h, x.(*Node))
}

func (h *Nodes) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func sortList(node *Node) *Node {
	var nodes Nodes
	heap.Init(&nodes)

	// Add each node in the list into the heap
	curr := node
	for curr != nil {
		heap.Push(&nodes, curr)
		curr = curr.Next
	}

	// Pop the node from the heap and make it a link list
	dummy := &Node{}
	curr = dummy
	for nodes.Len() > 0 {
		popNode := heap.Pop(&nodes).(*Node)

		// Clear the next of original node
		popNode.Next = nil

		curr.Next = popNode
		curr = popNode
	}

	return dummy.Next
}
