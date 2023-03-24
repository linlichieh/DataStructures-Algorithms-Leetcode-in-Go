package leetcode226

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Same as postorder implementation
	if root.Left != nil {
		invertTree(root.Left)
	}

	if root.Right != nil {
		invertTree(root.Right)
	}

	// Swap left and right nodes to invert the tree
	root.Left, root.Right = root.Right, root.Left
	return root
}
