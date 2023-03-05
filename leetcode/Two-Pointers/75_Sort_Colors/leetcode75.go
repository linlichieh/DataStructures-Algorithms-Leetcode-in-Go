package leetcode75

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}

	// Use 3 pointers approach
	//  - pointer0 at the beginer of array
	//  - pointer2 at the end of array
	//  - curr pointer at the beginner and for checking the value
	ptr0, curr, ptr2 := 0, 0, len(nums)-1
	for curr <= ptr2 {
		switch nums[curr] {
		case 0:
			// swap the value `0` to the left
			nums[curr], nums[ptr0] = nums[ptr0], nums[curr]
			curr++
			ptr0++
			// current pointer moves on
		case 1:
			curr++
		case 2:
			// swap the value `2` to the right
			nums[curr], nums[ptr2] = nums[ptr2], nums[curr]
			ptr2--
		}
	}
}
