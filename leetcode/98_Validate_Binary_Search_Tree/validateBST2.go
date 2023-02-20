package leetcode98

func isValidBST2(root *TreeNode) bool {
	// Use in-order traversal to sort the tree
	var list []int
	traverseBST2(root, &list)

	// Compare each element with previous one
	for i := 0; i < len(list)-1; i++ {
		if list[i] >= list[i+1] {
			return false
		}
	}

	return true
}

func traverseBST2(root *TreeNode, list *[]int) {
	if root.Left != nil {
		traverseBST2(root.Left, list)
	}

	*list = append(*list, root.Val)

	if root.Right != nil {
		traverseBST2(root.Right, list)
	}
}
