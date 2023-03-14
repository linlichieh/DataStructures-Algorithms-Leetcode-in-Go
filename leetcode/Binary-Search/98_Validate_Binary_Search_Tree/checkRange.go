package leetcode98

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return checkRange(root, math.MinInt64, math.MaxInt64)
}

func checkRange(root *TreeNode, min int64, max int64) bool {
	if root == nil {
		return true
	}
	if int64(root.Val) <= min || int64(root.Val) >= max {
		return false
	}

	return checkRange(root.Left, min, int64(root.Val)) && checkRange(root.Right, int64(root.Val), max)
}
