package bubbleSort

func bubble(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		swapped := false
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// If no two elements were swapped by inner loop, it indicates the numbers are sorted
		if !swapped {
			break
		}
	}
	return arr
}
