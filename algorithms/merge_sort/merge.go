package mergeSort

func merge(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	midIdx := len(arr) / 2
	left := arr[0:midIdx]
	right := arr[midIdx:]

	return mergeArrays(merge(left), merge(right))
}

func mergeArrays(leftArr []int, rightArr []int) []int {
	var mergedArr []int
	var leftIdx, rightIdx int
	for leftIdx < len(leftArr) && rightIdx < len(rightArr) {
		if leftArr[leftIdx] < rightArr[rightIdx] {
			mergedArr = append(mergedArr, leftArr[leftIdx])
			leftIdx++

		} else {
			mergedArr = append(mergedArr, rightArr[rightIdx])
			rightIdx++
		}
	}

	if leftIdx == len(leftArr) {
		// If all left numbers have been added into the merged array, then add the rest of the right array into the merged array
		mergedArr = append(mergedArr, rightArr[rightIdx:]...)
	} else {
		mergedArr = append(mergedArr, leftArr[leftIdx:]...)
	}

	return mergedArr
}
