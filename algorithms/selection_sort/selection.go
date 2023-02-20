package selectionSort

func selection(arr []int) []int {
	var tmp, minIdx int
	for i := 0; i < len(arr)-1; i++ {
		tmp = arr[i]
		minIdx = i
		for j := i; j < len(arr)-1; j++ {
			if arr[minIdx] > arr[j+1] {
				minIdx = j + 1
			}
		}
		arr[i] = arr[minIdx]
		arr[minIdx] = tmp
	}
	return arr
}
