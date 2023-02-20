package heapSort

func sort(arr []int) []int {
	// Make unsorted array a max-heap
	// Tip: sorting only the first 4 elements will result in the entire array being sorted.
	size := len(arr)
	for i := size/2 - 1; i >= 0; i-- {
		heapify(arr, size, i)
	}

	// Extract maximum from the root
	for i := size - 1; i > 0; i-- {
		// Swap maximum with the last element
		arr[0], arr[i] = arr[i], arr[0]

		// Call heapify
		heapify(arr, i, 0)
	}

	return arr
}

// Size is for only sorting unordered elements because the elements that are shifted to the right are sorted
func heapify(arr []int, size int, i int) {
	l := i*2 + 1
	r := i*2 + 2

	largest := i
	if l < size && arr[i] < arr[l] {
		largest = l
	}
	if r < size && arr[largest] < arr[r] {
		largest = r
	}

	if i == largest {
		return
	}

	arr[i], arr[largest] = arr[largest], arr[i]
	heapify(arr, size, largest)
}
