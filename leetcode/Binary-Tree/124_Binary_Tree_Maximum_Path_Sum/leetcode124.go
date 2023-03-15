package leetcode124

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	// if no child, just return its own value
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	// use MinInt32 instead of int due to the tree with all negative
	maximum := math.MinInt32
	recurGetMaximum(root, &maximum)
	return maximum
}

func recurGetMaximum(root *TreeNode, maximum *int) int {
	if root == nil {
		return 0
	}

	// get the left or right max value if existed
	leftMax := recurGetMaximum(root.Left, maximum)
	// compare with 0 because the node doesn't have to be combined if it's negative value
	leftMax = max(leftMax, 0)
	rightMax := recurGetMaximum(root.Right, maximum)
	rightMax = max(rightMax, 0)

	// get the maximum of split path
	splitMax := root.Val + leftMax + rightMax
	*maximum = max(*maximum, splitMax)

	// return the maximum without spliting
	return root.Val + max(leftMax, rightMax)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
