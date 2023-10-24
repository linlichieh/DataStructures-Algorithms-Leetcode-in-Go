package leetcode295

import (
	"container/heap"
)

type minHeap []int
type maxHeap []int

type MedianFinder struct {
	small maxHeap
	large minHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		small: []int{},
		large: []int{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.small.Len() == 0 || num <= this.small[0] {
		heap.Push(&this.small, num)
	} else {
		heap.Push(&this.large, num)
	}

	// Balance heaps
	// small large
	//   1     0
	//   1     1
	//   2     1
	//   3     1     <- balance
	//   1     2     <- balance
	if this.small.Len() > this.large.Len()+1 {
		heap.Push(&this.large, heap.Pop(&this.small))
	} else if this.large.Len() > this.small.Len() {
		heap.Push(&this.small, heap.Pop(&this.large))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	// Avoid raising panic when both heap are empty
	if this.small.Len() == 0 && this.large.Len() == 0 {
		return float64(0)
	}
	if this.small.Len() > this.large.Len() {
		return float64(this.small[0])
	}
	return float64(this.small[0]+this.large[0]) / 2
}

// min heap implementation
func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// max heap implementation
func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
