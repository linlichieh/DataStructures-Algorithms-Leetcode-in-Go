package leetcode450

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		// node matched

		// no child and 1 child scenario
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// Find the leftmost in right subtree
		leftmost := root.Right
		for leftmost.Left != nil {
			leftmost = leftmost.Left
		}

		// Remove the leftmost from the node that matches the key
		root = deleteNode(root, leftmost.Val)

		// Replace the root value with leftmost value
		root.Val = leftmost.Val
	}
	return root
}
