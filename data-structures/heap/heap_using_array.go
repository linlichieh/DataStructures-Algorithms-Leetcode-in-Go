package heap

import (
	"math"
)

// left(i) = i*2 + 1
// right(i) = i*2 + 2
// parent(i) = (i-1)/2

type MinHeapArray struct {
	list []int
}

func (h *MinHeapArray) Insert(val int) {
	h.list = append(h.list, val)
	newIdx := len(h.list) - 1
	h.HeapifyUp(newIdx)
}

// Maintain the heap property after an insertion or deletion
// the parent node must be less than or equal to (for a min heap) its children nodes
func (h *MinHeapArray) HeapifyUp(idx int) {
	// Fix the min heap property if it is violated
	if idx < 1 {
		return
	}

	parentIdx := int(math.Floor(float64((idx - 1) / 2)))
	if h.list[idx] < h.list[parentIdx] {
		h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
	}
	h.HeapifyUp(parentIdx)
}

// Remove the root element because it's the minimum element
func (h *MinHeapArray) ExtractMin() {
	size := len(h.list)

	if size == 0 {
		return
	}

	// Swap the root element with the last element in the heap
	h.list[0], h.list[size-1] = h.list[size-1], h.list[0]

	// Remove the last element from the heap
	h.list = h.list[0 : size-1]

	// Compare the root element with its children and swap it with the smaller child
	h.HeapifyDown(0)
}

// root index
func (h *MinHeapArray) HeapifyDown(root int) {
	l := root*2 + 1 // left index
	r := root*2 + 2 // right index

	// The root node should be swapped with the node that is smaller between the left and right nodes
	smallest := root
	if len(h.list) > l && h.list[l] < h.list[root] {
		smallest = l
	}
	if len(h.list) > r && h.list[r] < h.list[smallest] {
		smallest = r
	}
	if smallest != root {
		h.list[root], h.list[smallest] = h.list[smallest], h.list[root]
		h.HeapifyDown(smallest)
	}
}
