package leetcode98

func isValidBST2(root *TreeNode) bool {
	// Use in-order traversal to sort the tree
	var list []int
	inOrder(root, &list)

	// Compare each element with previous one
	for i := 0; i < len(list)-1; i++ {
		if list[i] >= list[i+1] {
			return false
		}
	}

	return true
}

func inOrder(root *TreeNode, list *[]int) {
	if root.Left != nil {
		inOrder(root.Left, list)
	}

	*list = append(*list, root.Val)

	if root.Right != nil {
		inOrder(root.Right, list)
	}
}
