package quickSort

func quick(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	quickRecur(arr, 0, len(arr)-1)
	return arr
}

func quickRecur(arr []int, low int, high int) {
	if high < low {
		return
	}

	// Set the pivot
	pivot := arr[high]

	// i is the starting point
	i := low
	for j := i; j < high; j++ {
		if arr[j] < pivot {
			if j != i {
				// Swap when the value at index j is less than pivot
				swap(arr, i, j)
			}
			i++
		}
	}

	// Values starting at and after i are greater than the pivot.
	// Swap pivot with i to ensure that elements after pivot are greater than it, in order to divide the array into two based on the pivot.
	// high is pivot's index
	swap(arr, high, i)
	quickRecur(arr, 0, i-1)
	quickRecur(arr, i+1, high)
}

func swap(arr []int, idx1 int, idx2 int) {
	arr[idx1], arr[idx2] = arr[idx2], arr[idx1]
}
