package leetcode230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var count, result int
	inOrderTraverse(root, k, &count, &result)
	return result
}

func inOrderTraverse(node *TreeNode, k int, count *int, result *int) {
	// If count is more than k, it means that we have found the kth smallest value
	if node == nil || *count >= k {
		return
	}

	inOrderTraverse(node.Left, k, count, result)

	// Counter +1 and return if it find the kth smallest value
	*count++
	if *count == k {
		*result = node.Val
		return
	}

	inOrderTraverse(node.Right, k, count, result)
}
