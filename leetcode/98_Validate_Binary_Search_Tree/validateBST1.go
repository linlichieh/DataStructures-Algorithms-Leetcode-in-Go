package leetcode98

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST1(root *TreeNode) bool {
	return traverseBST1(root, math.MinInt64, math.MaxInt64)
}

func traverseBST1(root *TreeNode, min int64, max int64) bool {
	if root == nil {
		return true
	}
	if int64(root.Val) <= min || int64(root.Val) >= max {
		return false
	}

	return traverseBST1(root.Left, min, int64(root.Val)) && traverseBST1(root.Right, int64(root.Val), max)
}
